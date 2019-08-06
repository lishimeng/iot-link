package setup

import (
	log "github.com/jeanphorn/log4go"
	"github.com/lishimeng/iot-link/internal/downlink"
	"github.com/lishimeng/iot-link/internal/etc"
	"time"
)

func DownLink() error {
	log.Debug("setup downLink")
	downlink.Init(etc.Config.DownLink.FetchSize, time.Duration(etc.Config.DownLink.IdleTime) * time.Millisecond)
	run()
	return nil
}

func run() {
	go downlink.GetInstance().StartDownLink()
}
