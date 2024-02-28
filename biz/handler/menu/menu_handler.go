// Code generated by hertz generator.

package menu

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/feihua/hertz-admin/biz/model/api"
	"github.com/feihua/hertz-admin/biz/model/menu"
	"github.com/feihua/hertz-admin/gen/model"
	"github.com/feihua/hertz-admin/gen/query"
	"net/http"
)

// MenuList 查询菜单列表
// @router /menu_list [POST]
func MenuList(ctx context.Context, c *app.RequestContext) {
	resp := new(menu.MenuListResp)
	var err error
	var req menu.MenuListReq
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
		hlog.CtxErrorf(ctx, "查询菜单列表异常: %v", err)
		resp.Code = api.Code_DBErr
		resp.Msg = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	//返回列表
	var list []*menu.MenuListData

	for _, item := range sysMenus {
		var updateTime string
		if item.UpdateTime != nil {
			updateTime = item.UpdateTime.Format("2006-01-02 15:04:05")
		}
		list = append(list, &menu.MenuListData{
			Id:         item.ID,
			CreateTime: item.CreateTime.Format("2006-01-02 15:04:05"),
			UpdateTime: updateTime,
			StatusId:   item.StatusID,
			Sort:       item.Sort,
			ParentId:   item.ParentID,
			MenuName:   item.MenuName,
			MenuUrl:    *item.MenuURL,
			ApiUrl:     *item.APIURL,
			MenuIcon:   *item.MenuIcon,
			Remark:     *item.Remark,
			MenuType:   item.MenuType,
		})
	}

	resp.Code = api.Code_Success
	resp.Msg = "查询菜单列表成功"
	resp.Data = list

	c.JSON(http.StatusOK, resp)
}

// MenuSave 添加菜单
// @router /menu_save [POST]
func MenuSave(ctx context.Context, c *app.RequestContext) {
	resp := new(menu.MenuSaveResp)
	var err error
	var req menu.MenuSaveReq
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.Code = api.Code_ParamInvalid
		resp.Msg = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	//添加菜单
	err = query.SysMenu.WithContext(ctx).Create(&model.SysMenu{
		MenuName: req.MenuName,
		MenuType: req.MenuType,
		StatusID: req.StatusId,
		Sort:     req.Sort,
		ParentID: req.ParentId,
		MenuURL:  &req.MenuUrl,
		APIURL:   &req.ApiUrl,
		MenuIcon: &req.MenuIcon,
		Remark:   &req.Remark,
	})

	if err != nil {
		hlog.CtxErrorf(ctx, "添加菜单异常: %v", err)
		resp.Code = api.Code_DBErr
		resp.Msg = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	resp.Code = api.Code_Success
	resp.Msg = "添加菜单成功"

	c.JSON(http.StatusOK, resp)
}

// MenuUpdate 修改菜单
// @router /menu_update [POST]
func MenuUpdate(ctx context.Context, c *app.RequestContext) {
	resp := new(menu.MenuUpdateResp)
	var err error
	var req menu.MenuUpdateReq
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.Code = api.Code_ParamInvalid
		resp.Msg = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	//判断菜单是否存在
	sysMenu := query.SysMenu
	_, err = sysMenu.WithContext(ctx).Where(sysMenu.ID.Eq(req.Id)).First()
	if err != nil {
		hlog.CtxErrorf(ctx, "菜单不存在: %v", err)
		resp.Msg = "菜单不存在"
		resp.Code = api.Code_OtherErr

		c.JSON(http.StatusOK, resp)
		return
	}

	//更新菜单
	_, err = sysMenu.WithContext(ctx).Where(sysMenu.ID.Eq(req.Id)).Updates(model.SysMenu{
		MenuName: req.MenuName,
		MenuType: req.MenuType,
		StatusID: req.StatusId,
		Sort:     req.Sort,
		ParentID: req.ParentId,
		MenuURL:  &req.MenuUrl,
		APIURL:   &req.ApiUrl,
		MenuIcon: &req.MenuIcon,
		Remark:   &req.Remark,
	})

	if err != nil {
		hlog.CtxErrorf(ctx, "修改菜单异常: %v", err)
		resp.Code = api.Code_DBErr
		resp.Msg = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	resp.Msg = "修改菜单成功"
	resp.Code = api.Code_Success

	c.JSON(http.StatusOK, resp)
}

// MenuDelete 删除菜单
// @router /menu_delete [POST]
func MenuDelete(ctx context.Context, c *app.RequestContext) {
	resp := new(menu.MenuDeleteResp)
	var err error
	var req menu.MenuDeleteReq
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.Msg = err.Error()
		resp.Code = api.Code_ParamInvalid
		c.JSON(http.StatusOK, resp)
		return
	}

	sysMenu := query.SysMenu
	//判断是否有子菜单,有则不能删除
	count, err := sysMenu.WithContext(ctx).Where(sysMenu.ParentID.In(req.Id)).Count()
	if count > 0 {
		hlog.CtxErrorf(ctx, "有子菜单,不能删除: %d", req.Id)
		resp.Msg = "有子菜单,不能删除"
		resp.Code = api.Code_OtherErr
		c.JSON(http.StatusOK, resp)
		return
	}

	//删除菜单
	_, err = sysMenu.WithContext(ctx).Where(sysMenu.ID.Eq(req.Id)).Delete()
	if err != nil {
		hlog.CtxErrorf(ctx, "删除菜单异常: %v", err)
		resp.Msg = err.Error()
		resp.Code = api.Code_DBErr
		c.JSON(http.StatusOK, resp)
		return
	}

	resp.Msg = "删除菜单成功"
	resp.Code = api.Code_Success

	c.JSON(http.StatusOK, resp)
}
