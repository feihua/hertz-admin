package system

import (
	"context"
	"errors"
	"github.com/feihua/hertz-admin/biz/dal"
	"github.com/feihua/hertz-admin/biz/model/system/menu"
	"github.com/feihua/hertz-admin/biz/pkg/mw"
	"github.com/feihua/hertz-admin/biz/pkg/utils"
	"github.com/feihua/hertz-admin/gen/model"
	"github.com/feihua/hertz-admin/gen/query"
	"gorm.io/gorm"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
)

// AddMenu 添加菜单信息
// @router /api/system/menu/addMenu [POST]
func AddMenu(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var req menu.AddMenuReq
	err := c.BindAndValidate(&req)
	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	name := req.MenuName
	path := req.MenuUrl

	m := query.SysMenu
	q := m.WithContext(ctx)

	// 1.查询菜单名称是否已存在,如果菜单已存在,则直接返回
	count, err := q.Where(m.MenuName.Eq(name)).Count()

	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	if count > 0 {
		resp.Error(c, "菜单名称已存在")
		return
	}

	// 2.查询菜单路由是否已存在,如果菜单已存在,则直接返回
	if len(path) != 0 {
		count, err = q.Where(m.MenuURL.Eq(path)).Count()

		if err != nil {
			resp.Error(c, err.Error())
			return
		}

		if count > 0 {
			resp.Error(c, "菜单路由已存在")
			return
		}
	}

	user, _ := c.Get(mw.IdentityKey)
	item := &model.SysMenu{
		MenuName: req.MenuName,         // 菜单名称
		MenuType: req.MenuType,         // 菜单类型(1：目录   2：菜单   3：按钮)
		Visible:  req.Visible,          // 显示状态（0:隐藏, 显示:1）
		Status:   req.Status,           // 菜单状态(1:正常，0:禁用)
		Sort:     req.Sort,             // 排序
		ParentID: req.ParentId,         // 父ID
		MenuURL:  req.MenuUrl,          // 路由路径
		APIURL:   req.ApiUrl,           // 接口URL
		MenuIcon: req.MenuIcon,         // 菜单图标
		Remark:   req.Remark,           // 备注
		CreateBy: user.(*mw.User).Name, // 创建者

	}

	err = q.WithContext(ctx).Create(item)

	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	resp.Success(c, "添加菜单信息成功")
}

// DeleteMenu 删除菜单信息
// @router /api/system/menu/deleteMenu [POST]
func DeleteMenu(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var req menu.DeleteMenuReq
	err := c.BindAndValidate(&req)
	if err != nil {
		resp.Error(c, err.Error())

		return
	}

	q := query.SysMenu

	count, err := q.WithContext(ctx).Where(q.ID.Eq(req.Id)).Count()

	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	if count == 0 {
		resp.Error(c, "菜单不存在")
		return
	}

	// 1.查询菜单是否有子菜单
	count, err = q.WithContext(ctx).Where(q.ParentID.Eq(req.Id)).Count()
	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	if count > 0 {
		resp.Error(c, "存在子菜单,不允许删除")
		return
	}

	// 2.查询菜单是否菜单已分配
	count, err = query.SysRoleMenu.WithContext(ctx).Where(query.SysRoleMenu.MenuID.Eq(req.Id)).Count()
	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	if count > 0 {
		resp.Error(c, "菜单已分配,不允许删除")
		return
	}

	_, err = q.WithContext(ctx).Where(q.ID.Eq(req.Id)).Delete()
	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	resp.Success(c, "删除菜单信息成功")
}

