package influx

import (
	"github.com/influxdata/influxdb1-client"
	log "github.com/jeanphorn/log4go"
	"net/url"
	"time"
)

var failTimes = 0

type Config struct {
	Host string
	Database string
}

func Save(hostName string, database string, measurement string, tags map[string]string, fields map[string]interface{}) {

	log.Debug("influx save: %s[%s:%s] tags:%s --> values:%s", hostName, database, measurement, tags, fields)
	host, err := url.Parse(hostName)
	if err != nil {
		log.Info(err)
		return
	}

	con, err := client.NewClient(client.Config{URL: *host})
	if err != nil {
		log.Info(err)
		return
	}

	point := client.Point{
		Measurement:measurement,
		Tags: tags,
		Fields: fields,
		Time: time.Now(),
	}

	points := []client.Point{point}

	bps := client.BatchPoints{
		Points: points,
		Database: database,
	}

	_, err = con.Write(bps)

	if err != nil {
		log.Info(err)
		failTimes++
	} else {
		failTimes = 0
	}

	if failTimes > 0x10 {
		panic("influx failed too much times")
	}
}