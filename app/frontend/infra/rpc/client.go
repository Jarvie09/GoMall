package rpc

import (
	"GoMall/rpc_gen/kitex_gen/product/productcatalogservice"
	"GoMall/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	"github.com/kitex-contrib/registry-nacos/v2/resolver"
	"sync"
)

var (
	UserClient    userservice.Client
	ProductClient productcatalogservice.Client
	//保证只能初始化一次
	once sync.Once
)

func Init() {
	once.Do(func() {
		iniUserClient()
		iniProductClient()
	})
}
func iniUserClient() {
	r, err := resolver.NewDefaultNacosResolver()
	if err != nil {
		hlog.Fatal(err)
	}
	UserClient, err = userservice.NewClient("user", client.WithResolver(r))
	if err != nil {
		hlog.Fatal(err)
	}

}
func iniProductClient() {
	r, err := resolver.NewDefaultNacosResolver()
	if err != nil {
		hlog.Fatal(err)
	}
	UserClient, err = userservice.NewClient("product", client.WithResolver(r))
	if err != nil {
		hlog.Fatal(err)
	}

}
