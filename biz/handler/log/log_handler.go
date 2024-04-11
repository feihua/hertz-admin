// Code generated by hertz generator.

package log

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/feihua/hertz-admin/biz/dal"
	"github.com/feihua/hertz-admin/biz/model/api"
	"github.com/feihua/hertz-admin/gen/query"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	log "github.com/feihua/hertz-admin/biz/model/log"
)

// LoginLogDelete .
// @router /api/log/login/deleteLoginLog [POST]
func LoginLogDelete(ctx context.Context, c *app.RequestContext) {
	resp := new(log.LoginLogDeleteResp)
	var err error
	var req log.LoginLogDeleteReq
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.Code = api.Code_ParamInvalid
		resp.Msg = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	loginLog := query.SysLoginLog
	_, err = loginLog.WithContext(ctx).Where(loginLog.ID.In(req.Ids...)).Delete()
	if err != nil {
		hlog.CtxErrorf(ctx, "删除登录日志异常: %v", err)
		resp.Msg = err.Error()
		resp.Code = api.Code_DBErr
		c.JSON(http.StatusOK, resp)
		return
	}

	resp.Code = api.Code_Success
	resp.Msg = "删除登录日志成功"

	c.JSON(http.StatusOK, resp)
}

// QueryLoginLogList .
// @router /api/log/login/queryLoginLogList [POST]
func QueryLoginLogList(ctx context.Context, c *app.RequestContext) {
	resp := new(log.QueryLoginLogListResp)
	var err error
	var req log.QueryLoginLogListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.Code = api.Code_ParamInvalid
		resp.Msg = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	result, count, err := dal.QueryLoginLogList(req.UserName, req.IP, req.PageNo, req.PageSize)

	if err != nil {
		hlog.CtxErrorf(ctx, "查询登录日志列表异常: %v", err)
		resp.Code = api.Code_DBErr
		resp.Msg = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	//返回数据
	var list []*log.LoginLogListData

	for _, item := range result {
		list = append(list, &log.LoginLogListData{
			Id:        int32(item.ID),
			UserName:  item.UserName,
			Status:    item.Status,
			Ip:        item.IP,
			LoginTime: item.LoginTime.Format("2006-01-02 15:04:05"),
		})
	}

	resp.Code = api.Code_Success
	resp.Msg = "查询登录日志列表成功"
	resp.Data = list
	resp.Total = count

	c.JSON(http.StatusOK, resp)
}

// StatisticsLoginLog .
// @router /api/log/login/statisticsLoginLog [GET]
func StatisticsLoginLog(ctx context.Context, c *app.RequestContext) {
	resp := new(log.StatisticsLoginLogResp)
	var err error
	var req log.StatisticsLoginLogReq
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.Code = api.Code_ParamInvalid
		resp.Msg = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	resp.Code = api.Code_Success
	resp.Msg = "查询登录日志统计成功"
	resp.Data = dal.QueryStatisticsLoginLog()

	c.JSON(http.StatusOK, resp)
}

// OperateLogDelete .
// @router /api/log/operate/deleteOperateLog [POST]
func OperateLogDelete(ctx context.Context, c *app.RequestContext) {
	resp := new(log.OperateLogDeleteResp)
	var err error
	var req log.OperateLogDeleteReq
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.Code = api.Code_ParamInvalid
		resp.Msg = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	operateLog := query.SysOperateLog
	_, err = operateLog.WithContext(ctx).Where(operateLog.ID.In(req.Ids...)).Delete()
	if err != nil {
		hlog.CtxErrorf(ctx, "删除操作日志异常: %v", err)
		resp.Msg = err.Error()
		resp.Code = api.Code_DBErr
		c.JSON(http.StatusOK, resp)
		return
	}

	resp.Code = api.Code_Success
	resp.Msg = "删除操作日志成功"

	c.JSON(http.StatusOK, resp)
}

// QueryOperateLogList .
// @router /api/log/operate/queryOperateLogList [POST]
func QueryOperateLogList(ctx context.Context, c *app.RequestContext) {
	resp := new(log.QueryOperateLogListResp)
	var err error
	var req log.QueryOperateLogListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.Code = api.Code_ParamInvalid
		resp.Msg = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	result, count, err := dal.QueryOperateLogList(req.UserName, req.IP, req.Method, req.PageNo, req.PageSize)

	if err != nil {
		hlog.CtxErrorf(ctx, "查询操作日志列表异常: %v", err)
		resp.Code = api.Code_DBErr
		resp.Msg = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	//返回数据
	var list []*log.OperateLogListData

	for _, item := range result {
		list = append(list, &log.OperateLogListData{
			Id:             int32(item.ID),
			UserName:       item.UserName,
			Operation:      item.Operation,
			Method:         item.Method,
			RequestParams:  item.RequestParams,
			ResponseParams: *item.ResponseParams,
			Time:           int32(item.Time),
			Ip:             *item.IP,
			OperationTime:  item.OperationTime.Format("2006-01-02 15:04:05"),
		})
	}

	resp.Code = api.Code_Success
	resp.Msg = "查询操作日志列表成功"
	resp.Data = list
	resp.Total = count

	c.JSON(http.StatusOK, resp)
}