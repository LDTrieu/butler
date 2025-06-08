package main

import (
	"butler/application/lib"
	"butler/config"
	_ "butler/pkg/log"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		logrus.Fatalf("Err get config: %v\n", err)
	}
	lib := lib.InitLib(cfg)

	runTestPickPack(lib, cfg)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	fmt.Println("Ctrl + C to exit program")
	<-quit
}

func runTestPickPack(lib *lib.Lib, cfg *config.Config) {}
