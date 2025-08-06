package system

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/feihua/hertz-admin/biz/dal"
	"github.com/feihua/hertz-admin/biz/handler/common"
	"github.com/feihua/hertz-admin/biz/model/system/user"
	"github.com/feihua/hertz-admin/biz/pkg/mw"
	"github.com/feihua/hertz-admin/biz/pkg/utils"
	"github.com/feihua/hertz-admin/gen/model"
	"github.com/feihua/hertz-admin/gen/query"
	"gorm.io/gorm"
	"strings"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
)

// AddUser 添加用户信息
// @router /api/system/user/addUser [POST]
func AddUser(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var req user.AddUserReq
	err := c.BindAndValidate(&req)
	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	q := query.SysUser

	// 1.查询用户名称是否存在
	name := req.UserName
	count, err := q.WithContext(ctx).Where(q.UserName.Eq(name)).Count()

	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	if count > 0 {
		resp.Error(c, "用户已存在")
		return
	}

	// 2.查询手机号是否存在
	mobile := req.Mobile
	count, err = q.WithContext(ctx).Where(q.Mobile.Eq(mobile)).Count()

	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	if count > 0 {
		resp.Error(c, "手机号已存在")
		return
	}

	// 3.查询邮箱是否存在
	email := req.Email
	count, err = q.WithContext(ctx).Where(q.Email.Eq(email)).Count()

	if err != nil {
		resp.Error(c, err.Error())

		return
	}

	if count > 0 {
		resp.Error(c, "邮箱已存在")
		return
	}

	avatar := req.Avatar
	// 默认用户头像
	if len(avatar) == 0 {
		avatar = "https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png"
	}
	user, _ := c.Get(mw.IdentityKey)
	item := &model.SysUser{
		Mobile:   req.Mobile,           // 手机号码
		UserName: req.UserName,         // 用户账号
		NickName: req.NickName,         // 用户昵称
		UserType: req.UserType,         // 用户类型（00系统用户）
		Avatar:   req.Avatar,           // 头像路径
		Email:    req.Email,            // 用户邮箱
		Password: req.Password,         // 密码
		Status:   req.Status,           // 状态(1:正常，0:禁用)
		DeptID:   req.DeptId,           // 部门ID
		Remark:   req.Remark,           // 备注
		CreateBy: user.(*mw.User).Name, // 更新者
	}

	err = query.Q.Transaction(func(tx *query.Query) error {
		// 4.用户不存在时,则直接添加用户
		err = tx.SysUser.WithContext(ctx).Create(item)

		if err != nil {
			return err
		}

		var userPosts []*model.SysUserPost
		for _, postId := range req.PostIds {
			userPosts = append(userPosts, &model.SysUserPost{
				UserID: item.ID,
				PostID: postId,
			})
		}

		postDo := tx.SysUserPost.WithContext(ctx)
		// 5.清空用户与岗位关联
		_, err = postDo.Where(tx.SysUserPost.UserID.Eq(item.ID)).Delete()
		if err != nil {
			return err
		}

		// 6.添加用户与岗位关联
		err = postDo.Create(userPosts...)
		if err != nil {
			return err
		}

		roleDo := tx.SysUserRole.WithContext(ctx)
		// 7.清空用户与角色关联(防止脏数据)
		_, err = roleDo.Where(tx.SysUserRole.UserID.Eq(item.ID)).Delete()
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	resp.Success(c, "添加用户信息成功")
}

// DeleteUser 删除用户信息
// @router /api/system/user/deleteUser [POST]
func DeleteUser(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var req user.DeleteUserReq
	err := c.BindAndValidate(&req)
	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	ids := req.Ids

	for _, userId := range ids {
		// 1.判断是不是超级管理员
		isAdmin, _ := common.IsAdmin(ctx, userId)
		if isAdmin {
			resp.Error(c, "不允许操作超级管理员用户")
			return
		}

	}

	err = query.Q.Transaction(func(tx *query.Query) error {
		// 2.删除用户
		u := tx.SysUser
		if _, err = u.WithContext(ctx).Where(u.ID.In(ids...)).Delete(); err != nil {
			return err
		}

		// 3.删除用户与角色的关联
		role := tx.SysUserRole
		if _, err = role.WithContext(ctx).Where(role.UserID.In(ids...)).Delete(); err != nil {
			return err
		}

		// 4.删除用户与岗位的关联
		post := tx.SysUserPost
		if _, err = post.WithContext(ctx).Where(post.UserID.In(ids...)).Delete(); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	resp.Success(c, "删除用户信息成功")
}

// UpdateUser 更新用户信息
// @router /api/system/user/updateUser [POST]
func UpdateUser(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var req user.UpdateUserReq
	err := c.BindAndValidate(&req)
	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	isAdmin, _ := common.IsAdmin(ctx, req.Id)
	if isAdmin {
		resp.Error(c, "不允许操作超级管理员用户")
		return
	}

	name := req.UserName // 用户名
	mobile := req.Mobile // 手机号
	email := req.Email   // 邮箱

	u := query.SysUser
	q := u.WithContext(ctx)

	// 1.根据用户id查询用户是否已存在
	item, err := q.Where(u.ID.Eq(req.Id)).First()

	// 1.判断用户是否存在
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		resp.Error(c, "用户不存在")
		return
	case err != nil:
		resp.Error(c, err.Error())
		return
	}

	// 2.查询用户名称是否存在
	count, err := q.Where(u.ID.Neq(req.Id), u.UserName.Eq(name)).Count()

	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	if count > 0 {
		resp.Error(c, "用户已存在")
		return
	}

	// 3.查询手机号是否存在
	count, err = q.Where(u.ID.Neq(req.Id), u.Mobile.Eq(mobile)).Count()

	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	if count > 0 {
		resp.Error(c, "手机号已存在")
		return
	}

	// 4.查询邮箱是否存在
	count, err = q.Where(u.ID.Neq(req.Id), u.Email.Eq(email)).Count()

	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	if count > 0 {
		resp.Error(c, "邮箱已存在")
		return
	}

	// 5.用户存在时,则直接更新用户
	user, _ := c.Get(mw.IdentityKey)
	updateTime := time.Now()
	sysUser := &model.SysUser{
		ID:            req.Id,               // 主键
		Mobile:        req.Mobile,           // 手机号码
		UserName:      req.UserName,         // 用户账号
		NickName:      req.NickName,         // 用户昵称
		UserType:      req.UserType,         // 用户类型（00系统用户）
		Avatar:        req.Avatar,           // 头像路径
		Email:         req.Email,            // 用户邮箱
		Password:      item.Password,        // 密码
		Status:        req.Status,           // 状态(1:正常，0:禁用)
		DeptID:        req.DeptId,           // 部门ID
		LoginIP:       item.LoginIP,         // 最后登录IP
		LoginDate:     item.LoginDate,       // 最后登录时间
		LoginBrowser:  item.LoginBrowser,    // 浏览器类型
		LoginOs:       item.LoginOs,         // 操作系统
		PwdUpdateDate: item.PwdUpdateDate,   // 密码最后更新时间
		Remark:        req.Remark,           // 备注
		DelFlag:       item.DelFlag,         // 删除标志（0代表删除 1代表存在）
		CreateBy:      item.CreateBy,        // 创建者
		CreateTime:    item.CreateTime,      // 创建时间
		UpdateBy:      user.(*mw.User).Name, // 更新者
		UpdateTime:    &updateTime,          // 更新时间
	}

	err = query.Q.Transaction(func(tx *query.Query) error {
		err = dal.DB.Model(&model.SysUser{}).WithContext(ctx).Where(tx.SysUser.ID.Eq(sysUser.ID)).Save(sysUser).Error
		if err != nil {
			return err
		}

		var userPosts []*model.SysUserPost
		for _, postId := range req.PostIds {
			userPosts = append(userPosts, &model.SysUserPost{
				UserID: sysUser.ID,
				PostID: postId,
			})
		}

		postDo := tx.SysUserPost.WithContext(ctx)
		// 6.清空用户与岗位关联
		_, err = postDo.Where(tx.SysUserPost.UserID.Eq(sysUser.ID)).Delete()
		if err != nil {
			return err
		}

		// 7.添加用户与岗位关联
		err = postDo.CreateInBatches(userPosts, len(userPosts))
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	resp.Success(c, "更新用户信息成功")
}

// UpdateUserStatus 用户信息状态
// @router /api/system/user/updateUserStatus [POST]
func UpdateUserStatus(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var req user.UpdateUserStatusReq
	err := c.BindAndValidate(&req)
	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	q := query.SysUser
	ids := req.Ids // 用户id

	for _, id := range ids {
		// 1.判断是不是超级管理员
		b, _ := common.IsAdmin(ctx, id)
		if b {
			resp.Error(c, "不允许操作超级管理员用户")
			return
		}

		// 2.判断用户是否存在
		count, err := q.WithContext(ctx).Where(q.ID.Eq(id)).Count()
		if err != nil {
			resp.Error(c, err.Error())
			return
		}

		if count == 0 {
			resp.Error(c, "用户不存在")
			return
		}
	}

	// 2.更新用户状态
	_, err = q.WithContext(ctx).Where(q.ID.In(ids...)).Update(q.Status, req.Status)

	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	resp.Success(c, "更新用户信息状态成功")
}

// QueryUserDetail 查询用户信息详情
// @router /api/system/user/queryUserDetail [POST]
func QueryUserDetail(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var req user.QueryUserDetailReq
	err := c.BindAndValidate(&req)
	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	item, err := query.SysUser.WithContext(ctx).Where(query.SysUser.ID.Eq(req.Id)).First()
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		resp.Error(c, "用户信息不存在")
		return
	case err != nil:
		resp.Error(c, err.Error())
		return
	}

	data := &user.QueryUserDetailResp{
		Id:            item.ID,                                // 主键
		Mobile:        item.Mobile,                            // 手机号码
		UserName:      item.UserName,                          // 用户账号
		NickName:      item.NickName,                          // 用户昵称
		UserType:      item.UserType,                          // 用户类型（00系统用户）
		Avatar:        item.Avatar,                            // 头像路径
		Email:         item.Email,                             // 用户邮箱
		Status:        item.Status,                            // 状态(1:正常，0:禁用)
		DeptId:        item.DeptID,                            // 部门ID
		LoginIp:       item.LoginIP,                           // 最后登录IP
		LoginDate:     utils.TimeToString(item.LoginDate),     // 最后登录时间
		LoginBrowser:  item.LoginBrowser,                      // 浏览器类型
		LoginOs:       item.LoginOs,                           // 操作系统
		PwdUpdateDate: utils.TimeToString(item.PwdUpdateDate), // 密码最后更新时间
		Remark:        item.Remark,                            // 备注
		DelFlag:       item.DelFlag,                           // 删除标志（0代表删除 1代表存在）
		CreateBy:      item.CreateBy,                          // 创建者
		CreateTime:    utils.TimeToStr(item.CreateTime),       // 创建时间
		UpdateBy:      item.UpdateBy,                          // 更新者
		UpdateTime:    utils.TimeToString(item.UpdateTime),    // 更新时间

	}

	resp.Success(c, data)
}

// QueryUserList 查询用户信息列表
// @router /api/system/user/queryUserList [POST]
func QueryUserList(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var err error
	var req user.QueryUserListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	u := query.SysUser
	q := u.WithContext(ctx)
	// 手机号码
	if len(req.Mobile) > 0 {
		q = q.Where(u.Mobile.Like("%" + req.Mobile + "%"))
	}
	// 用户账号
	if len(req.UserName) > 0 {
		q = q.Where(u.UserName.Like("%" + req.UserName + "%"))
	}
	// 用户昵称
	if len(req.NickName) > 0 {
		q = q.Where(u.NickName.Like("%" + req.NickName + "%"))
	}
	// 用户邮箱
	if len(req.Email) > 0 {
		q = q.Where(u.Email.Like("%" + req.Email + "%"))
	}
	// 状态(1:正常，0:禁用)
	if req.Status != 2 {
		q = q.Where(u.Status.Eq(req.Status))
	}
	// 部门ID
	if req.DeptId != 0 {
		var deptIds []int64
		dal.DB.Model(model.SysDept{}).WithContext(ctx).Select("id").Where("find_in_set(?, ancestors)", req.DeptId).Scan(&deptIds)
		deptIds = append(deptIds, req.DeptId)
		q = q.Where(u.DeptID.In(deptIds...))
	}
	result, count, err := q.FindByPage(int((req.PageNum-1)*req.PageSize), int(req.PageSize))

	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	var list []*user.QueryUserListResp

	for _, item := range result {
		list = append(list, &user.QueryUserListResp{
			Id:            item.ID,                                // 主键
			Mobile:        item.Mobile,                            // 手机号码
			UserName:      item.UserName,                          // 用户账号
			NickName:      item.NickName,                          // 用户昵称
			UserType:      item.UserType,                          // 用户类型（00系统用户）
			Avatar:        item.Avatar,                            // 头像路径
			Email:         item.Email,                             // 用户邮箱
			Status:        item.Status,                            // 状态(1:正常，0:禁用)
			DeptId:        item.DeptID,                            // 部门ID
			LoginIp:       item.LoginIP,                           // 最后登录IP
			LoginDate:     utils.TimeToString(item.LoginDate),     // 最后登录时间
			LoginBrowser:  item.LoginBrowser,                      // 浏览器类型
			LoginOs:       item.LoginOs,                           // 操作系统
			PwdUpdateDate: utils.TimeToString(item.PwdUpdateDate), // 密码最后更新时间
			Remark:        item.Remark,                            // 备注
			DelFlag:       item.DelFlag,                           // 删除标志（0代表删除 1代表存在）
			CreateBy:      item.CreateBy,                          // 创建者
			CreateTime:    utils.TimeToStr(item.CreateTime),       // 创建时间
			UpdateBy:      item.UpdateBy,                          // 更新者
			UpdateTime:    utils.TimeToString(item.UpdateTime),    // 更新时间
		})
	}

	resp.SuccessPage(c, list, count)
}

// QueryUserMenu .
// @router /api/system/user/queryUserMenu [GET]
func QueryUserMenu(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var err error
	var req user.UserInfoReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	user1, _ := c.Get(mw.IdentityKey)
	userId := user1.(*mw.User).Id

	// 1.根据id查询用户信息
	q := query.SysUser
	info, err := q.WithContext(ctx).Where(q.ID.Eq(userId)).First()

	// 2.判断用户是否存在
	if errors.Is(err, gorm.ErrRecordNotFound) {
		resp.Error(c, "用户不存在")

		return
	}

	if err != nil {
		resp.Error(c, err.Error())

		return
	}

	// 3.查询用户菜单和权限
	menuList, apiUrls, isAdmin := queryApis(ctx, userId)

	resp.Success(c, user.UserInfoResp{
		Avatar:  info.Avatar,
		Name:    info.UserName,
		SysMenu: menuList,
		BtnMenu: apiUrls,
		IsAdmin: isAdmin,
	})
}

// 查询用户菜单和权限
func queryApis(ctx context.Context, userId int64) ([]*user.ListMenuTree, []string, bool) {
	var result []*model.SysMenu
	isAdmin, _ := common.IsAdmin(ctx, userId)

	if isAdmin {
		result, _ = query.SysMenu.WithContext(ctx).Where(query.SysMenu.Status.Eq(1)).Find()
	} else {
		sql := `
				select sm.*
				from sys_user_role sur
						 left join sys_role sr on sur.role_id = sr.id
						 left join sys_role_menu srm on sr.id = srm.role_id
						 left join sys_menu sm on srm.menu_id = sm.id
				where sur.user_id = ? and sm.status = 1
				order by sm.id
				`
		dal.DB.WithContext(ctx).Raw(sql, userId).Scan(&result)
	}
	return buildMenuTree(result, isAdmin)
}

// 构建返回值
func buildMenuTree(menus []*model.SysMenu, isAdmin bool) ([]*user.ListMenuTree, []string, bool) {
	var menuListTrees []*user.ListMenuTree
	var urls []string
	for _, menu := range menus {
		if menu.MenuType != 3 && menu.Visible == 1 {
			menuListTrees = append(menuListTrees, &user.ListMenuTree{
				Id:       menu.ID,
				Name:     menu.MenuName,
				Icon:     menu.MenuIcon,
				ParentId: menu.ParentID,
				Path:     menu.MenuURL,
			})
		}

		if len(strings.TrimSpace(menu.APIURL)) != 0 {
			urls = append(urls, menu.APIURL)
		}

	}
	return menuListTrees, urls, isAdmin
}

// QueryUserRoleList .
// @router /api/system/user/queryUserRole [POST]
func QueryUserRoleList(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var err error
	var req user.QueryUserRoleListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	count, err := query.SysUser.WithContext(ctx).Where(query.SysUser.ID.Eq(req.UserId)).Count()
	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	if count == 0 {
		resp.Error(c, "用户不存在")
		return
	}

	// 1.查询所有角色
	result, count, err := query.SysRole.WithContext(ctx).FindByPage(int((req.PageNum-1)*req.PageSize), int(req.PageSize))

	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	var roleList []*user.UserRoleListData
	var roleIds []int64

	for _, item := range result {
		roleList = append(roleList, &user.UserRoleListData{
			Id:         item.ID,                             // 主键
			RoleName:   item.RoleName,                       // 名称
			RoleKey:    item.RoleKey,                        // 角色权限字符串
			DataScope:  item.DataScope,                      // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
			Status:     item.Status,                         // 状态(1:正常，0:禁用)
			Remark:     item.Remark,                         // 备注
			DelFlag:    item.DelFlag,                        // 删除标志（0代表删除 1代表存在）
			CreateBy:   item.CreateBy,                       // 创建者
			CreateTime: utils.TimeToStr(item.CreateTime),    // 创建时间
			UpdateBy:   item.UpdateBy,                       // 更新者
			UpdateTime: utils.TimeToString(item.UpdateTime), // 更新时间
		})
		// admin账号全部角色
		roleIds = append(roleIds, item.ID)
	}

	// 2.如果角色不是admin则根据userId查询角色
	q := query.SysUserRole
	isAdmin, _ := common.IsAdmin(ctx, req.UserId)
	if !isAdmin {
		var ids []int64
		_ = q.WithContext(ctx).Select(q.RoleID).Where(q.UserID.Eq(req.UserId)).Scan(&ids)
		roleIds = ids
	}

	data := &user.QueryUserRoleListResp{
		SysRoleList: roleList,
		UserRoleIds: roleIds,
	}

	resp.Success(c, data)
}

// AddUserRole .
// @router /api/system/user/updateUserRole [POST]
func AddUserRole(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var err error
	var req user.AddUserRoleReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	// 1.判断是否为超级管理员
	isAdmin, _ := common.IsAdmin(ctx, req.UserId)
	if isAdmin {
		resp.Error(c, "不允许操作超级管理员用户")
		return
	}

	err = query.Q.Transaction(func(tx *query.Query) error {

		q := tx.SysUserRole
		userId := req.UserId

		// 2.删除用户与角色的关联
		if _, err = q.WithContext(ctx).Where(q.UserID.Eq(userId)).Delete(); err != nil {
			return err
		}

		var userRoles []*model.SysUserRole
		for _, roleId := range req.RoleIds {
			userRoles = append(userRoles, &model.SysUserRole{
				RoleID: roleId,
				UserID: userId,
			})
		}

		// 3.添加用户与角色的关联
		if err = q.WithContext(ctx).CreateInBatches(userRoles, len(userRoles)); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	resp.Success(c, "分配角色成功")
}
