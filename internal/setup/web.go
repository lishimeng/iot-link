package setup

import (
	"github.com/kataras/iris"
	"github.com/lishimeng/iot-link/internal/api"
	"github.com/lishimeng/iot-link/internal/web"
)

func Web() error {

	var components = []func(application *iris.Application){
		api.SetupDataPoint,
		api.SetupApplication,
		api.SetupLogic,
		api.SetupConnector,
		api.SetupCodecJs,
		api.SetupTrigger,
	}
	go web.Run(components...)
	return nil
}
