package http

import (
	"github.com/bilibili/kratos/pkg/net/http/blademaster/render"
	"github.com/wq1019/short/internal/service"
	"net/http"

	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/log"
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
)

var (
	svc *service.Service
)

// New new a bm server.
func New(s *service.Service) (engine *bm.Engine) {
	var (
		hc struct {
			Server *bm.ServerConfig
		}
	)
	if err := paladin.Get("http.toml").UnmarshalTOML(&hc); err != nil {
		if err != paladin.ErrNotExist {
			panic(err)
		}
	}
	svc = s
	engine = bm.DefaultServer(hc.Server)
	initRouter(engine)
	if err := engine.Start(); err != nil {
		panic(err)
	}
	return
}

func initRouter(e *bm.Engine) {
	e.Ping(ping)
	g := e.Group("/demo")
	{
		//g.GET("/start", howToStart)
	}
	api := e.Group("/api/v1")
	{
		api.GET("/test", func(c *bm.Context) {
			c.Render(200, render.JSON{Code: http.StatusOK, Message: "hello", TTL: 3600, Data: "33"})
		})
		api.GET("/user:id", QueryUserInfo)

	}
}

func ping(ctx *bm.Context) {
	if err := svc.Ping(ctx); err != nil {
		log.Error("ping error(%v)", err)
		ctx.AbortWithStatus(http.StatusServiceUnavailable)
	}
}

//// example for http request handler.
//func howToStart(c *bm.Context) {
//	k := &model.Kratos{
//		Hello: "Golang 大法好 !!!",
//	}
//	c.JSON(k, nil)
//}