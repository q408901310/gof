package controller

import (
	"context"
	v1 "gof/api/v1"
	"gof/internal/service"
)

var (
	Server = cServer{}
)

type cServer struct{}

func (c *cServer) List(ctx context.Context, req *v1.ServerListReq) (res *v1.ServerListRes, err error) {
	res = &v1.ServerListRes{}
	res.List, err = service.Server().List(ctx)
	return
}
