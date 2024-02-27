// Code generated by hertz generator. DO NOT EDIT.

package Menu

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	menu "github.com/feihua/hertz-admin/biz/handler/menu"
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
			_system := _api.Group("/system", _systemMw()...)
			{
				_menu := _system.Group("/menu", _menuMw()...)
				_menu.POST("/addMenu", append(_menusaveMw(), menu.MenuSave)...)
				_menu.POST("/deleteMenu", append(_menudeleteMw(), menu.MenuDelete)...)
				_menu.GET("/queryMenuList", append(_menulistMw(), menu.MenuList)...)
				_menu.POST("/updateMenu", append(_menuupdateMw(), menu.MenuUpdate)...)
			}
		}
	}
}
