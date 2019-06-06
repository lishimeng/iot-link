package setup

import (
	"encoding/json"
	log "github.com/jeanphorn/log4go"
	"github.com/lishimeng/iot-link/internal/connector"
	"github.com/lishimeng/iot-link/internal/connector/lorawan"
	"github.com/lishimeng/iot-link/internal/db/repo"
	"github.com/lishimeng/iot-link/internal/message"
	"time"
)

var ConnectorRepository connector.Repository

func Connector() error {
	ConnectorRepository = connector.New()

	go loadConnectors()
	return nil
}

func loadConnectors() {

	var _initAllConnector = func() {

		log.Debug("check connectors")
		// get all config
		configs, size := repo.GetConnectConfigs()
		if size > 0 {
			// loop config to connect message platform
			log.Debug("connector size:%d", size)
			for _, config := range configs {
				loadConnector(*config)
			}
		}
	}

	for {
		_initAllConnector()
		time.Sleep(time.Second * 10)
		break
	}
}

func loadConnector(connConf repo.ConnectorConfig) {

	if !connExist(connConf.Name) {
		log.Debug("load connector[%s]", connConf.Name)
		var props map[string]string
		err := json.Unmarshal([]byte(connConf.Props), &props)
		if err == nil {
			forConnector:= connector.Config{
				ID: connConf.Id,
				Name: connConf.Name,
				Type: connConf.Type,
				Props: props,
			}
			c := lorawan.Create(forConnector)
			messageEngine := message.GetEngine()
			c.SetListener(messageEngine.OnDataUplink)
			ConnectorRepository.Register(c)
		}
	}
}

func connExist(id string) bool {
	_, err := ConnectorRepository.GetByID(id)
	// TODO check status
	return err == nil
}