package handler

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/forms"
	"github.com/pocketbase/pocketbase/mails"
	m "github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tokens"
	"github.com/xnpltn/hcc/internal/models"
	form "github.com/xnpltn/hcc/internal/templates/forms"
	"github.com/xnpltn/hcc/internal/templates/pages"
	"github.com/xnpltn/hcc/internal/utils"
)

func RenderView(c echo.Context, view templ.Component, layoutPath string) error {
	if c.Request().Header.Get("Hx-Request") == "true" {
		return view.Render(c.Request().Context(), c.Response().Writer)
	}
	return pages.AuthPage(layoutPath).Render(c.Request().Context(), c.Response().Writer)

}

func LoginForn() echo.HandlerFunc {
	return func(c echo.Context) error {
		return RenderView(c, form.Login(), "/auth/login")
	}
}

func SignupForn() echo.HandlerFunc {
	return func(c echo.Context) error {
		return RenderView(c, form.Signup(), "/auth/signup")
	}
}

func SignupHandler(app *pocketbase.PocketBase) echo.HandlerFunc {
	return func(c echo.Context) error {
		collection, err := app.Dao().FindCollectionByNameOrId("users")
		if err != nil {
			utils.WriteToLogs(err)
		}
		values, err := c.FormValues()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "something went wrong"})
		}
		record := m.NewRecord(collection)
		record.SetUsername(values.Get("username"))
		record.SetEmail(values.Get("email"))
		record.Set("name", values.Get("name"))
		record.SetPassword(values.Get("password"))
		if err := app.Dao().SaveRecord(record); err != nil {
			utils.WriteToLogs(err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "something went wrong"})
		}
		err = mails.SendRecordVerification(app, record)
		if err != nil {
			fmt.Println(err)
			utils.WriteToLogs(err)
		}
		return c.HTML(201, `
		<div class="flex items-center flex-col text-center">
		<p>Account created successfully <br/>
			Check your inbox to veirfy email.
		</p>
		</div>
	`)
	}
}

func LoginHandler(app *pocketbase.PocketBase) echo.HandlerFunc {
	return func(c echo.Context) error {
		formValues, err := c.FormValues()
		if err != nil {
			utils.WriteToLogs(err)
			return c.HTML(200, "<p>something went wrong</p>")
		}
		users, err := app.Dao().FindCollectionByNameOrId("users")
		if err != nil {
			utils.WriteToLogs(err)
			return c.HTML(200, "<p>something went wrong</p>")
		}
		form := forms.NewRecordPasswordLogin(app, users)
		form.Password = formValues.Get("password")
		form.Identity = formValues.Get("email")
		err = form.Validate()
		if err != nil {
			utils.WriteToLogs(err)
			return c.HTML(200, "<p>invalid username of password</p>")
		}
		authRecord, err := form.Submit()
		if err != nil {
			utils.WriteToLogs(err)
			return c.HTML(200, "<p>invalid username of password</p>")
		}
		token, err := tokens.NewRecordAuthToken(app, authRecord)
		if err != nil {
			return c.HTML(200, "<p>SOmething went wrong</p>")
		}
		c.SetCookie(&http.Cookie{
			Name:     "auth_token",
			Value:    token,
			Secure:   true,
			SameSite: http.SameSiteStrictMode,
			HttpOnly: true,
			MaxAge:   int(app.Settings().RecordAuthToken.Duration),
			Path:     "/",
		})
		return c.HTML(200, `
		<div class="flex items-center flex-col text-center">
		<p>Login Successful<br/>
			<a href="/" class="text-blue-700">go back to home</a>.
		</p>
		</div>
	`)
	}
}

func UpdateProfile(app *pocketbase.PocketBase) echo.HandlerFunc {
	return func(c echo.Context) error {
		values, err := c.FormValues()
		if err != nil {
			utils.WriteToLogs(err)
		}
		cookie, err := c.Cookie("auth_token")
		if err != nil {
			if !errors.Is(err, http.ErrNoCookie) {
				utils.WriteToLogs(err)
			}
		}
		if cookie != nil {
			authRecord, err := app.Dao().FindAuthRecordByToken(cookie.Value, app.Settings().RecordAuthToken.Secret)
			if err != nil {
				utils.WriteToLogs(err)
				return c.HTML(200, "<p>Profile Updated Something went wrong</p>")
			}
			err = updateCreds(app, authRecord, values)
			if err != nil {
				utils.WriteToLogs(fmt.Errorf("soemthing went wrong updating user profile: %w", err))
				return c.HTML(200, "<p>Something went wrong</p>")
			}
			err = mails.SendRecordVerification(app, authRecord)
			if err != nil {
				fmt.Println(err)
			}
			return c.HTML(
				200,
				`
				<div class="flex items-center flex-col text-center">
				<p>Profile Updated Successfully <br/>
					Check your inbox to veirfy email.
				</p>
				</div>
			`,
			)
		} else {
			return c.HTML(200, "<p>Something went wrong</p>")
		}
	}
}

