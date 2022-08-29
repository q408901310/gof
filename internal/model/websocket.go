package model

import "github.com/gogf/gf/v2/net/ghttp"

type WebsocketInitInput struct {
	R  *ghttp.Request
	Ws *ghttp.WebSocket
}
