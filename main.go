package main

import (
	"github.com/Somaycon/atelie-api/config"
	"github.com/Somaycon/atelie-api/router"
)

var (
	logger *config.Logger
)

func main() {
	logger := *config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.Errorf("Config Initialization error: %v",err);
		return
	}
	router.Initalize()
}