package repo

import (
	"github.com/lishimeng/iot-link/internal/db"
	"time"
)

func GetDataPoint(appId string) (dp DataPoint, err error) {

	dp = DataPoint{AppId: appId}
	err = db.Orm.Context.Read(&dp)
	return dp, err
}

func CreateDataPoint(appId string, content string) (dp DataPoint, err error) {

	dp = DataPoint{
		AppId:      appId,
		DataPoints: content,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}
	_, err = db.Orm.Context.Insert(&dp)
	err = checkErr(err)
	return dp, err
}

func UpdateDataPoint(appId string, content string) (dp DataPoint, err error) {

	dp = DataPoint{
		AppId:      appId,
		DataPoints: content,
		UpdateTime: time.Now().Unix(),
	}
	_, err = db.Orm.Context.Update(&dp, "DataPoints", "UpdateTime")
	err = checkErr(err)
	return dp, err
}

func DeleteDataPoint(appId string) (dp DataPoint, err error) {

	dp = DataPoint{AppId: appId}
	_, err = db.Orm.Context.Delete(&dp)
	err = checkErr(err)
	return dp, err
}
