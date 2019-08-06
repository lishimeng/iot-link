package message

import (
	"errors"
	"github.com/lishimeng/iot-link/internal/codec"
	"github.com/lishimeng/iot-link/internal/codec/intoyun"
	"github.com/lishimeng/iot-link/internal/codec/raw"
)

func decode(appId string, decodeType string, rawData []byte) (data map[string]interface{}, err error) {
	switch decodeType {
	case codec.Javascript:
		// find from raw js repo
		data, err = raw.New().Decode(appId, rawData)
		break
	case codec.IntoyunTLV:
		// find from tlv repo
		data, err = intoyun.New().Decode(appId, rawData)
		break
	default:
		// no codec plugin
		err = errors.New("unknown codec type")
		break
	}
	return data, err
}

func encode(appId string, encodeType string, msg map[string]interface{}) (data []byte, err error) {
	switch encodeType {
	case codec.Javascript:
		// find from raw js repo
		data, err = raw.New().Encode(appId, msg)
		break
	case codec.IntoyunTLV:
		// find from tlv repo
		data, err = intoyun.New().Encode(appId, msg)
		break
	default:
		// no codec plugin
		err = errors.New("unknown codec type")
		break
	}
	return data, err
}
