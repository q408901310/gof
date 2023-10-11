package service

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"gof/internal/consts"
	"gof/internal/model"
	"gof/internal/model/entity"
)

const (
	sessionKeyUser         = "SessionKeyUser"    // 用户信息存放在Session中的Key
	sessionKeyLoginReferer = "SessionKeyReferer" // Referer存储，当已存在该session时不会更新。用于用户未登录时引导用户登录，并在登录后跳转到登录前页面。
	sessionKeyNotice       = "SessionKeyNotice"  // 存放在Session中的提示信息，往往使用后则删除
)

type (
	sSession struct{}
)

var (
	insSession = sSession{}
)

func Session() *sSession {
	return &insSession
}

// 设置用户Session.
func (s *sSession) SetUser(ctx context.Context, user *model.SessionUser) (err error) {
	r := g.RequestFromCtx(ctx)
	if err = r.Session.Set(consts.UserSessionKey, user); err != nil {
		return
	}
	return
}

// 获取当前登录的用户信息对象，如果用户未登录返回nil。
func (s *sSession) GetUserBySessionId(ctx context.Context, sessionId string) *entity.User {
	r := g.RequestFromCtx(ctx)
	_ = r.Session.SetId(sessionId)
	v, _ := r.Session.Get(consts.UserSessionKey)
	if !v.IsNil() {
		var user *entity.User
		_ = v.Struct(&user)
		return user
	}
	return nil
}

// 获取当前登录的用户信息对象，如果用户未登录返回nil。
func (s *sSession) GetUser(ctx context.Context) *entity.User {
	customCtx := BizCtx().Get(ctx)
	if customCtx != nil {
		v, _ := customCtx.Session.Get(sessionKeyUser)
		if !v.IsNil() {
			var user *entity.User
			_ = v.Struct(&user)
			return user
		}
	}
	return &entity.User{}
}

// 删除用户Session。
func (s *sSession) RemoveUser(ctx context.Context) error {
	customCtx := BizCtx().Get(ctx)
	if customCtx != nil {
		return customCtx.Session.Remove(sessionKeyUser)
	}
	return nil
}

func (s *sSession) SetSessionId(ctx context.Context, id string) error {
	return g.RequestFromCtx(ctx).Session.SetId(id)
}
