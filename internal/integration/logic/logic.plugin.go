package logic

import (
	"fmt"
	log "github.com/jeanphorn/log4go"
	"github.com/lishimeng/go-libs/script"
	"github.com/lishimeng/iot-link/internal/downlink"
	"github.com/lishimeng/iot-link/internal/model"
	"github.com/robertkrimen/otto"
)

const Tpl = `
// object -> byte array
function execute(message){
	if (message.data) {
		return message.data;
	} else {
		return {};
	}
}
`

// 收到业务数据后执行
type MessageLogicHandler interface {
	OnData(msg model.LinkMessage) model.LinkMessage
}

func New(logicScript string) MessageLogicHandler {

	logicHandler := msgLogicHandler{
		script: logicScript,
	}
	var h = &logicHandler
	return h
}

type msgLogicHandler struct {
	script string
}

func (handler msgLogicHandler) OnData(msg model.LinkMessage) model.LinkMessage {
	logicContent := handler.script
	engine, err := script.Create(logicContent)
	if err != nil {
		log.Debug(err)
		log.Debug("create logic vm failed")
		return msg
	}
	engine.Inject("event", callback)
	var result otto.Value
	result, err = engine.Invoke("execute", msg.Data)
	if err != nil {
		log.Debug(err)
		log.Debug("invoke logic script failed")
	} else {
		value, ok := exportValue(result)
		if ok {
			msg.Data = value
		} else {
			// logic执行异常不清除data内容
			//msg.Data = make(map[string]interface{})
		}
	}
	return msg
}

func callback(call otto.FunctionCall) otto.Value {
	defer func() {
		if exp := recover(); exp != nil {
			log.Debug("%s", exp)
		}
	}()
	var result = otto.Value{}
	if len(call.ArgumentList) != 2 {
		return result
	}
	targetValue := call.Argument(0)
	dataValue := call.Argument(1)
	target, hasTarget := exportValue(targetValue)

	if !hasTarget {
		log.Debug("target err")
		return result
	}

	data, hasData := exportValue(dataValue)

	if !hasData {
		log.Debug("data err")
		return result
	}

	appID, okAppID := target["applicationID"]
	deviceID, okDeviceID := target["deviceID"]
	if okAppID && okDeviceID {
		go _cb(model.Target{
			AppId:    appID.(string),
			DeviceId: deviceID.(string),
		}, data, false, false)// 设定即时消息
	}
	return result
}

func _cb(target model.Target, data map[string]interface{}, delayed bool, auto bool) {

	log.Debug("send lora downLink [%s:%s:%s] data:%s", target.ConnectorId, target.AppId, target.DeviceId, data)
	payload := model.EventPayload{
		Receiver: target,
		Data: data,
		Delayed: delayed,
		AutoDownLink: auto,
	}
	h := downlink.GetInstance()
	var handler = *h
	_ = handler.SaveMessage(payload)
}

func exportValue(value otto.Value) (target map[string]interface{}, ok bool) {

	v, err := value.Export()
	if err != nil {
		fmt.Println(err)
		ok = false
		return target, ok
	}
	switch v.(type) {
	case map[string]interface{}:
		target = v.(map[string]interface{})
		ok = true
		break
	default:
		ok = false
		break
	}
	return target, ok
}
