package setup

import "github.com/lishimeng/iot-link/internal/integration/persistent"

func Influx() error {
	err := persistent.Init()
	return err
}
