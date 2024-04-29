package main

import (
	"butler/application/server"
	"butler/config"
	"butler/pkg/graceful"
	_ "butler/pkg/log"

	"github.com/sirupsen/logrus"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		logrus.Fatalf("Err get config: %v\n", err)
	}

	server := server.NewServer(cfg)
	server.Start()
	defer server.Stop()

	logrus.Info("Program is now running. Press CTRL-C to exit.")
	graceful.ShutDown()
}
