package repo

import (
	"fmt"
	"github.com/lishimeng/iot-link/internal/db"
	"time"
)

func GetTriggerByApp(appId string) (triggers []TriggerConfig, err error) {

	q := TriggerConfig{AppId: appId}
	res := new([]TriggerConfig)
	_, err = db.Orm.Context.QueryTable(&q).All(res)
	if err == nil {
		triggers = *res
	}
	return triggers, err
}

func GetTrigger(id string) (trigger TriggerConfig, err error) {

	q := TriggerConfig{Id: id}
	err = db.Orm.Context.Read(&q)
	if err == nil {
		trigger = q
	}
	return trigger, err
}

func CreateTrigger(appId string, content string) (trigger TriggerConfig, err error) {

	now := time.Now().Unix()
	id := fmt.Sprintf("%d%s", now, appId)
	trigger = TriggerConfig{
		Id: id,
		AppId:appId,
		Content: content,
		CreateTime: now,
		UpdateTime: now,
		Status: 1,
	}
	_, err = db.Orm.Context.Insert(&trigger)
	err = checkErr(err)
	return trigger, err
}