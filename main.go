package main

import (
	"github.com/shakilbd009/hexagon-api/app"
	"github.com/shakilbd009/hexagon-api/logger"
)

func main() {
	logger.Info("starting the application")
	app.Start()
}
