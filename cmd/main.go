package main

import (
	"log"
	"github.com/xnpltn/hcc/internal/app"
)


func main(){
	
	err := app.App().Start()
	if err!= nil{
		log.Fatal(err)
	}
}