// UpdateMenu 更新菜单信息
// @router /api/system/menu/updateMenu [POST]
func UpdateMenu(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var req menu.UpdateMenuReq
	err := c.BindAndValidate(&req)
	if err != nil {
		resp.Error(c, err.Error())

		return
	}

	name := req.MenuName
	path := req.MenuUrl

	sMenu := query.SysMenu
	q := sMenu.WithContext(ctx)

	// 1.根据菜单信息id查询菜单信息是否已存在
	item, err := q.Where(query.SysMenu.ID.Eq(req.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		resp.Error(c, "菜单不存在")
		return
	case err != nil:
		resp.Error(c, err.Error())
		return
	}

	// 2.查询菜单名称是否已存在,如果菜单已存在,则直接返回
	count, err := q.Where(sMenu.ID.Neq(req.Id), sMenu.MenuName.Eq(name)).Count()
	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	if count > 0 {
		resp.Error(c, "菜单名称已存在")
		return
	}

	// 3.查询菜单路由是否已存在,如果菜单已存在,则直接返回
	if len(path) != 0 {
		count, err = q.Where(sMenu.ID.Neq(req.Id), sMenu.MenuURL.Eq(path)).Count()

		if err != nil {
			resp.Error(c, err.Error())
			return
		}

		if count > 0 {
			resp.Error(c, "菜单路由已存在")
			return
		}
	}

	user, _ := c.Get(mw.IdentityKey)
	updateTime := time.Now()
	sysMenu := &model.SysMenu{
		ID:         req.Id,               // 主键
		MenuName:   req.MenuName,         // 菜单名称
		MenuType:   req.MenuType,         // 菜单类型(1：目录   2：菜单   3：按钮)
		Visible:    req.Visible,          // 显示状态（0:隐藏, 显示:1）
		Status:     req.Status,           // 菜单状态(1:正常，0:禁用)
		Sort:       req.Sort,             // 排序
		ParentID:   req.ParentId,         // 父ID
		MenuURL:    req.MenuUrl,          // 路由路径
		APIURL:     req.ApiUrl,           // 接口URL
		MenuIcon:   req.MenuIcon,         // 菜单图标
		Remark:     req.Remark,           // 备注
		CreateBy:   item.CreateBy,        // 创建者
		CreateTime: item.CreateTime,      // 创建时间
		UpdateBy:   user.(*mw.User).Name, // 更新者
		UpdateTime: &updateTime,          // 更新时间

	}

	// 2.菜单信息存在时,则直接更新菜单信息
	// 4.菜单存在时,则直接更新菜单
	err = dal.DB.Model(&model.SysMenu{}).WithContext(ctx).Where(query.SysMenu.ID.Eq(req.Id)).Save(sysMenu).Error

	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	resp.Success(c, "更新菜单信息成功")
}

// UpdateMenuStatus 菜单信息状态
// @router /api/system/menu/updateMenuStatus [POST]
func UpdateMenuStatus(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var req menu.UpdateMenuStatusReq
	err := c.BindAndValidate(&req)
	if err != nil {
		resp.Error(c, err.Error())

		return
	}

	q := query.SysMenu
	_, err = q.WithContext(ctx).Where(q.ID.Eq(req.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		resp.Error(c, "菜单不存在")
		return
	case err != nil:
		resp.Error(c, err.Error())
		return
	}

	_, err = q.WithContext(ctx).Where(q.ID.Eq(req.Id)).Update(q.Status, req.Status)

	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	resp.Success(c, "更新菜单信息状态成功")
}

// QueryMenuDetail 查询菜单信息详情
// @router /api/system/menu/queryMenuDetail [POST]
func QueryMenuDetail(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var req menu.QueryMenuDetailReq
	err := c.BindAndValidate(&req)
	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	item, err := query.SysMenu.WithContext(ctx).Where(query.SysMenu.ID.Eq(req.Id)).First()
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		resp.Error(c, "菜单信息不存在")
		return
	case err != nil:
		resp.Error(c, err.Error())
		return
	}

	data := &menu.QueryMenuDetailResp{
		Id:         item.ID,                             // 主键
		MenuName:   item.MenuName,                       // 菜单名称
		MenuType:   item.MenuType,                       // 菜单类型(1：目录   2：菜单   3：按钮)
		Visible:    item.Visible,                        // 显示状态（0:隐藏, 显示:1）
		Status:     item.Status,                         // 菜单状态(1:正常，0:禁用)
		Sort:       item.Sort,                           // 排序
		ParentId:   item.ParentID,                       // 父ID
		MenuUrl:    item.MenuURL,                        // 路由路径
		ApiUrl:     item.APIURL,                         // 接口URL
		MenuIcon:   item.MenuIcon,                       // 菜单图标
		Remark:     item.Remark,                         // 备注
		CreateBy:   item.CreateBy,                       // 创建者
		CreateTime: utils.TimeToStr(item.CreateTime),    // 创建时间
		UpdateBy:   item.UpdateBy,                       // 更新者
		UpdateTime: utils.TimeToString(item.UpdateTime), // 更新时间

	}
	resp.Success(c, data)
}

// QueryMenuList 查询菜单信息列表
// @router /api/system/menu/queryMenuList [POST]
func QueryMenuList(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var err error
	var req menu.QueryMenuListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	m := query.SysMenu
	q := m.WithContext(ctx)
	// 菜单名称
	if len(req.MenuName) > 0 {
		q = q.Where(m.MenuName.Like("%" + req.MenuName + "%"))
	}

	menus, err := q.Find()

	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	var list []*menu.QueryMenuListResp

	for _, item := range menus {
		list = append(list, &menu.QueryMenuListResp{
			Id:         item.ID,                             // 主键
			MenuName:   item.MenuName,                       // 菜单名称
			MenuType:   item.MenuType,                       // 菜单类型(1：目录   2：菜单   3：按钮)
			Visible:    item.Visible,                        // 显示状态（0:隐藏, 显示:1）
			Status:     item.Status,                         // 菜单状态(1:正常，0:禁用)
			Sort:       item.Sort,                           // 排序
			ParentId:   item.ParentID,                       // 父ID
			MenuUrl:    item.MenuURL,                        // 路由路径
			ApiUrl:     item.APIURL,                         // 接口URL
			MenuIcon:   item.MenuIcon,                       // 菜单图标
			Remark:     item.Remark,                         // 备注
			CreateBy:   item.CreateBy,                       // 创建者
			CreateTime: utils.TimeToStr(item.CreateTime),    // 创建时间
			UpdateBy:   item.UpdateBy,                       // 更新者
			UpdateTime: utils.TimeToString(item.UpdateTime), // 更新时间

		})
	}

	resp.Success(c, list)
}
