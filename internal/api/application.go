package api

import (
	"github.com/kataras/iris"
	"github.com/lishimeng/iot-link/internal/db/repo"
)

func SetupApplication(app *iris.Application) {

	routeApplication(app)
}

func routeApplication(app *iris.Application) {

	p := app.Party("api/{ownerId}/application")
	{
		p.Get("/", getApps)
		p.Get("/{appId}/", getApp)
		p.Post("/", createApp)
		p.Post("/{appId}/del", delApp)
	}
}

// app配置
func getApp(ctx iris.Context) {

	appId := ctx.Params().Get("appId")
	app, err := repo.GetApp(appId)

	var res = NewBean()
	if err == nil {
		res.Item = &app
	} else {
		res.Code = -1
	}
	_, _ = ctx.JSON(&res)
}

// app列表
func getApps(ctx iris.Context) {

	res := NewBean()

	pageNo := ctx.URLParamIntDefault("page", 1)
	pageSize := ctx.URLParamIntDefault("size", 5)

	apps, page, err := repo.ListApp(pageNo, pageSize)
	if err != nil || len(apps) == 0 {
		res.Code = -1
	} else {
		res.Item = &apps
		res.Page = &page
	}

	_, _ = ctx.JSON(&res)
}

type AppForm struct {
	AppId string
	Name string
	CodecType string
	ConnectorId string
}

// 创建app
func createApp(ctx iris.Context) {
	res := NewBean()
	var form AppForm
	var app *repo.AppConfig
	var err error
	err = ctx.ReadForm(&form)
	if err == nil {
		app, err = repo.CreateApp(form.AppId, form.Name, form.CodecType, form.ConnectorId)
	}

	if err == nil {
		res.Item = &app
	} else {
		res.Code = -1
	}
	_, _ = ctx.JSON(&res)
}

// 删除app
func delApp(ctx iris.Context) {
	res := NewBean()
	appId := ctx.Params().Get("appId")
	err := repo.DeleteApp(appId)
	if err != nil {
		res.Code = -1
	}
	_, _ = ctx.JSON(&res)
}
