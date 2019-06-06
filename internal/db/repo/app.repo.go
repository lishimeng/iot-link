package repo

import (
	"github.com/lishimeng/iot-link/internal/db"
	"time"
)

func GetApp(appId string) (app AppConfig, err error) {
	app = AppConfig{AppId: appId}
	err = db.Orm.Context.Read(&app)
	return app, err
}

func ListApp(pageNum int, limit int) (apps []AppConfig, page Page, err error) {

	q := AppConfig{}
	res := new([]AppConfig)
	var size int64
	size, err = db.Orm.Context.QueryTable(&q).Count()

	if err != nil {
		return apps, page, err
	}
	count := int(size)
	page = NewPage(count, pageNum, limit)
	if count > 0 {
		limit, start := GetLimit(pageNum, limit)
		_, err = db.Orm.Context.QueryTable(&q).Limit(limit, start).OrderBy("CreateTime").All(res)
		if err == nil {
			apps = *res
		}
	}
	return apps, page, err
}

func CreateApp(appId string, name string, codecType string, connectorId string) (app *AppConfig, err error) {

	app = &AppConfig{
		AppId: appId,
		AppDescription: name,
		CodecType: codecType,
		Connector: connectorId,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}
	_, err = db.Orm.Context.Insert(app)
	err = checkErr(err)
	return app, err
}

func UpdateApp(appId string, name string, codecType string, connectorId string) (app AppConfig, err error) {
	app = AppConfig{
		AppId: appId,
	}
	app.UpdateTime = time.Now().Unix()
	app.Connector = connectorId
	app.CodecType = codecType
	app.AppDescription = name
	_, err = db.Orm.Context.Update(&app, "AppDescription", "CodecType", "Connector", "UpdateTime")
	return app, err
}

func DeleteApp(appId string) (err error) {

	q := AppConfig{AppId: appId}
	_, err = db.Orm.Context.Delete(&q)
	return err
}