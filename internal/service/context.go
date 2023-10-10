package service

import (
	"context"
	"gof/internal/consts"
	"gof/internal/module"
)

type (
	sContext struct{}
)

var (
	insContext = sContext{}
)

func Context() *sContext {
	return &insContext
}

// InitAction 初始化上下文，并将链接对象保存到上下文中，以便后续流程中使用。
func (s *sContext) InitAction(parent context.Context, c *module.Client) (ctx context.Context, cancel context.CancelFunc) {
	ctx = context.WithValue(parent, consts.ContextClient, c)
	ctx, cancel = context.WithCancel(ctx)
	return
}

func (s *sContext) GetClient(ctx context.Context) (c *module.Client) {
	return ctx.Value(consts.ContextClient).(*module.Client)
}
