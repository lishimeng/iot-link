package lorawan

import (
	"encoding/json"
)

func onDataUpLink(raw []byte) (payload PayloadRx, err error) {

	err = json.Unmarshal(raw, &payload)
	if err != nil {
		return
	}
	return payload, err
}

func convertJsonDownLinkData(payload PayloadTx) (jsonStr string) {

	bs, err := json.Marshal(payload)
	if err != nil {
		jsonStr = ""
	} else {
		jsonStr = string(bs)
	}
	return jsonStr
}
