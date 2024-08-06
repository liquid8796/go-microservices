package main

import (
	"banking/app"
	"banking/logger"
)

func main() {
	// log.Println("Starting our application...")
	logger.Log.Info("Starting our application...")
	app.Start()
}
