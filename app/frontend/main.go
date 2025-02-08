// Code generated by hertz generator.

package main

import (
	"GoMall/app/frontend/infra/rpc"
	"GoMall/app/frontend/middleware"
	"context"
	"github.com/hertz-contrib/sessions"
	"github.com/hertz-contrib/sessions/redis"
	"github.com/joho/godotenv"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"

	"GoMall/app/frontend/biz/router"
	"GoMall/app/frontend/conf"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/cors"
	"github.com/hertz-contrib/gzip"
	"github.com/hertz-contrib/logger/accesslog"
	hertzlogrus "github.com/hertz-contrib/logger/logrus"
	"github.com/hertz-contrib/pprof"
	"go.uber.org/zap/zapcore"
)

func main() {

	_ = godotenv.Load()
	rpc.Init()
	// init dal
	// dal.Init()
	address := conf.GetConf().Hertz.Address
	h := server.New(server.WithHostPorts(address))

	registerMiddleware(h)

	//Nacos注册中心
	//sc := []constant.ServerConfig{
	//	*constant.NewServerConfig("127.0.0.1", 8848),
	//}

	// add a ping route to test
	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{"ping": "pong"})
	})

	router.GeneratedRegister(h)
	h.LoadHTMLGlob("template/*")
	h.Static("/static", "./")
	//登录
	h.GET("/sign-in", func(c context.Context, ctx *app.RequestContext) {
		data := utils.H{"Title": "Sign In", "Next": ctx.Query("next")}
		ctx.HTML(consts.StatusOK, "sign-in.tmpl", data)
	})
	//注册
	h.GET("/sign-up", func(c context.Context, ctx *app.RequestContext) {
		ctx.HTML(consts.StatusOK, "sign-up.tmpl", utils.H{"Title": "Sign Up"})
	})
	//关于
	h.GET("/about", func(c context.Context, ctx *app.RequestContext) {
		ctx.HTML(consts.StatusOK, "about.tmpl", utils.H{"Title": "About"})
	})
	h.Spin()
}

func registerMiddleware(h *server.Hertz) {
	store, _ := redis.NewStore(10, "tcp", conf.GetConf().Redis.Address, "", []byte(os.Getenv("SESSION_SECRET")))
	h.Use(sessions.New("cloudwego-shop", store))
	// log
	logger := hertzlogrus.NewLogger()
	hlog.SetLogger(logger)
	hlog.SetLevel(conf.LogLevel())
	asyncWriter := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.GetConf().Hertz.LogFileName,
			MaxSize:    conf.GetConf().Hertz.LogMaxSize,
			MaxBackups: conf.GetConf().Hertz.LogMaxBackups,
			MaxAge:     conf.GetConf().Hertz.LogMaxAge,
		}),
		FlushInterval: time.Minute,
	}
	hlog.SetOutput(asyncWriter)
	h.OnShutdown = append(h.OnShutdown, func(ctx context.Context) {
		asyncWriter.Sync()
	})

	// pprof
	if conf.GetConf().Hertz.EnablePprof {
		pprof.Register(h)
	}

	// gzip
	if conf.GetConf().Hertz.EnableGzip {
		h.Use(gzip.Gzip(gzip.DefaultCompression))
	}

	// access log
	if conf.GetConf().Hertz.EnableAccessLog {
		h.Use(accesslog.New())
	}

	// recovery
	h.Use(recovery.Recovery())

	// cores
	h.Use(cors.Default())
	middleware.Register(h)
}
