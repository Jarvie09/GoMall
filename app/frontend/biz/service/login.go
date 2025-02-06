package service

import (
	"context"
	"github.com/hertz-contrib/sessions"

	frontend_auth "GoMall/app/frontend/hertz_gen/frontend_auth"
	"github.com/cloudwego/hertz/pkg/app"
)

type LoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLoginService(Context context.Context, RequestContext *app.RequestContext) *LoginService {
	return &LoginService{RequestContext: RequestContext, Context: Context}
}

func (h *LoginService) Run(req *frontend_auth.LoginReq) (redirect string, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	//登录态，引用session
	session := sessions.Default(h.RequestContext)
	//var count int
	//v := session.Get("count")
	//if v == nil {
	//	count = 0
	//} else {
	//	count = v.(int)
	//	count++
	//}
	session.Set("user_id", 1)
	err = session.Save()
	if err != nil {
		return "", err
	}
	if req.Next != "" {
		redirect = req.Next
	}
	return
}
