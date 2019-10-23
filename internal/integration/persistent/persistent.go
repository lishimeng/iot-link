package persistent

import (
	log "github.com/jeanphorn/log4go"
	"github.com/lishimeng/go-connector/influx"
	"github.com/lishimeng/iot-link/internal/etc"
)

var influxClient *influx.Connector

func Init() (err error) {

	if etc.Config.Influx.Enable == 1 {
		influxClient, err = influx.New(etc.Config.Influx.Host)
	}
	return err
}

func Save(tags map[string]string, fields map[string]interface{}) {

	go func() {
		if influxClient != nil {
			err := influxClient.Save(etc.Config.Influx.Database, "sensor", tags, fields)
			if err != nil {
				log.Debug(err)
			}
		}
	}()
}
