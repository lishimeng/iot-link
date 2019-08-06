package message

import (
	"encoding/json"
	log "github.com/jeanphorn/log4go"
	"github.com/lishimeng/iot-link/internal/codec"
	"github.com/lishimeng/iot-link/internal/db/repo"
	"github.com/lishimeng/iot-link/internal/model"
)

func processDownLink(target model.Target, msg map[string]interface{}, onDownLink func(model.Target, []byte)) {
	var appConfig, err = repo.GetApp(target.AppId)
	if err != nil {
		// can't find application config
		log.Debug("can't find application[%s] config", target.AppId)
		return
	}

	var data []byte
	if appConfig.CodecType != codec.None {

		log.Debug("use downLink Encode:%s", appConfig.CodecType)

		data, err = encode(target.AppId, appConfig.CodecType, msg)

	} else {
		data, err = json.Marshal(&msg)
	}
	if err != nil {
		log.Debug(err)
		log.Debug("encode err for application %s", target.AppId)
	} else {
		if len(data) > 0 {
			onDownLink(target, data)
		}
	}
}
