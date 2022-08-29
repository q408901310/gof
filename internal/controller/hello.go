package controller

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	v1 "gof/api/v1"
)

var (
	Hello = cHello{}
)

type cHello struct{}

func (c *cHello) Hello(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error) {
	g.RequestFromCtx(ctx).Response.Writeln("Hello 世界!")
	return
}

//func (c *cHello) CardSelect(ctx context.Context, req *v1.CardSelectReq) (res *v1.CardSelectRes, err error) {
//	g.RequestFromCtx(ctx).Response.Writeln("Hello 世界!")
//	return
//}