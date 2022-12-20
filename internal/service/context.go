package service

import (
	"context"
	"gof/internal/consts"
	"gof/internal/model"
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

// InitAction 初始化上下文对象指针到上下文对象中，以便后续的请求流程中可以修改。
func (s *sContext) InitAction(parent context.Context, c *model.Client) (ctx context.Context, cancel context.CancelFunc) {
	ctx = context.WithValue(parent, consts.ContextClient, c)
	ctx, cancel = context.WithCancel(ctx)
	return
}

func (s *sContext) GetClient(ctx context.Context) (c *model.Client) {
	return ctx.Value(consts.ContextClient).(*model.Client)
}
