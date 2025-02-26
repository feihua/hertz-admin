// Code generated by hertz generator. DO NOT EDIT.

package role

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	role "github.com/feihua/hertz-admin/biz/handler/system/role"
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
				_role := _system.Group("/role", _roleMw()...)
				_role.POST("/addRole", append(_addroleMw(), role.AddRole)...)
				_role.POST("/batchAuthUser", append(_batchauthuserMw(), role.BatchAuthUser)...)
				_role.POST("/batchCancelAuthUser", append(_batchcancelauthuserMw(), role.BatchCancelAuthUser)...)
				_role.POST("/cancelAuthUser", append(_cancelauthuserMw(), role.CancelAuthUser)...)
				_role.POST("/deleteRole", append(_deleteroleMw(), role.DeleteRole)...)
				_role.POST("/queryAllocatedList", append(_queryallocatedlistMw(), role.QueryAllocatedList)...)
				_role.POST("/queryRoleDetail", append(_queryroledetailMw(), role.QueryRoleDetail)...)
				_role.POST("/queryRoleList", append(_queryrolelistMw(), role.QueryRoleList)...)
				_role.POST("/queryRoleMenu", append(_queryrolemenulistMw(), role.QueryRoleMenuList)...)
				_role.POST("/queryUnallocatedList", append(_queryunallocatedlistMw(), role.QueryUnallocatedList)...)
				_role.POST("/updateRole", append(_updateroleMw(), role.UpdateRole)...)
				_role.POST("/updateRoleMenu", append(_addrolemenuMw(), role.AddRoleMenu)...)
				_role.POST("/updateRoleStatus", append(_updaterolestatusMw(), role.UpdateRoleStatus)...)
			}
		}
	}
}
