// Code generated by hertz generator.

package notice

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/feihua/hertz-admin/biz/pkg/mw"
)

func rootMw() []app.HandlerFunc {
	var hf []app.HandlerFunc
	hf = append(hf, mw.JwtMiddleware.MiddlewareFunc())
	return hf
}

func _apiMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _systemMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _noticeMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _addnoticeMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _deletenoticeMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _querynoticedetailMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _querynoticelistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _updatenoticeMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _updatenoticestatusMw() []app.HandlerFunc {
	// your code...
	return nil
}
