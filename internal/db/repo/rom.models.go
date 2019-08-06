package repo

const (
	Push = iota
	Pull
)
const (
	Disabled = iota
	Enabled
)

// 暂时不需要
type DeviceConfig struct {
	DeviceId string `orm:"pk"`
	AppId    string
	// lora 设备接收通道
	FPort      int32 `orm:"column(f_port)"`
	CreateTime int64
}

type AppConfig struct {
	AppId string `orm:"pk"`

	AppDescription string

	// 编码插件
	CodecType string

	// Proxy ID
	Connector string

	CreateTime int64

	UpdateTime int64
}

type ConnectorConfig struct {
	Id         string `orm:"pk"`
	Name       string `orm:"unique"`
	Type       string `orm:"column(type)"`
	Props      string `orm:"column(props)"`
	CreateTime int64
	UpdateTime int64
}

type LogicScript struct {
	AppId      string `orm:"pk"`
	Content    string `orm:"column(content)"`
	CreateTime int64
	UpdateTime int64
}

type CodecScript struct {
	AppId         string `orm:"pk"`
	EncodeContent string
	DecodeContent string
	CreateTime    int64
	UpdateTime    int64
}

type DataPoint struct {
	AppId      string `orm:"pk"`
	DataPoints string
	CreateTime int64
	UpdateTime int64
}

type DownLinkData struct {
	Id string `orm:"pk"`
	DeviceId string
	AppId string
	Payload string
	Status int
	CreateTime int64
	UpdateTime int64
}

type DelayedDownLinkData struct {
	Id string `orm:"pk"`
	DeviceId string
	AppId string
	Payload string
	Type int // 1:Push平台主动发送,0:Pull设备主动拉取
	Status int
	CreateTime int64
	UpdateTime int64
}

type DownLinkLog struct {
	Id string `orm:"pk"`
	DeviceId string
	AppId string
	Payload string
	Status int
	CreateTime int64
	UpdateTime int64
}

type TriggerConfig struct {
	Id string `orm:"pk"`
	AppId string
	Content string
	CreateTime int64
	UpdateTime int64
	Status int
}