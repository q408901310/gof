package service

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/gogo/protobuf/proto"
	"gof/internal/action"
	"gof/internal/model/custom/merror"
	"gof/internal/module"
	"gof/internal/pb"

	"github.com/gogf/gf/v2/container/gmap"

	"github.com/gogf/gf/v2/frame/g"
)

type (
	sWebsocket struct {
		clientMap *gmap.Map
		pool      *grpool.Pool
		Close     bool
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
	c := module.NewClient(ws, user.Id)
	err = s.AddClient(c)
	if err != nil {
		return
	}
	for {
		// Blocking reading message from current websocket.
		msgType, msgByte, err = ws.ReadMessage()
		if err != nil {
			// Remove session info.
			s.clientMap.Remove(ws)
			return gerror.New("读取数据错误")
		}
		c.MsgType = msgType
		err = proto.Unmarshal(msgByte, &msg)
		if err != nil {
			// TODO 记录错误
			c.SendErrorMessage(msg.Action, pb.CODE_SYS_UNMARSHAL)
			continue
		}
		actCtx, cancel := Context().InitAction(ctx, c)
		c.CtxCancel = cancel
		// 子线程执行
		err = s.pool.Add(ctx, func(ctx context.Context) {
			action.Run(actCtx, msg.Action, msg.DataMsg)
		})
		if err != nil {
			// TODO 记录错误
			c.SendErrorMessage(msg.Action, pb.CODE_SYS_UNKNOWN)
			continue
		}
	}
}

func (s *sWebsocket) AddClient(c *module.Client) error {
	if s.Close {
		//return gerror.New("维护中")
		return merror.New("维护中")
	}
	// Create session data for current websocket.
	s.clientMap.Set(c.Uid, c)
	return nil
}

func (s *sWebsocket) CloseClient(c *module.Client) {
	s.clientMap.Remove(c.Uid)
	err := c.Close()
	if err != nil {
		return
	}
}

func (s *sWebsocket) CloseAllClient() {
	s.Close = true
	s.clientMap.Iterator(func(uid interface{}, c interface{}) bool {
		err := c.(*module.Client).Close()
		return err == nil
	})
	s.clientMap.Clear()
}
