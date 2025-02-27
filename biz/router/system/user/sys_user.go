// Code generated by hertz generator. DO NOT EDIT.

package user

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	user "github.com/feihua/hertz-admin/biz/handler/system/user"
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
				_user := _system.Group("/user", _userMw()...)
				_user.POST("/addUser", append(_adduserMw(), user.AddUser)...)
				_user.POST("/deleteUser", append(_deleteuserMw(), user.DeleteUser)...)
				_user.POST("/queryUserDetail", append(_queryuserdetailMw(), user.QueryUserDetail)...)
				_user.POST("/queryUserList", append(_queryuserlistMw(), user.QueryUserList)...)
				_user.GET("/queryUserMenu", append(_queryusermenuMw(), user.QueryUserMenu)...)
				_user.POST("/queryUserRole", append(_queryuserrolelistMw(), user.QueryUserRoleList)...)
				_user.POST("/updateUser", append(_updateuserMw(), user.UpdateUser)...)
				_user.POST("/updateUserRole", append(_adduserroleMw(), user.AddUserRole)...)
				_user.POST("/updateUserStatus", append(_updateuserstatusMw(), user.UpdateUserStatus)...)
			}
		}
	}
}
