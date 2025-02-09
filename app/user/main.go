package main

import (
	"GoMall/app/user/biz/dal"
	"github.com/joho/godotenv"
	"github.com/kitex-contrib/registry-nacos/v2/example/hello/kitex_gen/api"
	"github.com/kitex-contrib/registry-nacos/v2/registry"
	"golang.org/x/net/context"
	//"github.com/kitex-contrib/registry-nacos/v2/registry"
	"gopkg.in/natefinch/lumberjack.v2"
	"net"
	"time"

	"GoMall/app/user/conf"
	"GoMall/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"go.uber.org/zap/zapcore"
)

type HelloImpl struct{}

func (h *HelloImpl) Echo(_ context.Context, req *api.Request) (resp *api.Response, err error) {
	resp = &api.Response{
		Message: req.Message,
	}
	return
}
func main() {
	_ = godotenv.Load("./.env")
	dal.Init()
	opts := kitexInit()

	svr := userservice.NewServer(new(UserServiceImpl), opts...)

	err := svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}
}

func kitexInit() (opts []server.Option) {
	// address
	addr, err := net.ResolveTCPAddr("tcp", conf.GetConf().Kitex.Address)
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithServiceAddr(addr))

	// service info
	opts = append(opts, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: conf.GetConf().Kitex.Service,
	}))

	//注册中心  Nacos
	r, err := registry.NewDefaultNacosRegistry()
	if err != nil {
		panic(err)
	}
	//
	//svr := hello.NewServer(
	//	new(HelloImpl),
	//	server.WithRegistry(r),
	//	server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "user"}),
	//	server.WithServiceAddr(&net.TCPAddr{IP: net.IPv4(192, 168, 2, 21), Port: 8888}),
	//)
	//if err := svr.Run(); err != nil {
	//	log.Println("server stopped with error:", err)
	//} else {
	//	log.Println("server stopped")
	//}
	opts = append(opts, server.WithRegistry(r))
	// klog
	logger := kitexlogrus.NewLogger()
	klog.SetLogger(logger)
	klog.SetLevel(conf.LogLevel())
	asyncWriter := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.GetConf().Kitex.LogFileName,
			MaxSize:    conf.GetConf().Kitex.LogMaxSize,
			MaxBackups: conf.GetConf().Kitex.LogMaxBackups,
			MaxAge:     conf.GetConf().Kitex.LogMaxAge,
		}),
		FlushInterval: time.Minute,
	}
	klog.SetOutput(asyncWriter)
	server.RegisterShutdownHook(func() {
		asyncWriter.Sync()
	})
	return
}
