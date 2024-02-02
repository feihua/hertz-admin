// Code generated by hertz generator.

package user

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"hertz_admin/biz/model/api"
	"hertz_admin/gen/model"
	"hertz_admin/gen/query"
	"net/http"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"hertz_admin/biz/model/user"
)

// QueryUserMenu .
// @router /query_user_menu [POST]
func QueryUserMenu(ctx context.Context, c *app.RequestContext) {
	resp := new(user.QueryUserMenuResp)
	var err error
	var req user.QueryUserMenuReq
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.Code = api.Code_ParamInvalid
		resp.Msg = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UserList .
// @router /user_list [POST]
func UserList(ctx context.Context, c *app.RequestContext) {
	resp := new(user.UserListResp)
	var err error
	var req user.UserListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.Code = api.Code_ParamInvalid
		resp.Msg = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	result, count, err := query.SysUser.WithContext(ctx).FindByPage(int(req.PageNo), int(req.PageSize))
	//users, total, err := entity.QuerySysUser(nil, req.PageNo, req.PageSize)

	if err != nil {
		hlog.CtxErrorf(ctx, "查询用户列表异常: %v", err)
		resp.Code = api.Code_DBErr
		resp.Msg = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	var list []*user.UserListData

	for _, item := range result {
		list = append(list, &user.UserListData{
			Id:         item.ID,
			StatusId:   item.StatusID,
			Sort:       item.Sort,
			Mobile:     item.Mobile,
			UserName:   item.UserName,
			Remark:     *item.Remark,
			CreateTime: item.CreateTime.Format("http.StatusOK6-01-02 15:04:05"),
			UpdateTime: item.UpdateTime.Format("http.StatusOK6-01-02 15:04:05"),
		})
	}

	resp.Code = api.Code_Success
	resp.Msg = "查询用户列表成功"
	resp.Data = list
	resp.Total = count

	hlog.CtxDebugf(ctx, "查询用户列表成功: %v", resp)
	c.JSON(http.StatusOK, resp)
}

// UserSave .
// @router /user_save [POST]
func UserSave(ctx context.Context, c *app.RequestContext) {
	resp := new(user.UserSaveResp)
	var err error
	var req user.UserSaveReq
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.Code = api.Code_ParamInvalid
		resp.Msg = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	err = query.SysUser.WithContext(ctx).Create(&model.SysUser{
		Mobile:     req.Mobile,
		UserName:   req.UserName,
		Password:   nil,
		StatusID:   req.StatusID,
		Sort:       req.Sort,
		Remark:     &req.Remark,
		CreateTime: time.Now(),
	})

	if err != nil {
		hlog.CtxErrorf(ctx, "添加用户异常: %v", err)
		resp.Code = api.Code_DBErr
		resp.Msg = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	resp.Code = api.Code_Success
	resp.Msg = "添加用户成功"

	hlog.CtxDebugf(ctx, "添加用户成功: %v", resp)
	c.JSON(http.StatusOK, resp)
}

// UserUpdate .
// @router /user_update [POST]
func UserUpdate(ctx context.Context, c *app.RequestContext) {
	resp := new(user.UserUpdateResp)
	var err error
	var req user.UserUpdateReq
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.Code = api.Code_ParamInvalid
		resp.Msg = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	sysUser := query.SysUser
	_, err = sysUser.WithContext(ctx).Where(sysUser.ID.Eq(req.Id)).First()
	if err != nil {
		hlog.CtxErrorf(ctx, "用户不存在: %v", err)
		resp.Msg = "用户不存在"
		resp.Code = api.Code_OtherErr

		c.JSON(http.StatusOK, resp)
		return
	}

	_, err = sysUser.WithContext(ctx).Where(sysUser.ID.Eq(req.Id)).Updates(model.SysUser{
		Mobile:     req.Mobile,
		UserName:   req.UserName,
		StatusID:   req.StatusId,
		Sort:       req.Sort,
		Remark:     &req.Remark,
		UpdateTime: time.Now(),
	})

	if err != nil {
		hlog.CtxErrorf(ctx, "修改用户异常: %v", err)
		resp.Code = api.Code_DBErr
		resp.Msg = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	resp.Msg = "修改用户成功"
	resp.Code = api.Code_Success

	hlog.CtxDebugf(ctx, "修改用户成功: %v", resp)
	c.JSON(http.StatusOK, resp)
}

// UserDelete 删除用户
// @router /user_delete [POST]
func UserDelete(ctx context.Context, c *app.RequestContext) {
	resp := new(user.UserDeleteResp)
	var err error
	var req user.UserDeleteReq
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.Code = api.Code_ParamInvalid
		resp.Msg = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	sysUser := query.SysUser
	_, err = sysUser.WithContext(ctx).Where(sysUser.ID.In(req.Ids...)).Delete()
	if err != nil {
		hlog.CtxErrorf(ctx, "删除用户异常: %v", err)
		resp.Msg = err.Error()
		resp.Code = api.Code_DBErr
		c.JSON(http.StatusOK, resp)
		return
	}

	resp.Code = api.Code_Success
	resp.Msg = "删除用户成功"

	hlog.CtxDebugf(ctx, "删除用户成功: %v", resp)
	c.JSON(http.StatusOK, resp)
}