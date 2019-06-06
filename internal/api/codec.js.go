package api

import (
	"github.com/kataras/iris"
	"github.com/lishimeng/iot-link/internal/db/repo"
)

func SetupCodecJs(app *iris.Application) {
	routeCodecJs(app)
}

func routeCodecJs(app *iris.Application) {

	p := app.Party("api/{owner}/application/{appId}/codec/js/")
	{
		p.Get("/", getCodecJs)
		p.Post("/", createOrUpdateCodecJs)
		p.Post("/del", delCodecJs)
	}
}

func getCodecJs(ctx iris.Context) {

	res := NewBean()
	appId := ctx.Params().Get("appId")
	js, err := repo.GetJs(appId)
	if err != nil {
		res.Code = -1
	} else {
		res.Item = &js
	}

	_, _ = ctx.JSON(&res)
}

type CodecJsForm struct {
	EncodeContent string
	DecodeContent string
}

func createOrUpdateCodecJs(ctx iris.Context) {

	res := NewBean()
	appId := ctx.Params().Get("appId")
	form := CodecJsForm{}
	err := ctx.ReadForm(&form)
	if err == nil {
		var js repo.CodecScript
		js, err = repo.CreateOrUpdateJs(appId, form.EncodeContent, form.DecodeContent)
		if err == nil {
			res.Item = &js
		}
	}

	if err != nil {
		res.Code = -1
	}
	_, _ = ctx.JSON(&res)
}

func delCodecJs(ctx iris.Context) {

	res := NewBean()
	appId := ctx.Params().Get("appId")
	err := repo.DeleteJs(appId)
	if err != nil {
		res.Code = -1
	}

	_, _ = ctx.JSON(&res)
}