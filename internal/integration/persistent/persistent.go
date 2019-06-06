package persistent

import (
	"github.com/lishimeng/iot-link/internal/etc"
	"github.com/lishimeng/iot-link/internal/plugin/influx"
)

func Save(tags map[string]string, fields map[string]interface{}) {
	go influx.Save(etc.Config.Influx.Host, etc.Config.Influx.Database, "sensor", tags, fields)
}
