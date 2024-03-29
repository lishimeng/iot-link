package lorawan

import (
	"encoding/base64"
	log "github.com/jeanphorn/log4go"
	"github.com/lishimeng/go-connector/lorawan"
	"github.com/lishimeng/iot-link/internal/connector"
	"github.com/lishimeng/iot-link/internal/model"
)

type connectorLoraWan struct {
	Id               string
	Name             string
	Proxy            *lorawan.Connector
	Listener         connector.UpLinkListener
	State            bool
	StateDescription string
}

func New(id string, name string, broker string, clientId string, topicUpLink string, topicDownLink string) (connector.Connector, error) {

	log.Debug("Lorawan connector[%s]", broker)

	c := connectorLoraWan{
		Id:    id,
		Name:  name,
		State: false,
	}

	proxy, err := lorawan.New(broker, clientId, topicUpLink, topicDownLink, 0)
	if err != nil {
		return nil, err
	}
	c.Proxy = proxy
	proxy.SetUpLinkListener(c.onMessage)
	err = proxy.ConnectOnce()
	if err != nil {
		log.Debug(err)
	}

	var conn connector.Connector = &c
	return conn, nil // TODO
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
	return c.Proxy.GetSession().State
}

func (c connectorLoraWan) GetName() string {
	return c.Name
}

func (c *connectorLoraWan) SetListener(listener connector.UpLinkListener) {
	c.Listener = listener
}

// 监听数据上传
///
func (c *connectorLoraWan) onMessage(payload lorawan.PayloadRx) {
	rawData, err := base64.StdEncoding.DecodeString(payload.Data)
	if err != nil {
		return
	}
	context := model.LinkMessage{}
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
}

func (c connectorLoraWan) DownLink(target model.Target, logicData []byte) {
	// 业务数据部分必须为base64格式
	raw := base64.StdEncoding.EncodeToString(logicData)
	log.Debug("lora down link [%s]%s:%s", target.ConnectorId, target.AppId, target.DeviceId)
	log.Debug("data object:%s", logicData)
	log.Debug("data raw:%s", raw)

	downLinkData := lorawan.PayloadTx{FPort: 3, Data: raw}

	go func() {
		err := c.Proxy.DownLink(target.AppId, target.DeviceId, downLinkData)
		if err != nil {
			log.Debug(err)
		}
	}()
}
