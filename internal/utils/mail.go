package utils

import (
	"net/mail"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/tools/mailer"
)





func SendMail(app *pocketbase.PocketBase, name string, email string, subject string, HTMLBody string) error {
	message := &mailer.Message{
		From: mail.Address{
			Address: app.Settings().Meta.SenderAddress,
			Name:    app.Settings().Meta.SenderName,
		},
		To:      []mail.Address{{Name: name, Address: email}},
		Subject: subject,
		HTML:    HTMLBody,
	}
	err := app.NewMailClient().Send(message)
	if err != nil {
		return err
	}
	return nil
}


