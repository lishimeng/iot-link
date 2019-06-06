package api

import (
	"github.com/kataras/iris"
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


}

func getConnector(ctx iris.Context) () {


}

func createConnector(ctx iris.Context) {

}

func updateConnector(ctx iris.Context) {

}

func delConnector(ctx iris.Context) {

}