package topics

import (
	"fmt"
	"strings"
)

func DeviceUpLinkParamTpl(tpl string, topic string) (res map[string]string, err error) {

	ss := strings.Split(tpl, "/")
	st := strings.Split(topic, "/")
	if len(ss) == 0 || len(st) != len(ss) {
		err = fmt.Errorf("topics is not match the template %s[%s]", tpl, topic)
	} else {
		res = make(map[string]string)
		for i, v := range ss {
			if strings.HasPrefix(v, "{") && strings.HasSuffix(v, "}") {
				name := v[1 : len(v)-1]
				value := st[i]
				res[name] = value
			}
		}
	}
	return res, err
}
