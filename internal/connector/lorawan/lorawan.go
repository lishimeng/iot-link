package lorawan

import (
	"encoding/base64"
	"fmt"
	log "github.com/jeanphorn/log4go"
	"github.com/lishimeng/go-connector/mqtt"
	"github.com/lishimeng/iot-link/internal/connector"
	"github.com/lishimeng/iot-link/internal/model"
)

type connectorLoraWan struct {
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

func New(id string, name string, mqttBroker string, mqttClientId string, topicUpLink string, topicDownLink string) (connector.Connector, error) {

	log.Debug("create mqtt connector[%d]", mqttBroker)
	c := connectorLoraWan{
		Id: id,
		Name: name,
		Host: mqttBroker,
		ClientId: mqttClientId,
		DownLinkTopicTpl: topicDownLink,
		UpLinkTopicTpl: topicUpLink,
		State: false,
	}

	var onConnect = func(s mqtt.Session) {
		log.Debug("lora mqtt subscribe upLink topic:%s", c.UpLinkTopicTpl)
		c.State = c.Session.State
		c.Session.Subscribe(c.UpLinkTopicTpl, 0, nil)
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
	return conn, nil// TODO
}

func Create(conf connector.Config) (c connector.Connector, err error) {

	c, err = New(
		conf.ID,
		conf.Name,
		conf.Props["broker"],
		conf.Props["clientId"],
		conf.Props["upLink"],
		conf.Props["downLink"],
		)
	return c, err
}

func (c connectorLoraWan) GetID() string {
	return c.Id
}

func (c connectorLoraWan) GetState() bool {
	return c.State
}

func (c connectorLoraWan) GetName() string {
	return c.Name
}

func (c *connectorLoraWan) SetListener(listener connector.UpLinkListener) {
	c.Listener = listener
}

// 监听数据上传
///
func(c *connectorLoraWan) messageCallback(mqSession mqtt.Session, topic string, mqttMsg []byte) {

	log.Debug("receive lora upLink data %s", topic)
	context := model.LinkMessage{}
	payload, err := onDataUpLink(mqttMsg)
	if err != nil {
		return
	}

	// lorawan 业务数据base64格式
	rawData, err := base64.StdEncoding.DecodeString(payload.Data)
	if err != nil {
		return
	}
	context.ApplicationID = payload.ApplicationID
	context.DeviceID = payload.DevEUI
	context.ApplicationName = payload.ApplicationName
	context.DeviceName = payload.DeviceName
	context.Raw = rawData

	// 解析object字段
	if payload.DataObj != nil {
		context.Data = *payload.DataObj
	}

	c.Listener(&context)

	go func() {
		// TODO if class_a, handle down link

	} ()
}

func (c connectorLoraWan) DownLink(target model.Target, logicData []byte) {
	// 业务数据部分必须为base64格式
	raw := base64.StdEncoding.EncodeToString(logicData)

	payload := PayloadTx{FPort: 2, Data: raw}
	data := convertJsonDownLinkData(payload) // 序列化得到json类型data字符串

	topic := fmt.Sprintf(c.DownLinkTopicTpl, target.AppId, target.DeviceId)

	go func() {
		err := c.Session.Publish(topic, 0, data)
		if err != nil {
			log.Debug(err)
		}
	}()
}