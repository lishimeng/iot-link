package message

import (
	log "github.com/jeanphorn/log4go"
	"github.com/lishimeng/iot-link/internal/codec"
	"github.com/lishimeng/iot-link/internal/db/repo"
	"github.com/lishimeng/iot-link/internal/downlink"
	"github.com/lishimeng/iot-link/internal/integration/logic"
	"github.com/lishimeng/iot-link/internal/integration/trigger"
	"github.com/lishimeng/iot-link/internal/model"
)

func processUpLink(message *model.LinkMessage) {

	// TODO update device last seen

	err := processDecode(message)
	if err != nil {
		log.Debug(err)
		return
	}

	// run logic script
	*message = processLogic(*message)

	saveMessage(*message)

	// invoke triggers
	processTriggers(*message)

	// active delayed message
	downlink.GetInstance().ActiveDelayedMessage(message.ApplicationID, message.DeviceID)
}

func processDecode(message *model.LinkMessage) (err error) {
	var appConfig repo.AppConfig
	appConfig, err = repo.GetApp(message.ApplicationID)
	if err != nil {
		// can't find application config
		log.Debug("can't find application[%s] config\n", message.ApplicationID)
		return err
	}

	if appConfig.CodecType != codec.None {
		if message.Data == nil { // 需要decode
			message.Data, err = decode(message.ApplicationID, appConfig.CodecType, message.Raw)
		}
	}
	return err
}

func processLogic(message model.LinkMessage) (msg model.LinkMessage) {

	msg = message
	logicScript, err := repo.GetLogic(message.ApplicationID)
	if err == nil {
		logicHandler := logic.New(logicScript.Content)
		log.Debug("handle application[%s] logic", logicScript.AppId)
		msg = logicHandler.OnData(message)
	} else {
		log.Debug(err)
		log.Debug("no application[%s] logic script", message.ApplicationID)
	}
	return msg
}

func processTriggers(message model.LinkMessage) {
	t := *trigger.New()
	eventMessages, err := t.OnData(message)
	if err != nil {
		log.Debug(err)
	} else {
		if len(eventMessages) > 0 {
			h := downlink.GetInstance()// TODO downLink类型/topic转发类型
			for _, eventPayload := range eventMessages {
				e := h.SaveMessage(eventPayload)
				if e != nil {
					log.Debug(e)
				}
			}
		}
	}
}