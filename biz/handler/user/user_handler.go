// Code generated by hertz generator.

package user

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/feihua/hertz-admin/biz/dal"
	"github.com/feihua/hertz-admin/biz/handler/role"
	"github.com/feihua/hertz-admin/biz/model/api"
	"github.com/feihua/hertz-admin/biz/model/user"
	"github.com/feihua/hertz-admin/gen/model"
	"github.com/feihua/hertz-admin/gen/query"
	"net/http"
)

// UserList 查询用户列表
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

	q := query.SysUser.WithContext(ctx)
	if len(req.UserName) != 0 {
		q = q.Where(query.SysUser.UserName.Like("%" + req.UserName + "%"))
	}

	if len(req.Mobile) != 0 {
		q = q.Where(query.SysUser.Mobile.Like("%" + req.Mobile + "%"))
	}

	result, count, err := q.FindByPage(int((req.PageNo-1)*req.PageSize), int(req.PageSize))

	if err != nil {
		hlog.CtxErrorf(ctx, "查询用户列表异常: %v", err)
		resp.Code = api.Code_DBErr
		resp.Msg = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	var list []*user.UserListData

	for _, item := range result {
		var updateTime string
		if item.UpdateTime != nil {
			updateTime = item.UpdateTime.Format("2006-01-02 15:04:05")
		}
		list = append(list, &user.UserListData{
			Id:         item.ID,
			StatusId:   item.StatusID,
			Sort:       item.Sort,
			Mobile:     item.Mobile,
			UserName:   item.UserName,
			Remark:     *item.Remark,
			CreateTime: item.CreateTime.Format("2006-01-02 15:04:05"),
			UpdateTime: updateTime,
		})
	}

	resp.Code = api.Code_Success
	resp.Msg = "查询用户列表成功"
	resp.Data = list
	resp.Total = count

	c.JSON(http.StatusOK, resp)
}

