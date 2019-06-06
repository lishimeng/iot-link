package mqtt

import (
	Mqtt"github.com/eclipse/paho.mqtt.golang"
	log "github.com/jeanphorn/log4go"
)

type MessageCallback func(session Session, topic string, msg []byte)
type ConnectedCallback func(session Session)
type ConnectionLostCallback func(session Session)

type Session struct {
	client *Mqtt.Client
	OnMessage MessageCallback
	OnConnected ConnectedCallback
	OnLostConnect ConnectionLostCallback
	State bool
	ErrorMessage string
}

func CreateSession(broker string, clientId string) *Session {
	session := Session{}
	opts := Mqtt.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID(clientId)
	opts.SetCleanSession(true)

	opts.SetDefaultPublishHandler(session.MessageHandler)
	opts.SetConnectionLostHandler(session.onConnLost)
	opts.SetOnConnectHandler(session.onConned)

	//opts.SetAutoReconnect(true)

	c := Mqtt.NewClient(opts)
	session.client = &c
	return &session
}

func (session *Session) onConnLost(mqttc Mqtt.Client, reason error) {
	log.Info("mqtt connection error: %s", reason)
	if session.OnLostConnect != nil {
		session.OnLostConnect(*session)
	}
}

func (session *Session) onConned(mqttc Mqtt.Client) {
	log.Debug("mqtt connect success")
	if session.OnConnected != nil {
		session.OnConnected(*session)
	}
}

func (session *Session) Publish(topic string, payload string) () {

	client := *session.client

	log.Debug("publish msg: %s:%s", topic, payload)
	_ = client.Publish(topic, 0, false, payload)
}

func (session *Session) Subscribe(topic string) bool {
	client := *session.client

	log.Debug("Subscribe: %s", topic)
	if token := client.Subscribe(topic, 0, nil); token.Wait() && token.Error() != nil {
		log.Info(token.Error())
		session.ErrorMessage = token.Error().Error()
		session.State = false
		return false
	} else {
		session.State = true
		return true
	}
}

func (session *Session) Unsubscribe(topic string) bool {
	client := *session.client
	if token := client.Unsubscribe(topic); token.Wait() && token.Error() != nil {
		log.Debug(token.Error())
		return false
	} else {
		return true
	}
}

func (session *Session) Connect() () {
	client := *session.client
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		session.ErrorMessage = "connect failed"
	} else {
		session.State = true
		session.ErrorMessage = ""
	}
}

func (session *Session) Close() () {
	client := *session.client
	client.Disconnect(250)
}

func (session *Session) MessageHandler(client Mqtt.Client, msg Mqtt.Message) {
	log.Debug("TOPIC: %s", msg.Topic())
	log.Debug("MSG RAW: %s", msg.Payload())
	if session.OnMessage != nil {
		session.OnMessage(*session, msg.Topic(), msg.Payload())
	}
}