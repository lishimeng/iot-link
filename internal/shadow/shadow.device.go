package shadow

// 设备影子
type Shadow struct {
	// profile
	Profiles map[string]Profile

	// 是否在线
	IsOnline bool

	// 是否需要自动同步profile
	// false: online时将profile重置
	// true: online时下发profile以达到恢复原值的目的
	IsAutoSync bool
}

type Profile struct {
	// 名称(ASCII码)
	Name string
	// 展示名(UTF-8可中文)
	Label string
	// 字段数据类型(枚举)
	Type string

	// 上报值
	Reported interface{}
	// 修改值
	Desired interface{}

	// upLink时间
	UpLinkTime int64
	// downLink时间
	DownLinkTime int64
}

type DataPointTypeEnum struct {
	string string
}
