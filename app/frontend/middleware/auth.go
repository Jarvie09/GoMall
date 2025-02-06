package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
)

type SessionUserIdKey string

const SessionUserId SessionUserIdKey = "user_id"

func GlobalAuth() app.HandlerFunc {
	//返回一个函数
	return func(ctx context.Context, c *app.RequestContext) {
		//获取用户身份相关内容
		session := sessions.Default(c)
		context.WithValue(ctx, SessionUserId, session.Get("user_id"))
		c.Next(ctx)
	}
}

// 鉴权
func Auth() app.HandlerFunc {
	//返回一个函数
	return func(ctx context.Context, c *app.RequestContext) {
		//获取用户身份相关内容
		session := sessions.Default(c)
		userId := session.Get("user_id")
		if userId == nil {
			c.Redirect(302, []byte("/sign-in?next="+c.FullPath()))
			//跳转之后打断后续的操作
			c.Abort()
			return
		}
		c.Next(ctx)
	}
}
