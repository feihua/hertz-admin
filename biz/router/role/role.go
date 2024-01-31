// Code generated by hertz generator. DO NOT EDIT.

package role

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	role "hertz_admin/biz/handler/role"
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
		_api.POST("/query_role_menu", append(_queryrolemenuMw(), role.QueryRoleMenu)...)
		_api.POST("/role_delete", append(_roledeleteMw(), role.RoleDelete)...)
		_api.POST("/role_list", append(_rolelistMw(), role.RoleList)...)
		_api.POST("/role_save", append(_rolesaveMw(), role.RoleSave)...)
		_api.POST("/role_update", append(_roleupdateMw(), role.RoleUpdate)...)
		_api.POST("/update_role_menu", append(_updaterolemenuMw(), role.UpdateRoleMenu)...)
	}
}
