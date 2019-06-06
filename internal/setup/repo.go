package setup

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	"github.com/lishimeng/iot-link/internal/codec"
	"github.com/lishimeng/iot-link/internal/connector"
	"github.com/lishimeng/iot-link/internal/db"
	"github.com/lishimeng/iot-link/internal/db/repo"
	"github.com/lishimeng/iot-link/internal/etc"
	"github.com/lishimeng/iot-link/internal/model"
	"github.com/lishimeng/persistence"
)

func DBRepo() (err error) {
	var config = persistence.PostgresConfig{
		UserName: etc.Config.Db.User,
		Password: etc.Config.Db.Password,
		Host: etc.Config.Db.Host,
		Port: etc.Config.Db.Port,
		DbName: etc.Config.Db.Database,
		MaxIdle: 5,
		MaxConn: 10,
		InitDb: true,
	}

	err = orm.RegisterDriver("postgres", orm.DRPostgres)
	//orm.RegisterModel(new(repo.DeviceConfig))
	orm.RegisterModel(new(repo.AppConfig))
	orm.RegisterModel(new(repo.DataPoint))
	orm.RegisterModel(new(repo.LogicScript))
	orm.RegisterModel(new(repo.ConnectorConfig))
	orm.RegisterModel(new(repo.CodecScript))
	err = db.Init(config)
	if err == nil {
		//initTestData()
	}

	return err
}

func initTestData() {

	// test code
	name := "LoraWan192.168.1.12"
	t := connector.LoraWanType
	props := map[string]string {
		"broker": "tcp://192.168.1.12:1883",
		"clientId": "iot_link_sample",
		"upLink": "application/+/device/+/rx",
		"downLink": "application/%s/device/%s/tx",
	}
	c, _ := repo.CreateConnectorConfig(name, t, props)

	app, _ := repo.CreateApp("4", "测试", codec.IntoyunTLV, c.Id)

	encodeScript := `
// object -> byte array
function encode(map) {
    return [0x11, 0x14, 0x15, 0x22];
}
`
	decodeScript := `
// byte array -> object
function decode(data) {
    return {"a": "this","b": "is", "c": "a", "d": "message"};
}
`
	_, _ = repo.CreateJs(app.AppId, encodeScript, decodeScript)

	logicContent := `
function execute(message){
    var target = {
		 "applicationID": "4",
		 "deviceID": "0fb778900"
	 };
	var properties = {
		"switch": "1"
	};
	// 控制另外一个设备
	event(target, properties);
}
`
	_, _ = repo.CreateLogic(app.AppId, logicContent)

	dps := make(map[string]model.DataPoint)
	dp1 := model.DataPoint{
		Index: 1,
		Name: "switch",
		Type: 1,
		Length: 1,
		UpLink: true,
		DownLink: true,
	}
	dps[fmt.Sprintf("%d", dp1.Index)] = dp1
	bs, _ := json.Marshal(&dps)
	_, _ = repo.CreateDataPoint(app.AppId, string(bs))
}
