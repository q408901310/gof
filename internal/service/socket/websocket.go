package socket

import (
	"context"
	"github.com/gogo/protobuf/proto"
	"gof/internal/action"
	"gof/internal/pb"

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

func (s *sWebsocket) Connect(ctx context.Context) (err error) {
	var (
		r       = g.RequestFromCtx(ctx)
		ws      *ghttp.WebSocket
		msgByte []byte
		msg     pb.Msg
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
			continue
		}
		err = proto.Unmarshal(msgByte, &msg)
		if err != nil {
			s.sendErrorMessage(c, msg.Action, pb.Code_SYS_PARAMETER)
			continue
		}
		//msg.SessionId
		act, ok := action.Action[msg.Action]
		if !ok {
			s.sendErrorMessage(c, msg.Action, pb.Code_SYS_ACTION)
		}
		//停服维护中 todo
		g.Dump(act)
		//act
	}
}

func (s *sWebsocket) sendErrorMessage(c *client, action string, code pb.Code) {
	//msg = Message().BuildMsg(action, code)
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
