package module

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/golang/protobuf/proto"
	"gof/internal/pb"
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
	build := &pb.Msg{
		Sequence: 0,
		Action:   action,
		Code:     code,
	}
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
