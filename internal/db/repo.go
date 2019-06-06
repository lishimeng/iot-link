package db

import (
	"github.com/lishimeng/persistence"
)
var Orm persistence.OrmContext

func Init(config persistence.PostgresConfig) (err error) {

	Orm, err = persistence.InitPostgresOrm(config)
	return err
}