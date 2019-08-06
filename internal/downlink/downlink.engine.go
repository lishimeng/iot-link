package downlink

import (
	"github.com/lishimeng/iot-link/internal/model"
	"time"
)

type MessageDownLinkEngine interface {
	StartDownLink()
	SaveMessage(data model.EventPayload) error // 保存消息
	ActiveDelayedMessage(appId string, deviceId string) // 激活消息,准备发送
	PullDelayedMessage(appId string, deviceId string) (eventPayload model.EventPayload, err error) // 查询设备消息
	CloseExpiredDelayedMessages() // 关闭过期的消息
}

var _singleton MessageDownLinkEngine

type msgEngine struct {
	// 每次读取数据量
	fetchSize      int
	idleTime       time.Duration
	downLinkStatus bool // true: running, false: stopped
}

func GetInstance() MessageDownLinkEngine {
	return _singleton
}

func Init(fetchSize int, idleTime time.Duration) {

	m := msgEngine{
		fetchSize:      fetchSize,
		idleTime:       idleTime,
		downLinkStatus: false,
	}
	_singleton = &m
}