func HandlerLogout(app *pocketbase.PocketBase) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("HX-Redirect", "/")

		c.SetCookie(&http.Cookie{
			Name:     "auth_token",
			Value:    "/",
			Secure:   true,
			SameSite: http.SameSiteStrictMode,
			HttpOnly: true,
			MaxAge:   0,
			Path:     "/",
		})
		return nil
	}
}

func updateCreds(app *pocketbase.PocketBase, authRecord *m.Record, values url.Values) error {
	if values.Get("name") != "" {
		authRecord.Set("name", values.Get("name"))
	}

	if values.Get("email") != "" {
		if err := authRecord.SetEmail(values.Get("email")); err != nil {
			return err
		}
	}
	if values.Get("username") != "" {
		if err := authRecord.SetUsername(values.Get("username")); err != nil {
			return err
		}
	}
	if values.Get("password") != "" {
		if err := authRecord.SetPassword(values.Get("password")); err != nil {
			return err
		}
	}
	if err := authRecord.SetVerified(false); err != nil {
		return err
	}
	return app.Dao().Save(authRecord)
}

func PassRequestHandler(app *pocketbase.PocketBase) echo.HandlerFunc {
	return func(c echo.Context) error {
		email := c.FormValue("email")

		if email == "" {
			return c.HTML(500, "please enter a valid email")
		}
		authCollection, err := app.Dao().FindCollectionByNameOrId("users")
		if err != nil {
			utils.WriteToLogs(err)
			return c.HTML(500, "something went wrong")
		}
		authRecord, err := app.Dao().FindAuthRecordByEmail(authCollection.Id, email)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return c.HTML(200, `
				<div class="flex items-center flex-col text-center">
				<p>Email not accociated with any account<br/>
					Click <a href="/auth/signup" class="text-blue-700">here</a> to create one.
				</p>
				</div>
			`)
			} else {
				return c.HTML(200, `
				<div class="flex items-center flex-col text-center">
				<p>
					Something went wrong
				</p>
				</div>
			`)
			}

		}
		err = mails.SendRecordPasswordReset(app, authRecord)
		if err != nil {
			utils.WriteToLogs(err)
			return c.HTML(
				200, `
			<div class="flex items-center flex-col text-center">
			<p>
			 Something went wrong<br/>
			</p>
			</div>
		`,
			)
		}

		return c.HTML(200, "link sent, check your inbox")
	}
}

func ChangePasshandler(app *pocketbase.PocketBase) echo.HandlerFunc {
	return func(c echo.Context) error {

		user := new(models.User)
		authRecord, err := app.Dao().FindAuthRecordByToken(c.PathParam("token"), app.Settings().RecordPasswordResetToken.Secret)
		if err != nil {
			fmt.Println(err)
			utils.WriteToLogs(err)
			return pages.ChangePasswordPage(*user).Render(c.Request().Context(), c.Response().Writer)
		}
		user.Email = authRecord.Email()
		user.Name = authRecord.GetString("name")
		user.Username = authRecord.Username()
		user.Verified = authRecord.Verified()
		user.ID = authRecord.Id
		return pages.ChangePasswordPage(*user).Render(c.Request().Context(), c.Response().Writer)
	}
}

func HandleResetpassword(app *pocketbase.PocketBase) echo.HandlerFunc {
	return func(c echo.Context) error {
		values, err := c.FormValues()
		if err != nil {
			utils.WriteToLogs(err)
			return c.HTML(200, "something went wrong changeing your password")

		}
		authRecord, err := app.Dao().FindAuthRecordByEmail("users", values.Get("email"))
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				c.HTML(http.StatusNotFound, "This email is not associated with any account or might be deleted")
			}
		}

		if values.Get("password") != values.Get("passwordconfirm") {
			return c.HTML(200, `
			<div class="flex items-center flex-col text-center">
			<p class="text-red-700" color="red">
				Passwords must be equal
			</p>
			</div>
		`)
		}

		err = authRecord.SetPassword(values.Get("password"))
		if err != nil {
			return c.HTML(200,
				`
				<div class="flex items-center flex-col text-center">
				<p class="text-red-700" color="red">
					Something went wrong
				</p>
				</div>
			`)
		}

		return c.HTML(
			201, `
			<div class="flex items-center flex-col text-center">
			<p>Password Changed Successfully<br/>
				Click <a href="/auth/login" class="text-blue-700">here</a> to login.
			</p>
			</div>
		`,
		)
	}
}
