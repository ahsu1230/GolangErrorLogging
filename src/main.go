package main

import (
	"github.com/ahsu1230/golangwebservertutorial/src/router"
	"github.com/ahsu1230/golangwebservertutorial/src/logger"
)

func main() {
	logger.SetupLogger()
	engine := router.Setup()
    engine.Run(":3000")
}