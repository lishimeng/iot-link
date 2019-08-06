package repo

import (
	"fmt"
	"github.com/lishimeng/iot-link/internal/db"
	"time"
)

func SaveDownLinkMessage(appId string, deviceId string, payload string) (msg DownLinkData, err error) {

	now := time.Now().Unix()
	msg = DownLinkData{
		Id: genDownLinkMsg(appId, deviceId),
		AppId: appId,
		DeviceId: deviceId,
		Payload: payload,
		Status: Enabled,
		CreateTime: now,
		UpdateTime: now,
	}

	_, err = db.Orm.Context.InsertOrUpdate(&msg, "id")
	err = checkErr(err)
	return msg, err
}

func SaveDelayedDownLinkMessage(appId string, deviceId string, messageType int, payload string) (msg DelayedDownLinkData, err error) {

	now := time.Now().Unix()
	msg = DelayedDownLinkData{
		Id: genDelayedDownLinkMsg(appId, deviceId, messageType),
		AppId: appId,
		DeviceId: deviceId,
		Payload: payload,
		Type: messageType,
		Status: Enabled,
		CreateTime: now,
		UpdateTime: now,
	}

	_, err = db.Orm.Context.InsertOrUpdate(&msg, "id")

	err = checkErr(err)
	if err == nil {
		// TODO save log
	}
	return msg, err
}

func DeleteDownLinkMessage(id string) (err error) {

	_, err = db.Orm.Context.Delete(&DownLinkData{Id: id})
	err = checkErr(err)
	return err
}

func DeleteDelayedDownLinkMessage(id string) (err error) {

	_, err = db.Orm.Context.Delete(&DelayedDownLinkData{Id: id})
	err = checkErr(err)
	return err
}

func GetDelayedMessage(appId string, deviceId string, messageType int) (msg DelayedDownLinkData, err error) {
	id := genDelayedDownLinkMsg(appId, deviceId, messageType)
	msg = DelayedDownLinkData{Id: id}
	err = db.Orm.Context.Read(&msg)
	return msg, err
}

func CountMessage() (size int64, err error) {

	q := DownLinkData{}
	size, err = db.Orm.Context.QueryTable(&q).Count()
	return size, err
}

func ListMessage(limit int) (msg []DownLinkData, err error) {

	q := DownLinkData{}
	res := new([]DownLinkData)
	_, err = db.Orm.Context.QueryTable(&q).Limit(limit, 0).OrderBy("CreateTime").All(res)
	if err == nil {
		msg = *res
	}
	return msg, err
}

func genDownLinkMsg(appId string, deviceId string) string {
	return fmt.Sprintf("%s_%s", appId, deviceId)
}

func genDelayedDownLinkMsg(appId string, deviceId string, messageType int) string {
	return fmt.Sprintf("%s_%s_%d", appId, deviceId, messageType)
}

func CloseDelayedMessage() {

}