package system

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	loginLog "github.com/feihua/hertz-admin/biz/model/system/loginLog"
	loginlog "github.com/feihua/hertz-admin/biz/model/system/loginlog"
	"github.com/feihua/hertz-admin/biz/pkg/utils"
	"github.com/feihua/hertz-admin/gen/query"
	"gorm.io/gorm"

	"github.com/cloudwego/hertz/pkg/app"
)

// DeleteLoginLog 删除系统访问记录
// @router /api/system/loginLog/deleteLoginLog [POST]
func DeleteLoginLog(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var req loginLog.DeleteLoginLogReq
	err := c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, resp.Error(err))
		return
	}

	q := query.SysLoginLog

	_, err = q.WithContext(ctx).Where(q.ID.In(req.Ids...)).Delete()
	if err != nil {
		c.JSON(consts.StatusOK, resp.Error(err))
		return
	}

	c.JSON(consts.StatusOK, resp.Success("删除系统访问记录成功"))
}

// QueryLoginLogDetail 查询系统访问记录详情
// @router /api/system/loginLog/queryLoginLogDetail [POST]
func QueryLoginLogDetail(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var req loginLog.QueryLoginLogDetailReq
	err := c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, resp.Error(err))
		return
	}

	item, err := query.SysLoginLog.WithContext(ctx).Where(query.SysLoginLog.ID.Eq(req.Id)).First()
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		c.JSON(consts.StatusOK, resp.ErrorMsg("系统访问记录不存在"))
		return
	case err != nil:
		c.JSON(consts.StatusOK, resp.ErrorMsg("查询系统访问记录异常"))
		return
	}

	data := &loginLog.QueryLoginLogDetailResp{
		Id:            item.ID,                         // 访问ID
		LoginName:     item.LoginName,                  // 登录账号
		Ipaddr:        item.Ipaddr,                     // 登录IP地址
		LoginLocation: item.LoginLocation,              // 登录地点
		Platform:      item.Platform,                   // 平台信息
		Browser:       item.Browser,                    // 浏览器类型
		Version:       item.Version,                    // 浏览器版本
		Os:            item.Os,                         // 操作系统
		Arch:          item.Arch,                       // 体系结构信息
		Engine:        item.Engine,                     // 渲染引擎信息
		EngineDetails: item.EngineDetails,              // 渲染引擎详细信息
		Extra:         item.Extra,                      // 其他信息（可选）
		Status:        item.Status,                     // 登录状态(0:失败,1:成功)
		Msg:           item.Msg,                        // 提示消息
		LoginTime:     utils.TimeToStr(item.LoginTime), // 访问时间

	}

	c.JSON(consts.StatusOK, resp.Success(data))
}

// QueryLoginLogList 查询系统访问记录列表
// @router /api/system/loginLog/queryLoginLogList [POST]
func QueryLoginLogList(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var err error
	var req loginLog.QueryLoginLogListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, resp.Error(err))
		return
	}

	g := query.SysLoginLog
	q := g.WithContext(ctx)
	// 登录账号
	if len(req.LoginName) > 0 {
		q = q.Where(g.LoginName.Like("%" + req.LoginName + "%"))
	}
	// 登录IP地址
	if len(req.Ipaddr) > 0 {
		q = q.Where(g.Ipaddr.Like("%" + req.Ipaddr + "%"))
	}
	// 登录地点
	if len(req.LoginLocation) > 0 {
		q = q.Where(g.LoginLocation.Like("%" + req.LoginLocation + "%"))
	}
	// 平台信息
	if len(req.Platform) > 0 {
		q = q.Where(g.Platform.Like("%" + req.Platform + "%"))
	}
	// 浏览器类型
	if len(req.Browser) > 0 {
		q = q.Where(g.Browser.Like("%" + req.Browser + "%"))
	}
	// 浏览器版本
	if len(req.Version) > 0 {
		q = q.Where(g.Version.Like("%" + req.Version + "%"))
	}
	// 操作系统
	if len(req.Os) > 0 {
		q = q.Where(g.Os.Like("%" + req.Os + "%"))
	}
	// 状态(1:正常，0:禁用)
	if req.Status != 2 {
		q = q.Where(g.Status.Eq(req.Status))
	}
	result, count, err := q.FindByPage(int((req.PageNum-1)*req.PageSize), int(req.PageSize))

	if err != nil {
		c.JSON(consts.StatusOK, resp.Error(err))
		return
	}

	var list []*loginLog.QueryLoginLogListResp

	for _, item := range result {
		list = append(list, &loginLog.QueryLoginLogListResp{
			Id:            item.ID,                                      // 访问ID
			LoginName:     item.LoginName,                               // 登录账号
			Ipaddr:        item.Ipaddr,                                  // 登录IP地址
			LoginLocation: item.LoginLocation,                           // 登录地点
			Platform:      item.Platform,                                // 平台信息
			Browser:       item.Browser,                                 // 浏览器类型
			Version:       item.Version,                                 // 浏览器版本
			Os:            item.Os,                                      // 操作系统
			Arch:          item.Arch,                                    // 体系结构信息
			Engine:        item.Engine,                                  // 渲染引擎信息
			EngineDetails: item.EngineDetails,                           // 渲染引擎详细信息
			Extra:         item.Extra,                                   // 其他信息（可选）
			Status:        item.Status,                                  // 登录状态(0:失败,1:成功)
			Msg:           item.Msg,                                     // 提示消息
			LoginTime:     item.LoginTime.Format("2006-01-02 15:04:05"), // 访问时间

		})
	}

	c.JSON(consts.StatusOK, resp.SuccessPage(list, count))
}
