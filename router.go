// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/feihua/hertz-admin/biz/handler"
	"github.com/feihua/hertz-admin/biz/pkg/mw"
)

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {
	r.POST("/api/system/user/login", mw.JwtMiddleware.LoginHandler)
	//auth := r.Group("/auth", mw.JwtMiddleware.MiddlewareFunc())
	//auth.GET("/ping", handler.Ping)
	r.GET("/ping", handler.Ping)

	// your code ...
}
