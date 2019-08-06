package downlink

import (
	"encoding/json"
	log "github.com/jeanphorn/log4go"
	"github.com/lishimeng/iot-link/internal/db/repo"
	"github.com/lishimeng/iot-link/internal/event"
	"github.com/lishimeng/iot-link/internal/model"
	"time"
)

func (m *msgEngine) StartDownLink() {

	if m.downLinkStatus {
		log.Debug("downLink is running, skip start")
		return
	} else {
		log.Debug("start downLink...")
		m.downLinkStatus = true
	}
	for {
		size := m.sendAll()

		if size == 0 {
			idleTime := m.idleTime
			if idleTime <= 0 {
				idleTime = time.Second * 10
			}
			time.Sleep(idleTime)
		}
		//break// TODO Test
	}
}

func (m msgEngine) getMsgCount() int64 {
	size, err := repo.CountMessage()
	if err != nil {
		log.Debug(err)
		size = 0
	}

	if LogEnable {
		log.Debug("message size:%d", size)
	}
	return size
}

func (m msgEngine) sendAll() (size int) {

	size = 0
	if m.getMsgCount() > 0 {

		datas, err := m.fetchData()
		if err != nil {
			log.Debug(err)
			return size
		}
		if len(datas) > 0 {
			for _, item := range datas {
				err = m.downLink(item)
				if err == nil {
					size += 1
				} else {
					log.Debug(err)
				}
			}
		}
	}
	return size
}

// 发送列表(自动发送列表)
func (m msgEngine) fetchData() (datas []repo.DownLinkData, err error) {
	if LogEnable {
		log.Debug("fetch data...")
	}
	datas, err = repo.ListMessage(m.fetchSize)

	if LogEnable {
		log.Debug("data size: %d", len(datas))
	}
	return datas, err
}

func (m msgEngine) downLink(data repo.DownLinkData) (err error) {

	// TODO 记录发送log

	appConfig, err := repo.GetApp(data.AppId)
	if err == nil {
		var payload = make(map[string]interface{})
		err = json.Unmarshal([]byte(data.Payload), &payload)
		if err == nil {
			connectorId := appConfig.Connector
			target := model.Target{
				ConnectorId: connectorId,
				AppId:       data.AppId,
				DeviceId:    data.DeviceId,
			}
			event.GetInstance().Send(target, payload)
			m.onDownLinkComplete(data.Id)
		}
	} else {
		log.Debug("skip downLink, no app config")
	}
	return err
}
func (m msgEngine) onDownLinkComplete(id string) {

	log.Debug("remove completed message:%s", id)
	err := repo.DeleteDownLinkMessage(id)
	if err != nil {
		log.Debug(err)
	}
}
