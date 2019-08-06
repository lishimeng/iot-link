package downlink

import (
	"encoding/json"
	"fmt"
	log "github.com/jeanphorn/log4go"
	"github.com/lishimeng/iot-link/internal/db/repo"
	"github.com/lishimeng/iot-link/internal/model"
)

func (m msgEngine) SaveMessage(eventPayload model.EventPayload) (err error) {

	// save to downLink storage or delayed downLink storage
	log.Debug("save message")
	datas, err := json.Marshal(&eventPayload.Data)
	if err != nil {
		return err
	}
	payload := string(datas)
	if eventPayload.Delayed {
		delayedType := 0
		if eventPayload.AutoDownLink {
			delayedType = 1
		}
		_, err = repo.SaveDelayedDownLinkMessage(eventPayload.Receiver.AppId, eventPayload.Receiver.DeviceId, delayedType, payload)
	} else {
		_, err = repo.SaveDownLinkMessage(eventPayload.Receiver.AppId, eventPayload.Receiver.DeviceId, payload)
	}
	return err
}

func (m msgEngine) ActiveDelayedMessage(appId string, deviceId string) {

	// Find by device & application.If the message is exist, active it(move to downLink storage).
	var forDownLink repo.DownLinkData
	msg, err := repo.GetDelayedMessage(appId, deviceId, repo.Push)
	if err == nil {
		if msg.Status == repo.Enabled {
			// TODO
			// copy to downLink storage
			forDownLink, err = repo.SaveDownLinkMessage(msg.AppId, msg.DeviceId, msg.Payload)

			if err == nil {
				log.Debug(forDownLink.Id)
				// change to disabled
				err = repo.DeleteDelayedDownLinkMessage(msg.Id)
			}

		} else {
			// TODO
		}
	}
}

func (m msgEngine) PullDelayedMessage(appId string, deviceId string) (eventPayload model.EventPayload, err error) {

	msg, err := repo.GetDelayedMessage(appId, deviceId, repo.Pull)
	if err != nil {
		return eventPayload, err
	}
	log.Debug("delayed message:%s, %d, %d", msg.Id, msg.Status, msg.Type)
	if msg.Status == repo.Enabled && msg.Type == repo.Pull {
		var data = make(map[string]interface{})
		err = json.Unmarshal([]byte(msg.Payload), &data)
		if err != nil {
			return eventPayload, err
		}
		eventPayload = model.EventPayload{
			Receiver: model.Target{AppId: msg.AppId, DeviceId: msg.DeviceId},
			Data: data,
			MessageTime: msg.CreateTime,
		}
		return eventPayload, err
	} else {
		return eventPayload, fmt.Errorf("message is invalid")
	}
}

func (m msgEngine) CloseExpiredDelayedMessages() {

	// TODO 查找过期的消息并删除
}
