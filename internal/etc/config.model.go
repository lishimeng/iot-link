package etc

type Configuration struct {
	Name    string
	Version string
	Db      db
	Web     web
	Influx  influx
	DownLink downLink `toml:"down-link"`
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
	Enable   int
}

type downLink struct {
	IdleTime int64 `toml:"idle-time"`
	FetchSize int `toml:"fetch-size"`
	LogEnable bool `toml:"log-enable"`
}
