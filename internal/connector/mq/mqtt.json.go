package mq

import (
	"encoding/json"
	"fmt"
	log "github.com/jeanphorn/log4go"
	"github.com/lishimeng/go-connector/mqtt"
	"github.com/lishimeng/iot-link/internal/connector"
	"github.com/lishimeng/iot-link/internal/model"
	"github.com/lishimeng/iot-link/internal/plugin/topics"
)

type mqttJsonConnector struct {
	Id               string
	Name             string
	Host             string
	ClientId         string
	DownLinkTopic    string
	UpLinkTopicTpl   string
	UpLinkTopic      string
	Session          *mqtt.Session
	Listener         connector.UpLinkListener
	State            bool
	StateDescription string
}

func New(id string, name string, mqttBroker string, mqttClientId string, topicUpLinkTpl string, topicUpLink string, topicDownLink string) (connector.Connector, error) {

	log.Debug("create mqtt connector[%s]", mqttBroker)
	c := mqttJsonConnector{
		Id:               id,
		Name:             name,
		Host:             mqttBroker,
		ClientId:         mqttClientId,
		DownLinkTopic:    topicDownLink,
		UpLinkTopic:      topicUpLink,
		UpLinkTopicTpl:   topicUpLinkTpl,
		State:            false,
	}

	var onConnect = func(s mqtt.Session) {
		log.Debug("mqtt connected")
		c.State = c.Session.State
		if len(c.UpLinkTopic) > 0 {
			log.Debug("mqtt subscribe upLink topics:%s", c.UpLinkTopic)
			c.Session.Subscribe(c.UpLinkTopic, 0, nil)
		}
	}
	var onConnLost = func(s mqtt.Session, reason error) {
		log.Debug("mqtt lost connection")
		log.Debug(reason)
		c.Session.State = false
		c.State = c.Session.State
	}
	c.Session = mqtt.CreateSession(c.Host, c.ClientId)

	c.Session.OnConnected = onConnect
	c.Session.OnLostConnect = onConnLost
	c.Session.OnMessage = c.messageCallback

	log.Debug("mqtt connect to broker %s", c.Host)
	c.Session.Connect()

	var conn connector.Connector = &c
	return conn, nil // TODO
}

func Create(conf connector.Config) (c connector.Connector, err error) {

	c, err = New(
		conf.ID,
		conf.Name,
		conf.Props["broker"],
		conf.Props["clientId"],
		conf.Props["upLinkTpl"],
		conf.Props["upLink"],
		conf.Props["downLink"],
	)
	return c, err
}

func (c mqttJsonConnector) GetID() string {
	return c.Id
}

func (c mqttJsonConnector) GetState() bool {
	return c.State
}

func (c mqttJsonConnector) GetName() string {
	return c.Name
}

func (c *mqttJsonConnector) SetListener(listener connector.UpLinkListener) {
	c.Listener = listener
}

// 监听数据上传
///
func (c *mqttJsonConnector) messageCallback(mqSession mqtt.Session, topic string, mqttMsg []byte) {

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

func (c mqttJsonConnector) DownLink(target model.Target, logicData []byte) {

	data := string(logicData)

	if len(c.DownLinkTopic) > 0 {
		topic := fmt.Sprintf(c.DownLinkTopic, target.AppId, target.DeviceId)
		log.Debug("down link: %s[%s]", data, topic)
		go func() {
			err := c.Session.Publish(topic, 0, data)
			if err != nil {
				log.Debug(err)
			}
		}()
	}
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
