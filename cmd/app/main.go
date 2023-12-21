package main

import (
	"PortfolioV2/internal/app"
	"PortfolioV2/internal/logger"
)

func main() {

	logger.Init()
	app.StartServer()

}
