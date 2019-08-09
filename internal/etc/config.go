package etc

import (
	"github.com/lishimeng/go-libs/etc"
)

var Config Configuration

var configName string

var envPath []string

func LoadEnvs() error {
	err := etc.LoadEnvs(configName, envPath, &Config)
	return err
}

func SetEnvPath(env []string) () {
	envPath = env
}

func SetConfigName(name string) () {
	configName = name
}
