package action

import (
	"github.com/gogf/gf/v2/frame/g"
	"reflect"
)

var (
	Action = g.Map{
		"login": Login,
	}
)
var (
	actions = map[string]reflect.Value{
		"l_l": reflect.ValueOf(Login.login),
	}
)

func GetAction(key string) (res reflect.Value, err error) {
	ok := false
	if res, ok = actions[key]; !ok {
		return
	}
	return
}
