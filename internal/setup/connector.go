package setup

import (
	"encoding/json"
	"fmt"
	log "github.com/jeanphorn/log4go"
	"github.com/lishimeng/iot-link/internal/connector"
	"github.com/lishimeng/iot-link/internal/connector/lorawan"
	"github.com/lishimeng/iot-link/internal/connector/mq"
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
		if err != nil {
			return
		}

		forConnector:= connector.Config{
			ID: connConf.Id,
			Name: connConf.Name,
			Type: connConf.Type,
			Props: props,
		}
		var c connector.Connector
		c, err = createConn(forConnector)
		if err != nil {
			return
		}
		messageEngine := message.GetEngine()
		c.SetListener(messageEngine.OnDataUplink)
		ConnectorRepository.Register(c)
	}
}

func createConn(conf connector.Config) (c connector.Connector, err error) {
	switch conf.Type {
	case connector.LoraWanType:
		c, err = lorawan.Create(conf)
		break
	case connector.MqttJson:
		c,err = mq.Create(conf)
		break
	default:
		err = fmt.Errorf("unknown connector type %s", conf.Type)
	}
	return c, err
}

func connExist(id string) bool {
	_, err := ConnectorRepository.GetByID(id)
	// TODO check status
	return err == nil
}