package api

import (
	"github.com/kataras/iris"
	"github.com/lishimeng/iot-link/internal/db/repo"
	"github.com/lishimeng/iot-link/internal/integration/logic"
)

func SetupLogic(app *iris.Application) {
	routLogic(app)
}

func routLogic(app *iris.Application) {

	p := app.Party("api/{owner}/application/{appId}/logic")
	{
		p.Get("", getLogic)
		p.Get("/tpl", logicTemplate)
		p.Post("/", updateLogic)
		p.Post("/del", delLogic)
	}
}

func logicTemplate(ctx iris.Context) {
	res := NewBean()

	item := repo.LogicScript{
		Content: logic.Tpl,
	}
	res.Item = &item

	_, _ = ctx.JSON(&res)
}

func getLogic(ctx iris.Context) {
	res := NewBean()
	appId := ctx.Params().Get("appId")
	logicScript, err := repo.GetLogic(appId)
	if err != nil {
		res.Code = -1
	} else {
		res.Item = &logicScript
	}

	_, _ = ctx.JSON(&res)
}

type LogicForm struct {
	Content string
}

// update or create
func updateLogic(ctx iris.Context) {

	res := NewBean()
	appId := ctx.Params().Get("appId")
	form := LogicForm{}
	err := ctx.ReadForm(&form)
	if err == nil {
		content := ctx.Params().Get("content")
		var logicScript repo.LogicScript
		logicScript, err = repo.CreateLogic(appId, content)
		if err == nil {
			res.Item = &logicScript
		}
	}

	if err != nil {
		res.Code = -1
	}
	_, _ = ctx.JSON(&res)
}

func delLogic(ctx iris.Context) {

	res := NewBean()
	appId := ctx.Params().Get("appId")

	err := repo.DeleteLogic(appId)
	if err != nil {
		res.Code = -1
	}
	_, _ = ctx.JSON(&res)
}
