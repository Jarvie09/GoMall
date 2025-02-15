package service

import (
	"GoMall/app/frontend/infra/rpc"
	"GoMall/rpc_gen/kitex_gen/user"
	"context"
	"fmt"
	"github.com/hertz-contrib/sessions"

	frontend_auth "GoMall/app/frontend/hertz_gen/frontend_auth"
	"github.com/cloudwego/hertz/pkg/app"
)

type RegisterService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewRegisterService(Context context.Context, RequestContext *app.RequestContext) *RegisterService {
	return &RegisterService{RequestContext: RequestContext, Context: Context}
}

func (h *RegisterService) Run(req *frontend_auth.RegisterReq) (resp *frontend_auth.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	//对接userService对接
	fmt.Println("开始RPC调用User服务")
	userResp, err := rpc.UserClient.Register(h.Context, &user.RegisterReq{
		Email:           req.Email,
		Password:        req.Password,
		PasswordConform: req.PasswordConfirm,
	})
	if userResp == nil {
		fmt.Println("RPC调用User服务失败")
		fmt.Println("userResp == nil")
		fmt.Println("userResp为空")

	}
	session := sessions.Default(h.RequestContext)
	session.Set("user_id", userResp.UserId)
	fmt.Println("register:", userResp.UserId)
	err = session.Save()
	if err != nil {
		return nil, err
	}
	return
}
