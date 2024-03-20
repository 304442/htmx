package main

import (
	"errors"
	"log"
	"os"

	// "github.com/stripe/stripe-go/v72"
	"github.com/xnpltn/hcc/internal/app"
)


func main(){
	f, err := os.Create("logs")
	
	if err!= nil{
		if !errors.Is(err, os.ErrExist){
			log.Println(err)
		}
	}
	defer f.Close()

	err = app.App().Start()
	if err!= nil{
		log.Fatal(err)
	}
}