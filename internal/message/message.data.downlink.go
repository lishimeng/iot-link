package message

import (
	log "github.com/jeanphorn/log4go"
	"github.com/lishimeng/iot-link/internal/codec"
	"github.com/lishimeng/iot-link/internal/codec/intoyun"
	"github.com/lishimeng/iot-link/internal/codec/raw"
	"github.com/lishimeng/iot-link/internal/db/repo"
	"github.com/lishimeng/iot-link/internal/model"
)

func processDownLink(target model.Target, msg map[string]interface{}, onDownLink func(target model.Target, data []byte)) {
	var appConfig, err = repo.GetApp(target.AppId)
	if err != nil {
		// can't find application config
		log.Debug("can't find application[%s] config", target.AppId)
		return
	}

	if appConfig.CodecType != codec.None {

		log.Debug("use downLink Encode:%s", appConfig.CodecType)
		var data []byte
		var err error
		switch appConfig.CodecType {
		case codec.Javascript:
			// find from raw js repo
			data, err = raw.New().Encode(target.AppId, msg)
			break
		case codec.IntoyunTLV:
			// find from tlv repo
			data, err = intoyun.New().Encode(target.AppId, msg)
			break
		default:
			// no codec plugin
			break
		}
		if err != nil {
			log.Debug(err)
			log.Debug("encode err for application %s", target.AppId)
		} else {
			onDownLink(target, data)
		}
	}
}
