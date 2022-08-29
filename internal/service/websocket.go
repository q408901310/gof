package service

import (
	"context"
	"gof/internal/pb"

	"github.com/gogo/protobuf/proto"

	"github.com/gogf/gf/v2/net/ghttp"

	"github.com/gogf/gf/v2/container/gmap"

	"github.com/gogf/gf/v2/frame/g"
)

type (
	sWebsocket struct {
		clientMap *gmap.Map
	}

	client struct {
		Ws       *ghttp.WebSocket
		Uid      int
		PlayerId int
	}
)

var (
	insWebsocket = sWebsocket{
		clientMap: gmap.New(true),
	}
)

func Websocket() *sWebsocket {
	return &insWebsocket
}

func (s *sWebsocket) Start(ctx context.Context) (err error) {
	var (
		r       = g.RequestFromCtx(ctx)
		ws      *ghttp.WebSocket
		msgByte []byte
		msg     *pb.Msg
	)

	if ws, err = r.WebSocket(); err != nil {
		g.Log().Error(ctx, err)
		return
	}
	c := &client{
		Ws:       ws,
		Uid:      0,
		PlayerId: 0,
	}
	// Create session data for current websocket.
	s.clientMap.Set(ws, c)
	for {
		// Blocking reading message from current websocket.
		_, msgByte, err = ws.ReadMessage()
		if err != nil {
			// Remove session info.
			s.clientMap.Remove(ws)
			return nil
		}
		err := proto.Unmarshal(msgByte, msg)
		if err != nil {
			return err
		}
		g.Dump(msg)
	}
}

func (s *sWebsocket) Close(ws *ghttp.WebSocket) {
	s.clientMap.Remove(ws)
	err := ws.Close()
	if err != nil {
		return
	}
}

func (s *sWebsocket) CloseAll() {
	s.clientMap.Iterator(func(ws interface{}, c interface{}) bool {
		err := ws.(*ghttp.WebSocket).Close()
		return err == nil
	})
	// TODO 中间可能有并发往 clientMap 添加新成员，后续完善
	s.clientMap.Clear()
}
