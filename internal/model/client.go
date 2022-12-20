package model

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/golang/protobuf/proto"
	"gof/internal/pb"
	"gof/internal/service"
)

type Client struct {
	ws        *ghttp.WebSocket
	Uid       uint
	PlayerId  int
	MsgType   int
	CtxCancel context.CancelFunc
}

func NewClient(ws *ghttp.WebSocket, uid uint) *Client {
	return &Client{
		ws:       ws,
		Uid:      uid,
		PlayerId: 0,
		MsgType:  0,
	}
}

func (c *Client) SendErrorMessage(action pb.ACTION, code pb.CODE) {
	build := service.Message().BuildPBMsg(action, code)
	msg, err := proto.Marshal(build)
	if err != nil {
		return
	}
	err = c.ws.WriteMessage(c.MsgType, msg)
	if err != nil {
		return
	}
}

func (c *Client) Close() error {
	return c.ws.Close()
}
