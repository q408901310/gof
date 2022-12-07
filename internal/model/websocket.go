package model

import "github.com/gogf/gf/v2/net/ghttp"

type WebsocketInitIn struct {
	R  *ghttp.Request
	Ws *ghttp.WebSocket
}