// UserSave 添加用户
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

	//判断用户名称或者手机号是否已存在
	u := query.SysUser
	count, err := u.WithContext(ctx).Where(u.Mobile.Eq(req.Mobile)).Or(u.UserName.Eq(req.UserName)).Count()
	if err != nil {
		hlog.CtxErrorf(ctx, "根据用户名称或者手机号查询异常: %v", err)
		resp.Code = api.Code_DBErr
		resp.Msg = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	if count > 0 {
		hlog.CtxErrorf(ctx, "用户名称或者手机号已存在: %v", req.Mobile)
		resp.Code = api.Code_OtherErr
		resp.Msg = "用户名称或者手机号已存在"
		c.JSON(http.StatusOK, resp)
		return
	}

	//默认的密码为123456
	var defaultPassword = "123456"

	//添加用户
	err = u.WithContext(ctx).Create(&model.SysUser{
		Mobile:   req.Mobile,
		UserName: req.UserName,
		Password: defaultPassword,
		StatusID: req.StatusID,
		Sort:     req.Sort,
		Remark:   &req.Remark,
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

	c.JSON(http.StatusOK, resp)
}

// UserUpdate 更新用户信息
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

	//判断是否为管理员
	if req.Id == 1 {
		hlog.CtxErrorf(ctx, "超级管理员,不能修改: %v", req.Id)
		resp.Msg = "超级管理员,不能修改"
		resp.Code = api.Code_OtherErr
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
		Mobile:   req.Mobile,
		UserName: req.UserName,
		StatusID: req.StatusId,
		Sort:     req.Sort,
		Remark:   &req.Remark,
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

	if role.Contains(req.Ids, 1) {
		hlog.CtxErrorf(ctx, "超级管理员,不能删除: %v", req.Ids)
		resp.Msg = "超级管理员,不能删除"
		resp.Code = api.Code_OtherErr
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

	c.JSON(http.StatusOK, resp)
}

// QueryUserMenu .
// @router /query_user_menu [GET]
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

	sysMenuList := make([]*model.SysMenu, 0)
	if req.UserId == 1 {
		//查询所有菜单
		sysMenuList, err = query.SysMenu.WithContext(ctx).Find()
	} else {
		sql := `select u.*
		from sys_user_role t
				 left join sys_role usr on t.role_id = usr.id
				 left join sys_role_menu srm on usr.id = srm.role_id
				 left join sys_menu u on srm.menu_id = u.id
		where t.user_id = ?`
		dal.DB.WithContext(ctx).Raw(sql, req.UserId).Scan(&sysMenuList)
	}

	if err != nil {
		resp.Code = api.Code_DBErr
		resp.Msg = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	menuList := make([]*user.UserMenuList, 0)
	var path []string

	for _, item := range sysMenuList {
		if item.MenuType != 3 {
			menuList = append(menuList, &user.UserMenuList{
				Id:       int32(item.ID),
				ParentId: int32(item.ParentID),
				Name:     item.MenuName,
				Path:     *item.MenuURL,
				ApiUrl:   *item.APIURL,
				MenuType: item.MenuType,
				Icon:     *item.MenuIcon,
			})
		}

		if *item.APIURL != "" {
			path = append(path, *item.APIURL)
		}

	}

	data := &user.QueryUserMenuData{
		SysMenu: menuList,
		BtnMenu: path,
		Avatar:  "https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png",
		Name:    req.UserName,
	}

	resp.Code = api.Code_Success
	resp.Msg = "查询用户菜单成功"
	resp.Data = data

	c.JSON(http.StatusOK, resp)
}

// QueryUserRole .
// @router /api/query_user_role [POST]
func QueryUserRole(ctx context.Context, c *app.RequestContext) {
	resp := new(user.QueryUserRoleResp)
	var err error
	var req user.QueryUserRoleReq
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.Code = api.Code_ParamInvalid
		resp.Msg = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	//查询所有角色
	roles, err := query.SysRole.WithContext(ctx).Find()
	if err != nil {
		resp.Code = api.Code_DBErr
		resp.Msg = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	//默认查询的是所有的
	var userRoleIds []int64
	roleList := make([]*user.UserRoleList, 0)

	for _, item := range roles {
		userRoleIds = append(userRoleIds, item.ID)
		var updateTime string
		if item.UpdateTime != nil {
			updateTime = item.UpdateTime.Format("http.StatusOK6-01-02 15:04:05")
		}
		roleList = append(roleList, &user.UserRoleList{
			Id:         item.ID,
			CreateTime: item.CreateTime.Format("http.StatusOK6-01-02 15:04:05"),
			UpdateTime: updateTime,
			StatusId:   item.StatusID,
			Sort:       item.Sort,
			RoleName:   item.RoleName,
			Remark:     item.Remark,
		})
	}

	if req.UserId != 1 {
		rm := query.SysUserRole

		sysRoles, _ := rm.WithContext(ctx).Select(rm.RoleID).Where(rm.UserID.Eq(req.UserId)).Find()

		var roleIds []int64
		for _, m := range sysRoles {
			roleIds = append(roleIds, m.RoleID)
		}

		//重新赋值
		userRoleIds = roleIds
	}

	data := &user.QueryUserRoleData{
		SysRoleList: roleList,
		UserRoleIds: userRoleIds,
	}

	resp.Code = api.Code_Success
	resp.Msg = "查询用户角色成功"
	resp.Data = data

	c.JSON(http.StatusOK, resp)
}

// UpdateUserRole .
// @router /api/update_user_role [POST]
func UpdateUserRole(ctx context.Context, c *app.RequestContext) {
	resp := new(user.UpdateUserRoleResp)
	var err error
	var req user.UpdateUserRoleReq
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.Code = api.Code_ParamInvalid
		resp.Msg = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	if req.UserId == 1 {
		resp.Code = api.Code_OtherErr
		resp.Msg = "不能修改超级管理员的角色"
		c.JSON(http.StatusOK, resp)
		return
	}

	userRole := query.SysUserRole
	//根据用户id删除关联
	_, err = userRole.WithContext(ctx).Where(userRole.UserID.Eq(req.UserId)).Delete()
	if err != nil {
		resp.Code = api.Code_DBErr
		resp.Msg = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	//组装批量数据
	userRoles := make([]*model.SysUserRole, 0)
	for _, roleId := range req.RoleIds {
		userRoles = append(userRoles, &model.SysUserRole{
			UserID: req.UserId,
			RoleID: roleId,
		})
	}

	//批量添加
	err = userRole.WithContext(ctx).CreateInBatches(userRoles, len(req.RoleIds))
	if err != nil {
		resp.Code = api.Code_DBErr
		resp.Msg = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	resp.Code = api.Code_Success
	resp.Msg = "分配角色成功"

	c.JSON(http.StatusOK, resp)
}
