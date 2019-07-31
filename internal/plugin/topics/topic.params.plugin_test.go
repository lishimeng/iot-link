package topics

import (
	"fmt"
	"testing"
)

func TestTopicTemplate(t *testing.T) {
	tpl := "application/{AppID}/{deviceID}"
	test := "application/abs/uuf"
	params, errr := DeviceUpLinkParamTpl(tpl, test)
	if errr != nil {
		fmt.Println(errr)
	}
	appId, ok := params["AppID"]
	if !ok {
		t.Errorf("expect param %s", "AppID")
		return
	}
	if appId != "abs" {
		t.Errorf("expect param %s value is %s, but %s", "AppID", "abs", appId)
		return
	}
	deviceID, ok := params["deviceID"]
	if !ok {
		t.Errorf("expect param %s", "deviceID")
		return
	}
	if deviceID != "uuf" {
		t.Errorf("expect param %s value is %s, but %s", "deviceID", "uuf", deviceID)
		return
	}
}

func TestTopicTemplateFormatErr(t *testing.T) {
	tpl := "application/{AppID}/{deviceEUI}/"
	test := "application/abs/uuf"
	_, err := DeviceUpLinkParamTpl(tpl, test)
	if err == nil {
		t.Errorf("expect an error")
	}
}

func TestTopicTemplateParamSize(t *testing.T) {
	tpl := "application/{AppID}/{deviceEUI"
	test := "application/abs/uuf"
	params, _ := DeviceUpLinkParamTpl(tpl, test)
	if len(params) != 1 {
		t.Errorf("expect %d params but see %d", 1, len(params))
	}
}
