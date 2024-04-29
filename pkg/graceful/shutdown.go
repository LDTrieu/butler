package graceful

import (
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/sirupsen/logrus"
)

var (
	stop sync.WaitGroup
)

func init() {
	stop.Add(1)
	go func() {
		q := make(chan os.Signal, 1)
		signal.Notify(q, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT)
		logrus.Debug("receive signal: ", <-q)
		stop.Done()
	}()
}

func ShutDown() {
	stop.Wait()
	logrus.Debug("Shutdown")
	os.Exit(0)
}
