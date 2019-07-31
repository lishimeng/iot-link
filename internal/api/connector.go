package api

import (
	"github.com/kataras/iris"
	"github.com/lishimeng/iot-link/internal/db/repo"
)

func SetupConnector(app *iris.Application) {
	routeConnector(app)
}

func routeConnector(app *iris.Application) {

	p := app.Party("api/{owner}/connector")

	{
		p.Get("/", getConnectors)
		p.Get("{connectorId}/", getConnector)
		p.Post("/", createConnector)
		p.Post("{connectorId}/", updateConnector)
		p.Post("{connectorId}/del", delConnector)
	}
}

func getConnectors(ctx iris.Context) () {

	res := NewBean()

	pageNo := ctx.URLParamIntDefault("page", 1)
	pageSize := ctx.URLParamIntDefault("size", 5)

	connectors, page, err := repo.ListConnector(pageNo, pageSize)
	if err != nil || len(connectors) == 0 {
		res.Code = -1
	} else {
		res.Item = &connectors
		res.Page = &page
	}

	_, _ = ctx.JSON(&res)
}

func getConnector(ctx iris.Context) () {

	connectorId := ctx.Params().Get("connectorId")
	connectorConfig, err := repo.GetConnectorConfig(connectorId)

	var res = NewBean()
	if err == nil {
		res.Item = &connectorConfig
	} else {
		res.Code = -1
	}
	_, _ = ctx.JSON(&res)
}

type ConnectorForm struct {
	Name  string
	Type  string
	Props string
}

func createConnector(ctx iris.Context) {
	res := NewBean()
	var err error
	form := ConnectorForm{}
	err = ctx.ReadForm(&form)
	var connConf repo.ConnectorConfig
	if err == nil {
		connConf, err = repo.CreateConnectorConf(form.Name, form.Type, form.Props)
	}

	if err == nil {
		res.Item = &connConf
	} else {
		res.Code = -1
	}
	_, _ = ctx.JSON(&res)
}

func updateConnector(ctx iris.Context) {

}

func delConnector(ctx iris.Context) {
	res := NewBean()
	connectorId := ctx.Params().Get("connectorId")
	err := repo.DeleteConnectorConfig(connectorId)
	if err != nil {
		res.Code = -1
	}
	_, _ = ctx.JSON(&res)
}
