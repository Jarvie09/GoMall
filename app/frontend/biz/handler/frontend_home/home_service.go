package frontend_home

import (
	"GoMall/app/frontend/biz/service"
	"GoMall/app/frontend/biz/utils"
	frontend_home "GoMall/app/frontend/hertz_gen/frontend_home"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Home .
// @router / [GET]
func Home(ctx context.Context, c *app.RequestContext) {
	var err error
	var req frontend_home.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	//resp := &frontend_home.Empty{}
	resp, err := service.NewHomeService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.HTML(consts.StatusOK, "home", utils.WarpResponse(ctx, c, resp))
	//utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
