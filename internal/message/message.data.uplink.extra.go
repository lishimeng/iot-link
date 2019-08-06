package message

import (
	"github.com/lishimeng/iot-link/internal/etc"
	"github.com/lishimeng/iot-link/internal/integration/persistent"
	"github.com/lishimeng/iot-link/internal/model"
)

func saveMessage(message model.LinkMessage) {
	// persistent data
	if etc.Config.Influx.Enable == 1 {
		tags := map[string]string{
			"applicationID":   message.ApplicationID,
			"applicationName": message.ApplicationName,
			"deviceName":      message.DeviceName,
			"deviceID":        message.DeviceID,
		}
		if len(message.Data) > 0 {
			var fields = make(map[string]interface{})
			for k, v := range message.Data {
				fields[k] = v
			}
			persistent.Save(tags, fields)
		}
	}
}
