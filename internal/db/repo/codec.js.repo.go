package repo

import (
	"github.com/lishimeng/iot-link/internal/db"
	"time"
)

func CreateJs(appId string, encodeScript string, decodeScript string) (codec CodecScript, err error) {

	c := CodecScript{
		AppId: appId,
		EncodeContent: encodeScript,
		DecodeContent: decodeScript,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}
	_, err = db.Orm.Context.Insert(&c)
	return c, err
}

func CreateOrUpdateJs(appId string, encodeScript string, decodeScript string) (codec CodecScript, err error) {
	codec, err = GetJs(appId)

	if err == nil {
		// update
		codec.DecodeContent = encodeScript
		codec.DecodeContent = decodeScript
		codec.UpdateTime = time.Now().Unix()
		_, err = db.Orm.Context.Update(&codec, "DecodeContent", "DecodeContent", "UpdateTime")
	} else {
		// create
		codec, err = CreateJs(appId, encodeScript, decodeScript)
	}

	return codec, err
}


func GetJs(appId string) (js CodecScript, err error) {

	js = CodecScript{AppId: appId}
	err = db.Orm.Context.Read(&js)
	return js, err
}

func DeleteJs(appId string) (err error) {

	q := CodecScript{AppId: appId}
	_, err = db.Orm.Context.Delete(&q)
	return err
}