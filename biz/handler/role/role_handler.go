// Code generated by hertz generator.

package role

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"hertz_admin/biz/dal"
	"hertz_admin/biz/model/api"
	"hertz_admin/biz/model/menu"
	"hertz_admin/gen/model"
	"hertz_admin/gen/query"
	"net/http"
	"strconv"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"hertz_admin/biz/model/role"
)

// RoleList 查询角色列表
// @router /role_list [POST]
func RoleList(ctx context.Context, c *app.RequestContext) {
	resp := new(role.RoleListResp)
	var err error
	var req role.RoleListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.Code = api.Code_ParamInvalid
		resp.Msg = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	//查询角色
	//result, count, err := query.SysRole.WithContext(ctx).FindByPage(int(req.PageNo), int(req.PageSize))
	result, count, err := dal.QueryRoleList(req.RoleName, req.PageNo, req.PageSize)

	if err != nil {
		hlog.CtxErrorf(ctx, "查询角色列表异常: %v", err)
		resp.Code = api.Code_DBErr
		resp.Msg = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	//返回数据
	var list []*role.RoleListData

	for _, item := range result {
		var updateTime string
		if item.UpdateTime != nil {
			updateTime = item.UpdateTime.Format("http.StatusOK6-01-02 15:04:05")
		}
		list = append(list, &role.RoleListData{
			Id:         item.ID,
			CreateTime: item.CreateTime.Format("http.StatusOK6-01-02 15:04:05"),
			UpdateTime: updateTime,
			StatusId:   item.StatusID,
			Sort:       item.Sort,
			RoleName:   item.RoleName,
			Remark:     item.Remark,
		})
	}

	resp.Code = api.Code_Success
	resp.Msg = "查询角色列表成功"
	resp.Data = list
	resp.Total = count

	hlog.CtxDebugf(ctx, "查询角色列表成功: %v", resp)
	c.JSON(http.StatusOK, resp)
}

// RoleSave 添加角色
// @router /role_save [POST]
func RoleSave(ctx context.Context, c *app.RequestContext) {
	resp := new(role.RoleSaveResp)
	var err error
	var req role.RoleSaveReq
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.Code = api.Code_ParamInvalid
		resp.Msg = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	//判断角色名称是否已存在
	count, err := query.SysRole.WithContext(ctx).Where(query.SysRole.RoleName.Eq(req.RoleName)).Count()
	if err != nil {
		hlog.CtxErrorf(ctx, "根据角色名称查询异常: %v", err)
		resp.Code = api.Code_DBErr
		resp.Msg = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	if count > 0 {
		hlog.CtxErrorf(ctx, "角色名称已存在: %v", req.RoleName)
		resp.Code = api.Code_OtherErr
		resp.Msg = "角色名称已存在"
		c.JSON(http.StatusOK, resp)
		return
	}

	//添加角色
	err = query.SysRole.WithContext(ctx).Create(&model.SysRole{
		RoleName:   req.RoleName,
		StatusID:   req.StatusId,
		Sort:       req.Sort,
		Remark:     req.Remark,
		CreateTime: time.Now(),
	})

	if err != nil {
		hlog.CtxErrorf(ctx, "添加角色异常: %v", err)
		resp.Code = api.Code_DBErr
		resp.Msg = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	resp.Code = api.Code_Success
	resp.Msg = "添加角色成功"

	hlog.CtxDebugf(ctx, "添加角色成功: %v", resp)
	c.JSON(http.StatusOK, resp)
}

// RoleUpdate .
// @router /role_update [POST]
func RoleUpdate(ctx context.Context, c *app.RequestContext) {
	resp := new(role.RoleUpdateResp)
	var err error
	var req role.RoleUpdateReq
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.Code = api.Code_ParamInvalid
		resp.Msg = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	//查询角色是否存在
	sysRole := query.SysRole
	_, err = sysRole.WithContext(ctx).Where(sysRole.ID.Eq(req.Id)).First()
	if err != nil {
		hlog.CtxErrorf(ctx, "角色不存在: %v", err)
		resp.Msg = "角色不存在"
		resp.Code = api.Code_OtherErr

		c.JSON(http.StatusOK, resp)
		return
	}

	//更新角色
	_, err = sysRole.WithContext(ctx).Where(sysRole.ID.Eq(req.Id)).Updates(model.SysRole{
		RoleName: req.RoleName,
		StatusID: req.StatusId,
		Sort:     req.Sort,
		Remark:   req.Remark,
	})

	if err != nil {
		hlog.CtxErrorf(ctx, "修改角色异常: %v", err)
		resp.Code = api.Code_DBErr
		resp.Msg = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	resp.Msg = "修改角色成功"
	resp.Code = api.Code_Success

	hlog.CtxDebugf(ctx, "修改角色成功: %v", resp)
	c.JSON(http.StatusOK, resp)
}

// RoleDelete 删除角色
// @router /role_delete [POST]
func RoleDelete(ctx context.Context, c *app.RequestContext) {
	resp := new(role.RoleDeleteResp)
	var err error
	var req role.RoleDeleteReq
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.Code = api.Code_ParamInvalid
		resp.Msg = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	ids := req.Ids
	//角色id为1的是超级管理员,不能删除
	if Contains(ids, 1) {
		hlog.CtxErrorf(ctx, "超级管理员,不能删除: %v", ids)
		resp.Msg = "超级管理员,不能删除"
		resp.Code = api.Code_OtherErr
		c.JSON(http.StatusOK, resp)
		return
	}
	sysRole := query.SysRole
	//删除角色
	_, err = sysRole.WithContext(ctx).Where(sysRole.ID.In(ids...)).Delete()
	if err != nil {
		hlog.CtxErrorf(ctx, "删除角色异常: %v", err)
		resp.Msg = err.Error()
		resp.Code = api.Code_DBErr
		c.JSON(http.StatusOK, resp)
		return
	}

	resp.Code = api.Code_Success
	resp.Msg = "删除角色成功"

	hlog.CtxDebugf(ctx, "删除角色成功: %v", resp)
	c.JSON(http.StatusOK, resp)
}

// Contains 判断是否包含某值
func Contains(s []int64, target int64) bool {
	for _, v := range s {
		if v == target {
			return true
		}
	}
	return false
}

// QueryRoleMenu .
// @router /query_role_menu [POST]
func QueryRoleMenu(ctx context.Context, c *app.RequestContext) {
	resp := new(role.QueryRoleMenuResp)
	var err error
	var req role.QueryRoleMenuReq
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.Code = api.Code_ParamInvalid
		resp.Msg = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	//查询所有菜单
	sysMenus, err := query.SysMenu.WithContext(ctx).Find()
	if err != nil {
		resp.Code = api.Code_DBErr
		resp.Msg = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	//默认查询的是所有的
	var roleMenu []int64
	menuList := make([]*role.MenuDataList, 0)

	for _, item := range sysMenus {
		roleMenu = append(roleMenu, item.ID)
		menuList = append(menuList, &role.MenuDataList{
			Id:       item.ID,
			ParentId: item.ParentID,
			Title:    item.MenuName,
			Key:      strconv.FormatInt(item.ID, 10),
			Label:    item.MenuName,
		})
	}

	if req.RoleId != 1 {
		rm := query.SysRoleMenu

		menus, _ := rm.WithContext(ctx).Select(rm.MenuID).Where(rm.RoleID.Eq(req.RoleId)).Find()

		var menuIds []int64
		for _, m := range menus {
			menuIds = append(menuIds, m.MenuID)
		}

		//重新赋值
		roleMenu = menuIds
	}

	data := &role.QueryRoleMenuData{
		RoleMenu: roleMenu,
		MenuList: menuList,
	}

	resp.Code = api.Code_Success
	resp.Msg = "查询角色菜单成功"
	resp.Data = data

	hlog.CtxDebugf(ctx, "查询角色菜单成功: %v", resp)
	c.JSON(http.StatusOK, resp)
}

// UpdateRoleMenu 更新角色与菜单的关联
// @router /update_role_menu [POST]
func UpdateRoleMenu(ctx context.Context, c *app.RequestContext) {
	resp := new(menu.MenuSaveResp)
	var err error
	var req role.UpdateRoleMenuReq
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.Code = api.Code_ParamInvalid
		resp.Msg = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	if req.RoleId == 1 {
		resp.Code = api.Code_OtherErr
		resp.Msg = "不能修改超级管理员的权限"
		c.JSON(http.StatusOK, resp)
		return
	}

	roleMenu := query.SysRoleMenu
	//根据角色id删除关联
	_, err = roleMenu.WithContext(ctx).Where(roleMenu.RoleID.Eq(req.RoleId)).Delete()
	if err != nil {
		resp.Code = api.Code_DBErr
		resp.Msg = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	//组装批量数据
	sysRoleMenus := make([]*model.SysRoleMenu, 0)
	for _, menuId := range req.MenuIds {
		sysRoleMenus = append(sysRoleMenus, &model.SysRoleMenu{
			RoleID: req.RoleId,
			MenuID: menuId,
		})
	}

	//批量添加
	err = roleMenu.WithContext(ctx).CreateInBatches(sysRoleMenus, len(req.MenuIds))
	if err != nil {
		resp.Code = api.Code_DBErr
		resp.Msg = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	resp.Code = api.Code_Success
	resp.Msg = "分配权限成功"

	hlog.CtxDebugf(ctx, "分配权限成功: %v", resp)
	c.JSON(http.StatusOK, resp)
}
