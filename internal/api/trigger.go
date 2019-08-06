package api

import (
	"github.com/kataras/iris"
	"github.com/lishimeng/iot-link/internal/db/repo"
)

func SetupTrigger(app *iris.Application) {
	routTrigger(app)
}

func routTrigger(app *iris.Application) {

	p := app.Party("api/{owner}/application/{appId}/trigger")
	{
		p.Get("", listTriggers)
		p.Get("/{id}", getTrigger)
		//p.Get("/tpl", triggerTemplate)
		p.Post("/", createTrigger)
		p.Post("/del/{id}", delTrigger)
	}
}

func listTriggers(ctx iris.Context) {
	res := NewBean()
	appId := ctx.Params().Get("appId")
	triggers, err := repo.GetTriggerByApp(appId)
	if err != nil {
		res.Code = -1
	} else {
		res.Item = &triggers
	}

	_, _ = ctx.JSON(&res)
}

func getTrigger(ctx iris.Context) {
	res := NewBean()
	appId := ctx.Params().Get("appId")
	triggerId := ctx.Params().Get("id")
	triggerConfig, err := repo.GetTrigger(triggerId)
	if err != nil {
		res.Code = -1
	} else {
		if triggerConfig.AppId == appId {
			res.Item = &triggerConfig
		} else {
			res.Code = -1
		}
	}

	_, _ = ctx.JSON(&res)
}

type TriggerForm struct {
	Content string
}

// update or create
func createTrigger(ctx iris.Context) {

	res := NewBean()
	appId := ctx.Params().Get("appId")
	form := TriggerForm{}
	err := ctx.ReadForm(&form)
	if err == nil {

	}
	var triggerConfig repo.TriggerConfig
	triggerConfig, err = repo.CreateTrigger(appId, form.Content)
	if err == nil {
		res.Item = &triggerConfig
	}

	if err != nil {
		res.Code = -1
	}
	_, _ = ctx.JSON(&res)
}

func delTrigger(ctx iris.Context) {

	res := NewBean()
	//appId := ctx.Params().Get("appId")
	// TODO check app_id

	_, _ = ctx.JSON(&res)
}
