package setup

import (
	log "github.com/jeanphorn/log4go"
	"github.com/lishimeng/iot-link/internal/message"
	"github.com/lishimeng/iot-link/internal/model"
)

func Message() (err error) {
	message.GetEngine().SetCallback(onDownLink)
	return err
}

func onDownLink(target model.Target, data []byte) {
	conn, err := ConnectorRepository.GetByID(target.ConnectorId)
	if err != nil {
		log.Debug("no connector, skip this data application:%s connector:%s", target.AppId, target.ConnectorId)
	} else {
		conn.DownLink(target, data)
		log.Debug("downLink completed %s:%s:%s", target.ConnectorId, target.AppId, target.DeviceId)
	}
}