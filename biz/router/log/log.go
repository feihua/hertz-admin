// Code generated by hertz generator. DO NOT EDIT.

package Log

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	log "github.com/feihua/hertz-admin/biz/handler/log"
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
			_log := _api.Group("/log", _logMw()...)
			{
				_login := _log.Group("/login", _loginMw()...)
				_login.POST("/deleteLoginLog", append(_loginlogdeleteMw(), log.LoginLogDelete)...)
				_login.POST("/queryLoginLogList", append(_queryloginloglistMw(), log.QueryLoginLogList)...)
				_login.GET("/statisticsLoginLog", append(_statisticsloginlogMw(), log.StatisticsLoginLog)...)
			}
			{
				_operate := _log.Group("/operate", _operateMw()...)
				_operate.POST("/deleteOperateLog", append(_operatelogdeleteMw(), log.OperateLogDelete)...)
				_operate.POST("/queryOperateLogList", append(_queryoperateloglistMw(), log.QueryOperateLogList)...)
			}
		}
	}
}
