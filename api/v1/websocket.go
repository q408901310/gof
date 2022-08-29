package v1

import "github.com/gogf/gf/v2/frame/g"

type WebsocketReq struct {
	g.Meta `path:"/ws" method:"get" tags:"Websocket" summary:"Send message"`
	//g.Meta `path:"/ws" method:"get" tags:"Websocket" summary:"You first hello api"`
}

type WebsocketRes struct {
	g.Meta `mime:"json" type:"string" example:"<html/>" dc:"It start websocket if success"`
}
