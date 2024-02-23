package cmd

import (
	"context"
	"github.com/lshuangquan/yunbo-demo/internal/controller/file"
	"github.com/lshuangquan/yunbo-demo/internal/controller/page"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.BindHandler("GET:/file/view", file.New().View)

			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					file.New(),
				)
			})

			s.Group("/page", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					page.New(),
				)
			})

			s.Run()
			return nil
		},
	}
)
