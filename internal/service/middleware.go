package service

import (
	"net/http"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	sMiddleware struct{}
)

var (
	insMiddleware = sMiddleware{}
)

func Middleware() *sMiddleware {
	return &insMiddleware
}

// Debug 前台系统权限控制，用户必须登录才能访问
func (s *sMiddleware) Debug(r *ghttp.Request) {
	g.Dump(r.Session.Id())
	token := r.Get("token")
	if token.String() == "123456" {
		r.Middleware.Next()
	} else {
		r.Response.WriteStatusExit(http.StatusNotFound)
	}
	r.Middleware.Next()
}

// Auth 前台系统权限控制，用户必须登录才能访问
func (s *sMiddleware) Auth(r *ghttp.Request) {
	token := r.Get("token")
	if token.String() == "123456" {
		r.Middleware.Next()
	} else {
		r.Response.WriteStatusExit(http.StatusForbidden)
	}
	r.Middleware.Next()
}
