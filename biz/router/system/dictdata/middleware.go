// Code generated by hertz generator.

package dictdata

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

func _dictdataMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _adddictdataMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _deletedictdataMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _querydictdatadetailMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _querydictdatalistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _updatedictdataMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _updatedictdatastatusMw() []app.HandlerFunc {
	// your code...
	return nil
}
