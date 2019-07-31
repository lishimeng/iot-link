package model

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestLinkMessageData(t *testing.T) {

	var d = LinkMessage{
		Data: map[string]interface{}{
			"gps": map[string]string{
				"latitude": "1234.33221234",
				"lat":      "N",
				"logitude": "313455.1133311",
				"log":      "E",
			},
			"rfid": map[string]string{
				"tid": "1111111111111",
				"epc": "2222222222222",
			},
			"bat": map[string]string{
				"level": "10",
			},
		},
	}

	b, err := json.Marshal(&d.Data)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	fmt.Println(string(b))
}

func TestA(t *testing.T) {

}
