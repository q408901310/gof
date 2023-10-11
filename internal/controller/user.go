package controller

import (
	"context"
	v1 "gof/api/v1"
	"gof/internal/model"
	"gof/internal/service"
)

var (
	User = cUser{}
)

type cUser struct{}

func (c *cUser) Register(ctx context.Context, req *v1.UserRegisterReq) (res *v1.UserRegisterRes, err error) {
	err = service.User().Register(ctx, &model.UserRegisterIn{
		Channel:  req.Channel,
		Passport: req.Passport,
		Password: req.Password,
		Confirm:  req.Confirm,
	})
	if err != nil {
		return nil, err
	}
	return
}

func (c *cUser) Login(ctx context.Context, req *v1.UserLoginReq) (res *v1.UserLoginRes, err error) {
	res = &v1.UserLoginRes{}
	err = service.User().Login(ctx, &model.UserLoginIn{
		Passport: req.Passport,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	res.List, err = service.Server().List(ctx)
	return
}

func (c *cUser) Guest(ctx context.Context, req *v1.UserGuestReq) (res *v1.UserGuestRes, err error) {
	err = service.User().Guest(ctx)
	return
}
