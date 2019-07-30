package intoyun

import (
	"encoding/json"
	"fmt"
	"github.com/lishimeng/go-libs/codec/tlv"
	"github.com/lishimeng/iot-link/internal/codec"
	"github.com/lishimeng/iot-link/internal/db/repo"
	"github.com/lishimeng/iot-link/internal/model"
	"github.com/lishimeng/iot-link/internal/tools"
)

type intoyunTlvCodec struct {

}

func New() codec.Coder {

	c := intoyunTlvCodec{}
	var plugin codec.Coder = &c
	return plugin
}

///
func (c intoyunTlvCodec) Decode(appId string, raw []byte) (props map[string]interface{}, err error) {
	var dp repo.DataPoint
	dp, err = repo.GetDataPoint(appId)
	if err != nil {
		return props, err
	}

	var structure = make(map[string]model.DataPoint)
	err = json.Unmarshal([]byte(dp.DataPoints), &structure)
	if err != nil {
		return props, err
	}
	payload := raw[1:]// 拆开包头
	frame := tlv.DecodeToFrame(payload)
	props = make(map[string]interface{})
	for _, tag := range frame.Tags {

		tagId := tag.TagId
		if value, ok := structure[fmt.Sprintf("%d", tagId)]; ok {
			props[value.Name] = tag.TagData
		}
	}
	return props, err
}

func (c intoyunTlvCodec) Encode(appId string, props map[string]interface{}) (raw []byte, err error) {

	var dp repo.DataPoint
	dp, err = repo.GetDataPoint(appId)
	if err != nil {
		return raw, err
	}

	var structure = make(map[string]model.DataPoint)
	err = json.Unmarshal([]byte(dp.DataPoints), &structure)
	if err != nil {
		return raw, err
	}

	frame := tlv.TagFrame{}
	seq := 0
	tags := make([]tlv.Tag, len(props))
	for key, value := range props {

		tagData := tools.Convert2Str(value)
		for _, tlvConfig := range structure {
			if tlvConfig.Name == key {
				tag := tlv.Tag{
					TagId: tlvConfig.Index,
					TagType: tlvConfig.Type,
					TagLength: tlvConfig.Length,
					TagData: tagData,
				}
				tags[seq] = tag
				seq++
				break
			}
		}
	}
	if seq > 0 {
		frame.Tags = tags[0:seq]
		result := tlv.Encode(&frame)
		result = "1" + result// 补齐包头
		raw = []byte(result)
	} else {
		err = fmt.Errorf("data points are empty")
	}

	return raw, err
}
