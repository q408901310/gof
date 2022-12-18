package service

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/grpool"
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
		pool      *grpool.Pool
		Close     bool
	}

	client struct {
		ws       *ghttp.WebSocket
		uid      uint
		PlayerId int
		msgType  int
	}
)

var (
	insWebsocket = sWebsocket{
		clientMap: gmap.New(true),
		pool:      grpool.New(1024),
		Close:     false,
	}
)

func Websocket() *sWebsocket {
	return &insWebsocket
}

func (s *sWebsocket) Connect(ctx context.Context) (err error) {
	var (
		r       = g.RequestFromCtx(ctx)
		ws      *ghttp.WebSocket
		msgType int
		msgByte []byte
		msg     pb.Msg
	)

	if ws, err = r.WebSocket(); err != nil {
		g.Log().Error(ctx, err)
		return
	}
	sessionId := r.GetHeader("sessionId")
	user := Session().GetUserBySessionId(ctx, sessionId)
	if user == nil {
		return gerror.New("请先登录")
	}
	c := s.AddClient(ws, user.Id)
	if c == nil {
		return gerror.New("维护中")
	}
	for {
		// Blocking reading message from current websocket.
		msgType, msgByte, err = ws.ReadMessage()
		if err != nil {
			// Remove session info.
			s.clientMap.Remove(ws)
			//continue
			return gerror.New("数据错误")
		}
		c.msgType = msgType
		err = proto.Unmarshal(msgByte, &msg)
		if err != nil {
			s.sendErrorMessage(c, msg.Action, pb.CODE_SYS_UNMARSHAL)
			continue
			//return
		}
		action.Run(ctx, msg.Action)
		//act, ok := action.Action[msg.Action]
		//if !ok {
		//	s.sendErrorMessage(c, msg.Action, pb.Code_SYS_ACTION)
		//	continue
		//}
		//停服维护中 todo
		//g.Dump(act)
		//act
	}
}

func (s *sWebsocket) sendErrorMessage(c *client, action pb.ACTION, code pb.CODE) {
	build := Message().buildPBMsg(action, code)
	msg, err := proto.Marshal(build)
	if err != nil {
		return
	}
	err = c.ws.WriteMessage(c.msgType, msg)
	if err != nil {
		return
	}
}

func (s *sWebsocket) AddClient(ws *ghttp.WebSocket, uid uint) (c *client) {
	if s.Close {
		return nil
	}
	c = &client{
		ws:  ws,
		uid: uid,
	}
	// Create session data for current websocket.
	s.clientMap.Set(c.uid, c)
	return c
}

func (s *sWebsocket) CloseClient(c *client) {
	s.clientMap.Remove(c.uid)
	err := c.ws.Close()
	if err != nil {
		return
	}
}

func (s *sWebsocket) CloseAllClient() {
	s.Close = true
	s.clientMap.Iterator(func(uid interface{}, c interface{}) bool {
		err := c.(*client).ws.Close()
		return err == nil
	})
	s.clientMap.Clear()
}
