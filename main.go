package main

import (
	"fmt"
	log "github.com/jeanphorn/log4go"
	"github.com/lishimeng/go-libs/shutdown"
	"github.com/lishimeng/iot-link/internal/etc"
	"github.com/lishimeng/iot-link/internal/setup"
)

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	} ()

	log.LoadConfiguration("log.json")
	defer log.Close()

	// load etc
	etc.SetConfigName("iot-link.toml")
	etc.SetEnvPath([]string{".", "/etc/iot-link"})
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

	shutdown.WaitExit(&shutdown.Configuration{
		BeforeExit: func(s string) {
			log.Info("Shutdown [ %s ] (%s)", etc.Config.Name, s)
		},
	})
}

func setupComponents() (err error) {
	components := []func() error{
		setup.DBRepo,
		setup.Event,
		setup.Message,
		setup.Web,
		setup.Influx,
		setup.DownLink,
		setup.Connector,
	}

	for _, component := range components {
		if err := component(); err != nil {
			return err
		}
	}

	return nil
}
