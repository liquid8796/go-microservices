package main

import (
	"banking/app"

	"github.com/liquid8796/banking-lib/logger"
)

func main() {
	// log.Println("Starting our application...")
	logger.Info("Starting our application...")
	app.Start()
}
