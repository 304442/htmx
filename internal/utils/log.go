package utils

import (
	"fmt"
	"log"
	"os"

)

func WriteToLogs(v interface{}){
		wkdir, err := os.Getwd()
		if err != nil{
			fmt.Println("error getting working dir: ", err)
		}
		logFile, err := os.OpenFile(fmt.Sprintf("%s/logs/logfile.log", wkdir), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatal(err)
		}
		defer logFile.Close()
		log.SetOutput(logFile)
		switch v.(type){
		case error:
			log.Printf("Error: %v", v)
		default:
			log.Println(v)
		}
}