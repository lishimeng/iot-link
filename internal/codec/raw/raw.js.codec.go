package raw

import (
	"fmt"
	"github.com/lishimeng/go-libs/script"
	"github.com/lishimeng/iot-link/internal/codec"
	"github.com/lishimeng/iot-link/internal/db/repo"
)

type jsRawCodec struct {
}

func New() codec.Coder {

	c := jsRawCodec{}
	var plugin codec.Coder = &c
	return plugin
}

// javascript 解析raw格式
func (c jsRawCodec) Decode(appId string, data []byte) (props map[string]interface{}, err error) {

	// TODO
	// find js javascript
	jsConfig, err := repo.GetJs(appId)
	if err != nil {
		return props, err
	}
	js := jsConfig.DecodeContent
	engine, err := script.Create(js)
	if err != nil {
		return props, err
	}
	value, err := engine.Invoke("decode", data)
	if err != nil {
		return props, err
	}

	raw, err := value.Export()
	if err != nil {
		return props, err
	}
	switch raw.(type) {
	case map[string]interface{}:
		props = raw.(map[string]interface{})
		break
	default:
		err = fmt.Errorf("decode result must be type of map[string]interface{}")
	}

	// decode
	return props, err
}

func (c jsRawCodec) Encode(appId string, props map[string]interface{}) (data []byte, err error) {

	jsConfig, err := repo.GetJs(appId)
	if err != nil {
		return data, err
	}

	js := jsConfig.EncodeContent
	engine, err := script.Create(js)
	if err != nil {
		return data, err
	}
	value, err := engine.Invoke("encode", props)
	if err != nil {
		return data, err
	}

	raw, err := value.Export()
	switch raw.(type) {
	case []byte:
		data = raw.([]byte)
		break
	case []int64:
		tmp := raw.([]int64)
		if len(tmp) > 0 {
			data = make([]byte, len(tmp))
			for index, item := range tmp {
				data[index] = byte(item)
			}
		}
		break
	default:
		err = fmt.Errorf("encode result must be type of byte array")
	}
	return data, err
}
