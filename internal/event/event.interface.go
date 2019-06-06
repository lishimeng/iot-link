package event

import (
	"github.com/lishimeng/iot-link/internal/model"
)

type Callback func(target model.Target, properties map[string]interface{})
type DownLinkHandler interface {
	Send(target model.Target, properties map[string]interface{})
	SetCallback(callback Callback)
}

var singleton DownLinkHandler

func GetInstance() DownLinkHandler {
	return singleton
}

func New() DownLinkHandler {
	h := sendHandler{}
	var sh DownLinkHandler = &h
	return sh
}

func init() {
	singleton = New()
}

type sendHandler struct {
	cb Callback
}

func (h sendHandler) Send(target model.Target, properties map[string]interface{}) {
	// app config
	// process down link
	if h.cb != nil {
		h.cb(target, properties)
	}
}

func (h *sendHandler) SetCallback(callback Callback) {
	h.cb = callback
}