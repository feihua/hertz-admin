// Code generated by hertz generator.

package Sys

import (
	"github.com/cloudwego/hertz/pkg/app"
	"hertz_admin/biz/pkg/mw"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _menudeleteMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _menulistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _menusaveMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _menuupdateMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _queryrolemenuMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _queryusermenuMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _roledeleteMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _rolelistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _rolesaveMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _roleupdateMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _updaterolemenuMw() []app.HandlerFunc {
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

func _apiMw() []app.HandlerFunc {
	// your code...
	var hf []app.HandlerFunc
	hf = append(hf, mw.JwtMiddleware.MiddlewareFunc())
	return hf
}
