package etc

import (
	"container/list"
	"fmt"
	"github.com/BurntSushi/toml"
	log "github.com/jeanphorn/log4go"
	"path/filepath"
)

var Config Configuration

var configName string

var envPath = new(list.List)

func LoadEnvs() error {
	for e := envPath.Front(); e != nil; e = e.Next() {
		tmp := e.Value
		var path = tmp.(string)
		file := fmt.Sprintf("%s/%s", path, configName)
		f, err := filepath.Abs(file)
		if err != nil {
			log.Info("Config file is not in %s", path)
		} else {
			if _, err := toml.DecodeFile(f, &Config); err != nil {
				log.Info("Config file is not in %s %s", path, err)
			} else {
				log.Info("%s. Version[%s]", Config.Name, Config.Version)
				break
			}
		}

	}

	return nil
}

func AddEnvPath(env string) () {
	envPath.PushBack(env)
}

func SetConfigName(name string) () {
	configName = name
}
