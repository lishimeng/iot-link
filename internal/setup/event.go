package setup

import (
	"github.com/lishimeng/iot-link/internal/event"
	"github.com/lishimeng/iot-link/internal/message"
	"github.com/lishimeng/iot-link/internal/model"
)

func Event() (err error) {
	event.GetInstance().SetCallback(onEvent)
	return err
}

func onEvent(target model.Target, properties map[string]interface{}) {
	message.GetEngine().OnDataDownLink(target, properties)
}