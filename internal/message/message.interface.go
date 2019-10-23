package message

import (
	"github.com/lishimeng/iot-link/internal/model"
)

type DataProcessEngine interface {
	OnDataUpLink(upLink *model.LinkMessage)
	OnDataDownLink(target model.Target, props map[string]interface{})
	SetCallback(cb OnDownLink)
}

type OnDownLink func(target model.Target, data []byte)

type dataProcessEngineImpl struct {
	cb OnDownLink
}

var singleton DataProcessEngine

func init() {
	d := dataProcessEngineImpl{}
	singleton = &d
}

func GetEngine() DataProcessEngine {
	return singleton
}

func (d *dataProcessEngineImpl) SetCallback(cb OnDownLink) {
	d.cb = cb
}

func (d dataProcessEngineImpl) OnDataUpLink(upLink *model.LinkMessage) {
	processUpLink(upLink)
}

func (d dataProcessEngineImpl) OnDataDownLink(target model.Target, props map[string]interface{}) {
	processDownLink(target, props, d.cb)
}
