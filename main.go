package main

import (
	log "github.com/jeanphorn/log4go"
	"github.com/lishimeng/iot-link/internal/etc"
	"github.com/lishimeng/iot-link/internal/setup"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	log.LoadConfiguration("log.json")
	defer log.Close()

	// load etc
	etc.SetConfigName("iot-link.toml")
	etc.AddEnvPath(".")
	etc.AddEnvPath("/etc/iot-link")
	err := etc.LoadEnvs()
	if err != nil {
		log.Info("%s", err)
		return
	}

	err = setupComponents()
	if err != nil {
		log.Info(err)
		return
	}

	waitExit()
}

func waitExit() {

	sigChan := make(chan os.Signal)
	exitChan := make(chan struct{})
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		// todo: handle graceful shutdown?
		//exitChan <- struct{}{}
		log.Info("stopping server[auto]")
	}()
	log.Info("wait for exit")
	select {
	case <-exitChan:
		log.Info("exit received, stopping immediately")
	case s := <-sigChan:
		log.Info("signal received, stopping immediately %s", s)
	}
}

func setupComponents() (err error) {
	components := []func() error {
		setup.DBRepo,
		setup.Event,
		setup.Message,
		setup.Connector,
		setup.Web,
	}

	for _, component := range components {
		if err := component(); err != nil {
			return err
		}
	}

	return nil
}

