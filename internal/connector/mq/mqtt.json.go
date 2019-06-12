package mq

import (
	"encoding/json"
	"fmt"
	log "github.com/jeanphorn/log4go"
	"github.com/lishimeng/iot-link/internal/connector"
	"github.com/lishimeng/iot-link/internal/model"
	"github.com/lishimeng/iot-link/internal/plugin/mqtt"
	"github.com/lishimeng/iot-link/internal/plugin/topics"
)

type connectorMqttJson struct {
	Id               string
	Name             string
	Host             string
	ClientId         string
	DownLinkTopicTpl string
	UpLinkTopicTpl   string
	Session          *mqtt.Session
	Listener         connector.UpLinkListener
	State            bool
	StateDescription string
}

func New(id string, name string, mqttBroker string, mqttClientId string, topicUpLink string, topicDownLink string) connector.Connector {

	log.Debug("create mqtt connector[%d]", mqttBroker)
	c := connectorMqttJson{
		Id: id,
		Name: name,
		Host: mqttBroker,
		ClientId: mqttClientId,
		DownLinkTopicTpl: topicDownLink,
		UpLinkTopicTpl: topicUpLink,
		State: false,
	}

	var onConnect = func(s mqtt.Session) {
		log.Debug("lora mqtt subscribe upLink topics:%s", c.UpLinkTopicTpl)
		c.State = c.Session.State
		c.Session.Subscribe(c.UpLinkTopicTpl)
		c.State = c.Session.State
	}
	var onConnLost = func(s mqtt.Session) {
		log.Debug("lora mqtt lost connection")
		c.Session.State = false
		c.State = c.Session.State
	}
	c.Session = mqtt.CreateSession(c.Host, c.ClientId)

	c.Session.OnConnected = onConnect
	c.Session.OnLostConnect = onConnLost
	c.Session.OnMessage = c.messageCallback

	log.Debug("lora mqtt connect to broker %s", c.Host)
	c.Session.Connect()

	var conn connector.Connector = &c
	return conn
}

func Create(conf connector.Config) (c connector.Connector) {

	c = New(
		conf.ID,
		conf.Name,
		conf.Props["broker"],
		conf.Props["clientId"],
		conf.Props["upLink"],
		conf.Props["downLink"],
	)
	return c
}

func (c connectorMqttJson) GetID() string {
	return c.Id
}

func (c connectorMqttJson) GetState() bool {
	return c.State
}

func (c connectorMqttJson) GetName() string {
	return c.Name
}

func (c *connectorMqttJson) SetListener(listener connector.UpLinkListener) {
	c.Listener = listener
}

// 监听数据上传
///
func(c *connectorMqttJson) messageCallback(mqSession mqtt.Session, topic string, mqttMsg []byte) {

	log.Debug("receive mqtt upLink data %s", topic)
	context := model.LinkMessage{}

	err := resolveMeta(c.UpLinkTopicTpl, topic, &context)
	if err != nil {
		// TODO log
		return
	}
	payload, err := onDataUpLink(mqttMsg)
	if err != nil {
		// TODO log
		return
	}

	// 业务原始数据json格式
	context.Raw = mqttMsg
	// 转换后map格式
	context.Data = payload
	c.Listener(&context)
}

func (c connectorMqttJson) DownLink(target model.Target, logicData []byte) {

	data := string(logicData)

	topic := fmt.Sprintf(c.DownLinkTopicTpl, target.AppId, target.DeviceId)
	c.Session.Publish(topic, data)
}

func onDataUpLink(raw []byte) (payload map[string]interface{}, err error) {

	err = json.Unmarshal(raw, &payload)
	if err != nil {
		return
	}
	return payload, err
}

func resolveMeta(tpl string, topic string, context *model.LinkMessage) (err error) {
	var header map[string]string
	header, err = topics.DeviceUpLinkParamTpl(tpl, topic)
	if err != nil {
		return err
	}
	appId, hasApp := header["ApplicationID"]
	deviceId, hasDev := header["DeviceID"]
	if hasApp && hasDev {
		context.ApplicationID = appId
		context.DeviceID = deviceId
	} else {
		err = fmt.Errorf("topic must contains ApplicationID and DeviceID")
	}
	return err
}