// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	log "github.com/feihua/hertz-admin/biz/router/log"
	menu "github.com/feihua/hertz-admin/biz/router/menu"
	role "github.com/feihua/hertz-admin/biz/router/role"
	user "github.com/feihua/hertz-admin/biz/router/user"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	log.Register(r)

	menu.Register(r)

	role.Register(r)

	user.Register(r)
}
