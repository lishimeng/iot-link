package repo

import (
	"encoding/json"
	"fmt"
	"github.com/lishimeng/iot-link/internal/db"
	"time"
)

func GetConnectorConfig(id string) (c ConnectorConfig, err error) {

	c = ConnectorConfig{Id: id}
	err = db.Orm.Context.Read(&c)
	return c, err
}

// TODO pagination

func ListConnector(pageNum int, limit int) (conns []ConnectorConfig, page Page, err error) {

	q := ConnectorConfig{}
	res := new([]ConnectorConfig)
	var size int64
	size, err = db.Orm.Context.QueryTable(&q).Count()

	if err != nil {
		return conns, page, err
	}
	count := int(size)
	page = NewPage(count, pageNum, limit)
	if count > 0 {
		limit, start := GetLimit(pageNum, limit)
		_, err = db.Orm.Context.QueryTable(&q).Limit(limit, start).OrderBy("CreateTime").All(res)
		if err == nil {
			conns = *res
		}
	}
	return conns, page, err
}

func GetConnectConfigs() (cs []*ConnectorConfig, size int64) {

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
	return CreateConnectorConf(name, Type, props)
}

func CreateConnectorConf(name string, Type string, props string) (c ConnectorConfig, err error) {

	c = ConnectorConfig{
		Id:         fmt.Sprintf("CONN%d", time.Now().Unix()),
		Name:       name,
		Type:       Type,
		Props:      props,
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
