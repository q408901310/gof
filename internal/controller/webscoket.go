package controller

import (
	"context"
	v1 "gof/api/v1"
	"gof/internal/service"
)

type cWebsocket struct{}

var (
	Websocket = cWebsocket{}
)

func (c *cWebsocket) Websocket(ctx context.Context, req *v1.WebsocketReq) (res *v1.WebsocketRes, err error) {
	err = service.Websocket().Start(ctx)
	return
}
