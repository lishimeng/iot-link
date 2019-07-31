package etc

type Configuration struct {
	Name    string
	Version string
	Db      db
	Web     web
	Influx  influx
}

type db struct {
	Host     string
	Database string
	Port     int
	User     string
	Password string
}

type web struct {
	Listen string
}

type influx struct {
	Host     string
	Database string
}
