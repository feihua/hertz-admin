// Code generated by hertz generator.

package user

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/feihua/hertz_admin/biz/pkg/mw"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _apiMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _userMw() []app.HandlerFunc {
	var hf []app.HandlerFunc
	hf = append(hf, mw.JwtMiddleware.MiddlewareFunc())
	return hf
}

func _queryusermenuMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _queryuserroleMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _updateuserroleMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _userdeleteMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _userlistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _usersaveMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _userupdateMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _query_user_roleMw() []app.HandlerFunc {
	// your code...
	return nil
}
