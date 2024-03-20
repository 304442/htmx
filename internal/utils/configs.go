package utils

import (
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
		log.Println("Remember to add configuration keys in 'configs' collection. ")
		log.Print("required keys are: \n 'stripe_publishable_key' \n 'stripe_secret_key' \n'stripe_webhook_secret', \n 'admin_email'\n exactly with those names")
	}
	return configs
}