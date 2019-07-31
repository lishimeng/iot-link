package web

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/lishimeng/iot-link/internal/etc"
	"github.com/lishimeng/iot-link/internal/static"
)

func Run(components ...func(app *iris.Application)) {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(logger.New())
	app.Favicon("./favicon.ico", "/favicon.ico")

	app.OnErrorCode(404, func(ctx iris.Context) {
		_, _ = ctx.Writef("404[not found]")
	})

	bs, err := static.Asset("index.html")
	indexHtml := ""
	if err == nil {
		indexHtml = string(bs)
	}

	app.StaticEmbedded("/", "", static.Asset, static.AssetNames)
	app.Get("/", func(c iris.Context) {
		_, _ = c.HTML(indexHtml)
	})

	if len(components) > 0 {
		for _, component := range components {
			component(app)
		}
	}

	_ = app.Run(iris.Addr(etc.Config.Web.Listen))

}
