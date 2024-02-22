// Code generated by hertz generator. DO NOT EDIT.

package menu

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	menu "hertz_admin/biz/handler/menu"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_api := root.Group("/api", _apiMw()...)
		{
			_menu := _api.Group("/menu", _menuMw()...)
			_menu.POST("/menuDelete", append(_menudeleteMw(), menu.MenuDelete)...)
			_menu.GET("/menuList", append(_menulistMw(), menu.MenuList)...)
			_menu.POST("/menuSave", append(_menusaveMw(), menu.MenuSave)...)
			_menu.POST("/menuUpdate", append(_menuupdateMw(), menu.MenuUpdate)...)
		}
	}
}
