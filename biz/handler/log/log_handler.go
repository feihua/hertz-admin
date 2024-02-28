// Code generated by hertz generator.

package log

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	log "github.com/feihua/hertz-admin/biz/model/log"
)

// LoginLogDelete .
// @router /api/log/login/deleteLoginLog [POST]
func LoginLogDelete(ctx context.Context, c *app.RequestContext) {
	var err error
	var req log.LoginLogDeleteReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	resp := new(log.LoginLogDeleteResp)

	c.JSON(200, resp)
}

// QueryLoginLogList .
// @router /api/log/login/queryLoginLogList [POST]
func QueryLoginLogList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req log.QueryLoginLogListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	resp := new(log.QueryLoginLogListResp)

	c.JSON(200, resp)
}

// StatisticsLoginLog .
// @router /api/log/login/statisticsLoginLog [GET]
func StatisticsLoginLog(ctx context.Context, c *app.RequestContext) {
	var err error
	var req log.StatisticsLoginLogReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	resp := new(log.StatisticsLoginLogResp)

	c.JSON(200, resp)
}

// OperateLogDelete .
// @router /api/log/operate/deleteOperateLog [POST]
func OperateLogDelete(ctx context.Context, c *app.RequestContext) {
	var err error
	var req log.OperateLogDeleteReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	resp := new(log.OperateLogDeleteResp)

	c.JSON(200, resp)
}

// QueryOperateLogList .
// @router /api/log/operate/queryOperateLogList [POST]
func QueryOperateLogList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req log.QueryOperateLogListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	resp := new(log.QueryOperateLogListResp)

	c.JSON(200, resp)
}
