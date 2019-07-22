package main

import (
	log "github.com/jeanphorn/log4go"
	"github.com/lishimeng/iot-link/internal/etc"
	"github.com/lishimeng/iot-link/internal/setup"
	"github.com/lishimeng/shutdown"
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

	shutdown.WaitExit(&shutdown.Configuration{
		BeforeExit: func(s string) {
			log.Info("Shutdown [ %s ] (%s)", etc.Config.Name, s)
		},
	})
}

func setupComponents() (err error) {
	components := []func() error {
		setup.DBRepo,
		setup.Event,
		setup.Message,
		setup.Connector,
		setup.Web,
		setup.Influx,
	}

	for _, component := range components {
		if err := component(); err != nil {
			return err
		}
	}

	return nil
}

