package connector

import (
	"fmt"
	"github.com/lishimeng/iot-link/internal/model"
)

const (
	// lorawan
	LoraWanType = "lora"

	// mqtt传输,json数据
	MqttJson = "mqtt_json"
)

type Config struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Props map[string]string `json:"props"`
}

type UpLinkListener func(context *model.LinkMessage)

type Connector interface {
	GetID() string
	GetName() string
	GetState() bool
	DownLink(target model.Target, data []byte)
	SetListener(OnDataUplink UpLinkListener)
}

type Repository interface {
	GetByName(name string) (Connector, error)
	GetByID(id string) (Connector, error)
	Register(c Connector)
}

type connRepo struct {
	connectors map[string]Connector
	name2id map[string]string
}

func New() Repository {
	r := connRepo{
		connectors: make(map[string]Connector),
		name2id: make(map[string]string),
	}
	var repo Repository = &r

	return repo
}

func (r connRepo) Register(c Connector) {
	id := c.GetID()
	name := c.GetName()
	r.connectors[id] = c
	r.name2id[name] = id
}

func (r connRepo) GetByID(id string) (c Connector, err error) {
	c, ok := r.connectors[id]
	if !ok {
		err = fmt.Errorf("no such connector id[%s]", id)
	}
	return c, err
}

func (r connRepo) GetByName(name string) (c Connector, err error) {
	id, ok := r.name2id[name]
	if !ok {
		err = fmt.Errorf("no such connector name[%s]", name)
	} else {
		c, err = r.GetByID(id)
		if err != nil {
			err = fmt.Errorf("no such connector name[%s]", name)
		}
	}
	return c, err
}
