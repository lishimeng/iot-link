package repo

// 暂时不需要
type DeviceConfig struct {

	DeviceId string	`orm:"pk"`
	AppId string
	// lora 设备接收通道
	FPort int32 `orm:"column(f_port)"`
	CreateTime int64
}

type AppConfig struct {
	AppId string `orm:"pk"`

	AppDescription string

	// 编码插件
	CodecType string

	// Connector ID
	Connector string

	CreateTime int64

	UpdateTime int64
}

type ConnectorConfig struct {
	Id string `orm:"pk"`
	Name string `orm:"unique"`
	Type string `orm:"column(type)"`
	Props string `orm:"column(props)"`
	CreateTime int64
	UpdateTime int64
}

type LogicScript struct {
	AppId string `orm:"pk"`
	Content string `orm:"column(content)"`
	CreateTime int64
	UpdateTime int64
}

type CodecScript struct {
	AppId string `orm:"pk"`
	EncodeContent string
	DecodeContent string
	CreateTime int64
	UpdateTime int64
}

type DataPoint struct {
	AppId string `orm:"pk"`
	DataPoints string
	CreateTime int64
	UpdateTime int64
}
