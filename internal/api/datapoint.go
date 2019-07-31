package api

import (
	"github.com/kataras/iris"
	"github.com/lishimeng/iot-link/internal/db/repo"
)

func SetupDataPoint(app *iris.Application) {

	routeDataPoint(app)
}

func routeDataPoint(app *iris.Application) {
	p := app.Party("api/{owner}/application/{appId}/dp")

	{
		p.Get("/", getDataPoint)
		p.Post("/", createDataPoint)
		p.Post("/delete", delDataPoint)
	}
}

func getDataPoint(ctx iris.Context) {

	appId := ctx.Params().Get("appId")
	dp, err := repo.GetDataPoint(appId)
	var res = NewBean()
	if err == nil {
		res.Item = &dp
	} else {
		res.Code = -1
	}
	_, _ = ctx.JSON(&res)
}

type DataPointForm struct {
	Content string
}

func createDataPoint(ctx iris.Context) {

	var res = NewBean()
	appId := ctx.Params().Get("appId")
	form := DataPointForm{}
	err := ctx.ReadForm(&form)
	if err != nil {
		res.Code = -1
	} else {
		dp, err := repo.CreateDataPoint(appId, form.Content)
		if err == nil {
			res.Item = &dp
		} else {
			res.Code = -1
		}
	}
	_, _ = ctx.JSON(&res)
}

func delDataPoint(ctx iris.Context) {
	_, _ = ctx.JSON(NewBean())
}
