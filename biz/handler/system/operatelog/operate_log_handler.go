package system

import (
	"context"
	"errors"
	operatelog "github.com/feihua/hertz-admin/biz/model/system/operateLog"
	"github.com/feihua/hertz-admin/biz/pkg/utils"
	"github.com/feihua/hertz-admin/gen/query"
	"gorm.io/gorm"

	"github.com/cloudwego/hertz/pkg/app"
)

// DeleteOperateLog 删除操作日志记录
// @router /api/system/operateLog/deleteOperateLog [POST]
func DeleteOperateLog(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var req operatelog.DeleteOperateLogReq
	err := c.BindAndValidate(&req)
	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	q := query.SysOperateLog

	_, err = q.WithContext(ctx).Where(q.ID.In(req.Ids...)).Delete()
	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	resp.Success(c, "删除操作日志记录成功")
}

// QueryOperateLogDetail 查询操作日志记录详情
// @router /api/system/operateLog/queryOperateLogDetail [POST]
func QueryOperateLogDetail(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var req operatelog.QueryOperateLogDetailReq
	err := c.BindAndValidate(&req)
	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	item, err := query.SysOperateLog.WithContext(ctx).Where(query.SysOperateLog.ID.Eq(req.Id)).First()
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		resp.Error(c, "操作日志记录不存在")
		return
	case err != nil:
		resp.Error(c, err.Error())
		return
	}

	data := &operatelog.QueryOperateLogDetailResp{
		Id:              item.ID,                           // 日志主键
		Title:           item.Title,                        // 模块标题
		BusinessType:    item.BusinessType,                 // 业务类型（0其它 1新增 2修改 3删除）
		Method:          item.Method,                       // 方法名称
		RequestMethod:   item.RequestMethod,                // 请求方式
		OperatorType:    item.OperatorType,                 // 操作类别（0其它 1后台用户 2手机端用户）
		OperateName:     item.OperateName,                  // 操作人员
		DeptName:        item.DeptName,                     // 部门名称
		OperateUrl:      item.OperateURL,                   // 请求URL
		OperateIp:       item.OperateIP,                    // 主机地址
		OperateLocation: item.OperateLocation,              // 操作地点
		OperateParam:    item.OperateParam,                 // 请求参数
		JsonResult:      item.JSONResult,                   // 返回参数
		Platform:        item.Platform,                     // 平台信息
		Browser:         item.Browser,                      // 浏览器类型
		Version:         item.Version,                      // 浏览器版本
		Os:              item.Os,                           // 操作系统
		Arch:            item.Arch,                         // 体系结构信息
		Engine:          item.Engine,                       // 渲染引擎信息
		EngineDetails:   item.EngineDetails,                // 渲染引擎详细信息
		Extra:           item.Extra,                        // 其他信息（可选）
		Status:          item.Status,                       // 操作状态(0:异常,正常)
		ErrorMsg:        item.ErrorMsg,                     // 错误消息
		OperateTime:     utils.TimeToStr(item.OperateTime), // 操作时间
		CostTime:        item.CostTime,                     // 消耗时间

	}

	resp.Success(c, data)
}

// QueryOperateLogList 查询操作日志记录列表
// @router /api/system/operateLog/queryOperateLogList [POST]
func QueryOperateLogList(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var err error
	var req operatelog.QueryOperateLogListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	p := query.SysOperateLog
	q := p.WithContext(ctx)
	// 模块标题
	if len(req.Title) > 0 {
		q = q.Where(p.Title.Like("%" + req.Title + "%"))
	}
	// 业务类型（0其它 1新增 2修改 3删除）
	if req.BusinessType != 2 {
		q = q.Where(p.BusinessType.Eq(req.BusinessType))
	}
	// 方法名称
	if len(req.Method) > 0 {
		q = q.Where(p.Method.Like("%" + req.Method + "%"))
	}
	// 请求方式
	if len(req.RequestMethod) > 0 {
		q = q.Where(p.RequestMethod.Like("%" + req.RequestMethod + "%"))
	}
	// 操作类别（0其它 1后台用户 2手机端用户）
	if req.OperatorType != 2 {
		q = q.Where(p.OperatorType.Eq(req.OperatorType))
	}
	// 操作人员
	if len(req.OperateName) > 0 {
		q = q.Where(p.OperateName.Like("%" + req.OperateName + "%"))
	}
	// 部门名称
	if len(req.DeptName) > 0 {
		q = q.Where(p.DeptName.Like("%" + req.DeptName + "%"))
	}

	// 操作地点
	if len(req.OperateLocation) > 0 {
		q = q.Where(p.OperateLocation.Like("%" + req.OperateLocation + "%"))
	}

	// 操作状态(0:异常,正常)
	if req.Status != 2 {
		q = q.Where(p.Status.Eq(req.Status))
	}
	// 平台信息
	if len(req.Platform) > 0 {
		q = q.Where(p.Platform.Like("%" + req.Platform + "%"))
	}
	// 浏览器类型
	if len(req.Browser) > 0 {
		q = q.Where(p.Browser.Like("%" + req.Browser + "%"))
	}
	// 浏览器版本
	if len(req.Version) > 0 {
		q = q.Where(p.Version.Like("%" + req.Version + "%"))
	}
	// 操作系统
	if len(req.Os) > 0 {
		q = q.Where(p.Os.Like("%" + req.Os + "%"))
	}

	result, count, err := q.FindByPage(int((req.PageNum-1)*req.PageSize), int(req.PageSize))

	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	var list []*operatelog.QueryOperateLogListResp

	for _, item := range result {
		list = append(list, &operatelog.QueryOperateLogListResp{
			Id:              item.ID,                           // 日志主键
			Title:           item.Title,                        // 模块标题
			BusinessType:    item.BusinessType,                 // 业务类型（0其它 1新增 2修改 3删除）
			Method:          item.Method,                       // 方法名称
			RequestMethod:   item.RequestMethod,                // 请求方式
			OperatorType:    item.OperatorType,                 // 操作类别（0其它 1后台用户 2手机端用户）
			OperateName:     item.OperateName,                  // 操作人员
			DeptName:        item.DeptName,                     // 部门名称
			OperateUrl:      item.OperateURL,                   // 请求URL
			OperateIp:       item.OperateIP,                    // 主机地址
			OperateLocation: item.OperateLocation,              // 操作地点
			OperateParam:    item.OperateParam,                 // 请求参数
			JsonResult:      item.JSONResult,                   // 返回参数
			Platform:        item.Platform,                     // 平台信息
			Browser:         item.Browser,                      // 浏览器类型
			Version:         item.Version,                      // 浏览器版本
			Os:              item.Os,                           // 操作系统
			Arch:            item.Arch,                         // 体系结构信息
			Engine:          item.Engine,                       // 渲染引擎信息
			EngineDetails:   item.EngineDetails,                // 渲染引擎详细信息
			Extra:           item.Extra,                        // 其他信息（可选）
			Status:          item.Status,                       // 操作状态(0:异常,正常)
			ErrorMsg:        item.ErrorMsg,                     // 错误消息
			OperateTime:     utils.TimeToStr(item.OperateTime), // 操作时间
			CostTime:        item.CostTime,                     // 消耗时间

		})
	}

	resp.SuccessPage(c, list, count)
}
