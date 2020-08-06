package logger

import (
	"fmt"
	"os"
	log "github.com/sirupsen/logrus"
)

type LogFields = log.Fields

var standardFields LogFields

func SetupLogger() {
	log.SetFormatter(&log.TextFormatter{})
	// for production, use JSON instead of text
	log.SetFormatter(&log.JSONFormatter{})

	// log.SetOutput(os.Stdout)
	// Or log to File
	file, err := os.OpenFile("service.log", os.O_APPEND | os.O_CREATE | os.O_RDWR, 0666)
    if err != nil {
		fmt.Printf("error opening file: %v", err)
    }
	log.SetOutput(file)
	log.Println("Service log re-opened")
	
	// log.SetLevel(log.DebugLevel)

	standardFields = LogFields{
		"hostAddress": "http://localhost:3000",
		"appName":  "golang-webserver-tutorial",
	}
}

func LogInfo(message string, fields LogFields) {
	log.WithFields(standardFields).
		WithFields(fields).
		Info(message)
}

func LogError(message string, err error, fields LogFields) {
	log.WithFields(standardFields).
		WithFields(fields).
		Error(fmt.Sprintf("%s (%w)", message, err))
}