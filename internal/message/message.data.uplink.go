package message

import (
	log "github.com/jeanphorn/log4go"
	"github.com/lishimeng/iot-link/internal/codec"
	"github.com/lishimeng/iot-link/internal/codec/intoyun"
	"github.com/lishimeng/iot-link/internal/codec/raw"
	"github.com/lishimeng/iot-link/internal/db/repo"
	"github.com/lishimeng/iot-link/internal/integration/logic"
	"github.com/lishimeng/iot-link/internal/integration/persistent"
	"github.com/lishimeng/iot-link/internal/model"
)

func processUpLink(message *model.LinkMessage) {

	var appConfig, err = repo.GetApp(message.ApplicationID)
	if err != nil {
		// can't find application config
		log.Debug("can't find application[%s] config\n", message.ApplicationID)
		return
	}

	if appConfig.CodecType != codec.None {
		if message.Data == nil { // 需要decode
			switch appConfig.CodecType {
			case codec.Javascript:
				// find from raw js repo
				message.Data, err = raw.New().Decode(message.ApplicationID, message.Raw)
				break
			case codec.IntoyunTLV:
				// find from tlv repo
				message.Data, err = intoyun.New().Decode(message.ApplicationID, message.Raw)
				break
			default:
				// no codec plugin
				break
			}
		}
	}
	if err != nil {
		log.Debug(err)
		return
	}

	// run logic script
	logicScript, err := repo.GetLogic(message.ApplicationID)
	if err == nil {
		logicHandler := logic.New(logicScript.Content)
		log.Debug("handle application[%s] logic", logicScript.AppId)
		*message = logicHandler.OnData(*message)
	} else {
		log.Debug(err)
		log.Debug("no application[%s] logic script", message.ApplicationID)
	}

	saveMessage(message)
}

func saveMessage(message *model.LinkMessage) {
	// persistent data
	tags := map[string]string{
		"applicationID":   message.ApplicationID,
		"applicationName": message.ApplicationName,
		"deviceName":      message.DeviceName,
		"deviceID":        message.DeviceID,
	}
	if len(message.Data) > 0 {
		log.Debug(message.Data)
		var fields = make(map[string]interface{})
		for k, v := range message.Data {
			fields[k] = v
		}
		persistent.Save(tags, fields)
	}
}
