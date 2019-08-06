package trigger

import (
	"encoding/json"
	log "github.com/jeanphorn/log4go"
	"github.com/lishimeng/iot-link/internal/db/repo"
	"github.com/lishimeng/iot-link/internal/model"
)

type EventProducer interface {
	OnData(msg model.LinkMessage) ([]model.EventPayload, error)
}

type eventProducer struct {
}

func New() *EventProducer {
	p := eventProducer{}
	var h EventProducer = &p
	return &h
}

func (e eventProducer) OnData(msg model.LinkMessage) (datas []model.EventPayload, err error) {

	// find trigger conditions
	triggers, err := repo.GetTriggerByApp(msg.ApplicationID)

	if err == nil && len(triggers) > 0 {
		var tmp = make([]*model.EventPayload, len(triggers))
		var i = 0
		for _, triggerConfig := range triggers {
			var data *model.EventPayload
			data, err = onData(triggerConfig, msg)
			if err != nil {
				log.Debug(err)
			} else {
				if data != nil {
					tmp[i] = data
					i++
				}
			}
		}
		datas = make([]model.EventPayload, i)
		if i > 0 {
			for j := 0; j < i; j++ {
				datas[j] = *tmp[j]
			}
		}
		// TODO
	}
	return datas, err
}

func onData(triggerConfig repo.TriggerConfig, msg model.LinkMessage) (payload *model.EventPayload, err error) {
	t := model.Trigger{}
	err = json.Unmarshal([]byte(triggerConfig.Content), &t)
	if err != nil {
		return payload, err
	}
	var ok bool
	ok, err = handleTrigger(t, msg)

	if err == nil {
		//log.Debug("trigger result %b", ok)
		if ok {
			// 条件触发器在本数据中命中,payload为响应数据
			log.Debug("trigger hit the target")
			payload = handleEvent(t, msg)
		}
	}
	return payload, err
}

func handleEvent(t model.Trigger, _ model.LinkMessage) (payload *model.EventPayload) {
	target := t.TargetEvent
	if target != nil {
		payload = target
	}
	return payload
}

func handleTrigger(t model.Trigger, msg model.LinkMessage) (b bool, err error) {

	javascript := buildTriggerContent(t, msg)
	b, err = calcTrigger(javascript)
	return b, err
}
