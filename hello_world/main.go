package main

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func main() {
	// server.Default() creates a Hertz with recovery middleware.
	// If you need a pure hertz, you can use server.New()
	h := server.Default()
	//hertz := server.New()

	h.GET("/hello", func(ctx context.Context, c *app.RequestContext) {
		//c.String(consts.StatusOK, "Hello hertz!")
		//c.JSON(200,"{\"status\":\"ok\"}")
		c.Data(consts.StatusOK, consts.MIMETextPlain, []byte("hello world"))
	})

	h.Spin()
}
