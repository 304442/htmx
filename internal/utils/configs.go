package utils

import (
	"fmt"
	"log"

	"github.com/pocketbase/pocketbase"
)

type Configs struct{
	StripePublishableKey string `db:"stripe_publishable_key"`
	StripeSecretKey string `db:"stripe_secret_key"`
	StripeWebhookKey string `db:"stripe_webhook_secret"`
	AdminEmail string  `db:"admin_email"`
}

func GetAdminConfigs(app *pocketbase.PocketBase) Configs{
	configs := Configs{}
	if err := app.Dao().DB().Select("*").From("configs").One(&configs); err!= nil{
		WriteToLogs(err)
		fmt.Println(err)
		log.Fatal(err)
	}
	return configs
}