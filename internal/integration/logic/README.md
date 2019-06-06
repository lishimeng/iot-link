## Logic Script
编写数据处理逻辑,入口参数为message:LinkMessage.

logic script返回message.data,可以在logic中编辑data部分数据

```text
function execute(message:LinkMessage);
```


message定义
```ecma script level 4

message = {
    "applicationID": "1",
    "applicationName": "app_test",
    "deviceID": "1001",
    "deviceName": "tbs001",
    "data": {
        "humidity": 45
    }
};
```
event函数定义
```ecma script level 4
function event(target, properties) {
    // dummy
    console.log(target, properties);
}
```
Logic script接口 (Demo)

function execute(message) return message.data
```ecma script level 4
// 
function execute(message){
    humidity = message.data["humidity"];
    if (humidity < 5) {
        
        var target = {
             "applicationID": "",
             "deviceID": ""
         };
        var properties = {
            "state": "1"
        };
        // 控制另外一个设备
        event(target, properties);
        if (message.data) {
            return message.data;
        } else {
            return {};
        }
    }
}
```
