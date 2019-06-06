package repo

import (
	"encoding/json"
	"fmt"
	"github.com/lishimeng/iot-link/internal/db"
	"time"
)

func GetConnectorConfig(id string) (c *ConnectorConfig, err error) {

	c = &ConnectorConfig{Id: id}
	err = db.Orm.Context.Read(c)
	return c, err
}

// TODO pagination
func GetConnectConfigs() (cs[]*ConnectorConfig, size int64) {

	q := ConnectorConfig{}
	size, err := db.Orm.Context.QueryTable(&q).OrderBy("CreateTime").All(&cs)
	if err != nil {
		fmt.Println(err)
	}

	return cs, size
}

func CreateConnectorConfig(name string, Type string, propsMap map[string]string) (c ConnectorConfig, err error) {

	props := ""
	bs, err := json.Marshal(propsMap)

	if err != nil {
		return c, err
	}
	props = string(bs)
	c = ConnectorConfig{
		Id: fmt.Sprintf("CONN%d", time.Now().Unix()),
		Name: name,
		Type: Type,
		Props: props,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}
	_, err = db.Orm.Context.InsertOrUpdate(&c, "name")
	err = checkErr(err)
	return c, err
}

func DeleteConnectorConfig(id string) (err error) {

	q := ConnectorConfig{Id: id}
	_, err = db.Orm.Context.Delete(&q)
	return err
}