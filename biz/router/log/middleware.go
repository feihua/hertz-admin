// Code generated by hertz generator.

package Log

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/feihua/hertz-admin/biz/pkg/mw"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _apiMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _logMw() []app.HandlerFunc {
	var hf []app.HandlerFunc
	hf = append(hf, mw.JwtMiddleware.MiddlewareFunc())
	return hf
}

func _loginMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _loginlogdeleteMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _queryloginloglistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _statisticsloginlogMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _operateMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _operatelogdeleteMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _queryoperateloglistMw() []app.HandlerFunc {
	// your code...
	return nil
}