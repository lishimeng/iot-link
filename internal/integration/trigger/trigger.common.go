package trigger

import (
	"fmt"
	"github.com/lishimeng/go-libs/script"
	"github.com/lishimeng/iot-link/internal/model"
	"github.com/robertkrimen/otto"
	"time"
)

func buildTriggerContent(t model.Trigger, msg model.LinkMessage) string {
	res := make([]string, len(t.Tags))
	for index, tag := range t.Tags {
		if v, ok := msg.Data[tag.Key]; ok {
			jsContent := buildJsContent(v, tag)
			res[index] = jsContent
		}
	}

	javascript := ""
	for _, js := range res {
		javascript += js
	}
	return javascript
}

func calcTrigger(javascript string) (b bool, err error) {
	vm := otto.New()
	var result otto.Value
	result, err = script.Execute(vm, javascript, 10 * time.Millisecond)
	if err == nil {
		if result.IsBoolean() {
			b, err = result.ToBoolean()
		}
	}
	return b, err
}

func buildJsContent(v interface{}, tag model.TriggerTag) string {
	value := strValue(v)
	compare := strValue(tag.Value)
	operator := strValue(tag.Operator)
	condition := tag.Condition
	c := getCondition(condition)
	return fmt.Sprintf(" %s (%s %s %s)", c, value, operator, compare)
}

func getCondition(c string) (s string) {
	s = ""
	if len(c) > 0 {
		if c == "AND" {
			s = "&&"
		} else if c == "OR" {
			s = "||"
		}
	}
	return s
}

func strValue(v interface{}) (str string) {

	str = ""
	switch v.(type) {
	case string:
		str = v.(string)
	case int:
		s := v.(int)
		str = fmt.Sprintf("%d", s)
	case uint:
		s := v.(int)
		str = fmt.Sprintf("%d", s)
	case int8:
		s := v.(int8)
		str = fmt.Sprintf("%d", s)
	case uint8:
		s := v.(uint8)
		str = fmt.Sprintf("%d", s)
	case int16:
		s := v.(int16)
		str = fmt.Sprintf("%d", s)
	case uint16:
		s := v.(uint16)
		str = fmt.Sprintf("%d", s)
	case int32:
		s := v.(int32)
		str = fmt.Sprintf("%d", s)
	case uint32:
		s := v.(uint32)
		str = fmt.Sprintf("%d", s)
	case int64:
		s := v.(int64)
		str = fmt.Sprintf("%d", s)
	case uint64:
		s := v.(uint64)
		str = fmt.Sprintf("%d", s)
	case float32:
		s := v.(float32)
		str = fmt.Sprintf("%f", s)
	case float64:
		s := v.(float64)
		str = fmt.Sprintf("%f", s)
	}
	return str
}