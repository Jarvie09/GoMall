package middleware

import "github.com/cloudwego/hertz/pkg/app/server"

func Register(he *server.Hertz) {
	he.Use(GlobalAuth())
}
