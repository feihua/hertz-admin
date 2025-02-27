// Code generated by hertz generator. DO NOT EDIT.

package dictdata

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	dictdata "github.com/feihua/hertz-admin/biz/handler/system/dictdata"
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
				_dictdata := _system.Group("/dictData", _dictdataMw()...)
				_dictdata.POST("/addDictData", append(_adddictdataMw(), dictdata.AddDictData)...)
				_dictdata.POST("/deleteDictData", append(_deletedictdataMw(), dictdata.DeleteDictData)...)
				_dictdata.POST("/queryDictDataDetail", append(_querydictdatadetailMw(), dictdata.QueryDictDataDetail)...)
				_dictdata.POST("/queryDictDataList", append(_querydictdatalistMw(), dictdata.QueryDictDataList)...)
				_dictdata.POST("/updateDictData", append(_updatedictdataMw(), dictdata.UpdateDictData)...)
				_dictdata.POST("/updateDictDataStatus", append(_updatedictdatastatusMw(), dictdata.UpdateDictDataStatus)...)
			}
		}
	}
}
