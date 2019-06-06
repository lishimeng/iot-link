package repo

import (
	"github.com/lishimeng/iot-link/internal/db"
	"time"
)

func GetLogic(appId string) (script LogicScript, err error) {

	script = LogicScript{
		AppId: appId,
	}
	err = db.Orm.Context.Read(&script)
	return script, err
}

func CreateLogic(appId string, content string) (script LogicScript, err error) {

	script, err = GetLogic(appId)
	if err == nil {
		script.Content = content
		script.UpdateTime = time.Now().Unix()
		_, err = db.Orm.Context.Update(&script)
	} else {
		script = LogicScript{AppId: appId, Content: content, CreateTime: time.Now().Unix(), UpdateTime: time.Now().Unix()}
		_, err = db.Orm.Context.Insert(&script)
	}
	return script, err
}

func DeleteLogic(appId string) (err error) {

	q := LogicScript{AppId: appId}
	_, err = db.Orm.Context.Delete(&q)
	return err
}