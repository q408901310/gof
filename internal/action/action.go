package action

import (
	"github.com/gogf/gf/v2/frame/g"
	"gof/internal/pb"
	"reflect"
)

var (
	actionMap = map[pb.ACTION]reflect.Value{
		pb.ACTION_LOGIN_ZONE: reflect.ValueOf(Login.zone),
	}
)

func Run(ctx g.Ctx, key pb.ACTION, params ...interface{}) error {
	action, ok := actionMap[key]
	if !ok {
		return nil
	}
	in := []reflect.Value{
		reflect.ValueOf(ctx),
	}
	if len(params) > 0 {
		for _, param := range params {
			in = append(in, reflect.ValueOf(param))
		}
	}
	action.Call(in)
	return nil
}
