package main

import (
	"log"
	"os"

	"github.com/stripe/stripe-go/v72"
	"github.com/xnpltn/hcc/internal/app"
)


func main(){
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
	err := app.App().Start()
	if err!= nil{
		log.Fatal(err)
	}
}