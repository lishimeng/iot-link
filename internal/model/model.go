package model

/// 事件接受者
type Target struct {
	// 目标设备
	DeviceId string
	// 目标app
	AppId string
	// 网络
	ConnectorId string
}

type LinkMessage struct {
	// APP ID
	ApplicationID string
	// APP名(label)
	ApplicationName string
	// 设备ID
	DeviceID string
	// 设备名(label)
	DeviceName string
	// logic数据
	Data map[string]interface{}
	// 原始类型数据
	Raw []byte
}

// 数据点结构
type DataPoint struct {
	// 数据点序号
	Index int32   `json:"index"`
	// 名称
	Name string   `json:"name"`
	// 数据类型
	/// bool   0
	/// number 1
	/// enum   2
	/// txt    3
	Type int32    `json:"type"`
	// 数据长度
	Length int32  `json:"length,omitempty"`
	// 出现在上报中
	UpLink bool   `json:"upLinkEnable,omitempty"`
	// 出现在下发中
	DownLink bool `json:"downLinkEnable,omitempty"`
}