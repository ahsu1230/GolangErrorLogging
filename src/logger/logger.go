package logger

import (
	"fmt"
	// "os"
	log "github.com/sirupsen/logrus"
)

type Fields = log.Fields

var standardFields Fields

func SetupLogger() {
	log.SetFormatter(&log.TextFormatter{})
	// for production, use JSON instead of text
	// log.SetFormatter(&log.JSONFormatter{})

	// Log to STDOUT (default is STDERR)
	// log.SetOutput(os.Stdout)
	// Or log to File
	// file, err := os.OpenFile("service.log", os.O_APPEND | os.O_CREATE | os.O_RDWR, 0666)
    // if err != nil {
	// 	fmt.Printf("error opening file: %v", err)
    // }
	// log.SetOutput(file)
	// log.Println("Service log re-opened")
	
	// log.SetReportCaller(true) // isn't great if you have a wrapper interface like here

	// log.SetLevel(log.DebugLevel)

	standardFields = Fields{
	// 	"hostAddress": "http://localhost:3000",
		"appName":  "golang-webserver",
	}
}

func Info(message string, fields Fields) {
	log.WithFields(standardFields).
		WithFields(fields).
		Info(message)
}

func Error(message string, err error, fields Fields) {
	log.WithFields(standardFields).
		WithFields(fields).
		Error(fmt.Sprintf("%s (%w)", message, err))
}