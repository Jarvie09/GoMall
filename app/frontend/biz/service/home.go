package service

import (
	frontend_home "GoMall/app/frontend/hertz_gen/frontend_home"
	"GoMall/app/frontend/infra/rpc"
	"GoMall/rpc_gen/kitex_gen/product"
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *frontend_home.Empty) (map[string]any, error) {
	fmt.Println("开始商品服务RPC调用")
	p, err := rpc.ProductClient.ListProducts(h.Context, &product.ListProductsReq{})
	if err != nil {
		fmt.Println(err)
		fmt.Println(p)
	}
	if p == nil {
		fmt.Println("P这里有问题")
		fmt.Println("商品服务有问题")
	}
	fmt.Println(p.Products)
	return utils.H{"title": "hot sale", "item": p.Products}, nil
	fmt.Println("这里有问题")
	item := []map[string]any{
		{"Name": "T-shirt", "Price": 100, "Picture": "/static/imag/logo.jpg"},
		{"Name": "T-shirt-1", "Price": 110, "Picture": "/static/imag/logo.jpg"},
		{"Name": "T-shirt-2", "Price": 120, "Picture": "/static/imag/logo.jpg"},
		{"Name": "T-shirt-3", "Price": 139, "Picture": "/static/imag/logo.jpg"},
		{"Name": "T-shirt-4", "Price": 140, "Picture": "/static/imag/logo.jpg"},
		{"Name": "T-shirt-5", "Price": 150, "Picture": "/static/imag/logo.jpg"},
	}
	resp := make(map[string]any)
	resp["item"] = item
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return resp, nil
}
