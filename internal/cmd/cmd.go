package cmd

import (
	"context"
	"gof/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"gof/internal/controller"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			// 静态目录设置
			uploadPath := g.Cfg().MustGet(ctx, "upload.path").String()
			if uploadPath == "" {
				g.Log().Fatal(ctx, "文件上传配置路径不能为空")
			}
			s.AddStaticPath("/upload", uploadPath)

			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareCORS, ghttp.MiddlewareHandlerResponse)
				group.Bind(
					controller.User,
				)

				//登录后可访问
				s.Group("/", func(group *ghttp.RouterGroup) {
					group.Bind(
						controller.Websocket,
					)
				})

				//登录后可访问
				s.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(service.Middleware().Auth)
					group.Bind(
						controller.Server,
					)
				})
			})

			//平滑重启
			s.BindMiddleware("/debug/*", service.Middleware().Debug)
			s.EnableAdmin()
			s.Run()
			return nil
		},
	}
)
