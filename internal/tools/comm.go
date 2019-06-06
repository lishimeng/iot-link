package tools

import "fmt"

func Convert2Str(inter interface{}) string {

	var str string
	switch inter.(type) {
	case string:
		str = inter.(string)
		break
	case int:
		str = fmt.Sprintf("%d", inter.(int))
		break
	case int8:
		str = fmt.Sprintf("%d", inter.(int8))
		break
	case int16:
		str = fmt.Sprintf("%d", inter.(int16))
		break
	case int64:
		str = fmt.Sprintf("%d", inter.(int64))
		break
	case int32:
		str = fmt.Sprintf("%d", inter.(int32))
		break
	case uint:
		str = fmt.Sprintf("%d", inter.(uint))
		break
	case uint8:
		str = fmt.Sprintf("%d", inter.(uint8))
		break
	case uint16:
		str = fmt.Sprintf("%d", inter.(uint16))
		break
	case uint32:
		str = fmt.Sprintf("%d", inter.(uint32))
		break
	case uint64:
		str = fmt.Sprintf("%d", inter.(uint64))
		break
	case float64:
		str = fmt.Sprintf("%f", inter.(float64))
		break
	case float32:
		str = fmt.Sprintf("%f", inter.(float32))
		break
	case []byte:
		str = string(inter.([]byte))
		break
	}
	return str
}
