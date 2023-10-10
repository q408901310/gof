package action

import (
	"github.com/gogf/gf/v2/frame/g"
	"gof/internal/pb"
	"reflect"
)

var (
	actionMap = map[pb.ACTION]reflect.Value{
		pb.ACTION_LOGIN_ZONE: reflect.ValueOf(Login.start),
	}
)

func Run(ctx g.Ctx, key pb.ACTION, data *pb.DataMsg) {
	action, ok := actionMap[key]
	if !ok {
		return
	}
	in := []reflect.Value{
		reflect.ValueOf(ctx),
		reflect.ValueOf(data),
	}
	action.Call(in)
}
