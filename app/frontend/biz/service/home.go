package service

import (
	"context"

	frontend_home "GoMall/app/frontend/hertz_gen/frontend_home"
	"github.com/cloudwego/hertz/pkg/app"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *frontend_home.Empty) (map[string]any, error) {

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
