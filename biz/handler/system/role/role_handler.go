package system

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/feihua/hertz-admin/biz/dal"
	"github.com/feihua/hertz-admin/biz/model/system/role"
	"github.com/feihua/hertz-admin/biz/pkg/mw"
	"github.com/feihua/hertz-admin/biz/pkg/utils"
	"github.com/feihua/hertz-admin/gen/model"
	"github.com/feihua/hertz-admin/gen/query"
	"gorm.io/gorm"
	"strconv"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
)

// AddRole 添加角色信息
// @router /api/system/role/addRole [POST]
func AddRole(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var req role.AddRoleReq
	err := c.BindAndValidate(&req)
	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	q := query.SysRole

	// 1.查询角色名称是否存在
	name := req.RoleName
	count, err := q.WithContext(ctx).Where(q.RoleName.Eq(name)).Count()

	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	if count > 0 {
		resp.Error(c, "角色名称已存在")
		return
	}

	// 2.查询权限字符是否存在
	roleKey := req.RoleKey
	count, err = q.WithContext(ctx).Where(q.RoleKey.Eq(roleKey)).Count()

	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	if count > 0 {
		resp.Error(c, "权限字符已存在")
		return
	}

	// 3.角色不存在时,则直接添加角色
	user, _ := c.Get(mw.IdentityKey)
	item := &model.SysRole{
		RoleName:  req.RoleName,         // 名称
		RoleKey:   req.RoleKey,          // 角色权限字符串
		DataScope: req.DataScope,        // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
		Status:    req.Status,           // 状态(1:正常，0:禁用)
		Remark:    req.Remark,           // 备注
		CreateBy:  user.(*mw.User).Name, // 创建者

	}

	err = q.WithContext(ctx).Create(item)

	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	resp.Success(c, "添加角色信息成功")
}

// DeleteRole 删除角色信息
// @router /api/system/role/deleteRole [POST]
func DeleteRole(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var req role.DeleteRoleReq
	err1 := c.BindAndValidate(&req)
	if err1 != nil {
		resp.Error(c, err1.Error())

		return
	}

	ids := req.Ids

	for _, roleId := range ids {

		// 1.判断是不是超级管理员
		if roleId == 1 {
			resp.Error(c, "不允许操作超级管理员角色")
			return
		}

		q := query.SysRole
		count, err := q.WithContext(ctx).Where(q.ID.Eq(roleId)).Count()

		if err != nil {
			resp.Error(c, err.Error())
			return
		}

		if count == 0 {
			resp.Error(c, "角色不存在")
			return
		}

		// 2.角色是否已使用
		q1 := query.SysUserRole
		count, err = q1.WithContext(ctx).Where(q1.RoleID.Eq(roleId)).Count()
		if err != nil {
			resp.Error(c, err.Error())
			return
		}
		if count > 0 {
			resp.Error(c, "已分配,不能删除")
			return
		}

	}

	err := query.Q.Transaction(func(tx *query.Query) error {

		// 3.删除角色
		r := tx.SysRole
		if _, err := r.WithContext(ctx).Where(r.ID.In(ids...)).Delete(); err != nil {
			return err
		}

		// 4.删除用角色与菜单的关联
		menu := tx.SysRoleMenu
		if _, err := menu.WithContext(ctx).Where(menu.RoleID.In(ids...)).Delete(); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	resp.Success(c, "删除角色信息成功")
}

// UpdateRole 更新角色信息
// @router /api/system/role/updateRole [POST]
func UpdateRole(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var req role.UpdateRoleReq
	err := c.BindAndValidate(&req)
	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	r := query.SysRole
	q := r.WithContext(ctx)

	if req.Id == 1 {
		resp.Error(c, "不允许操作超级管理员角色")
		return
	}

	// 1.查询角色角色是否已存在
	item, err := q.Where(r.ID.Eq(req.Id)).First()

	// 1.判断角色是否存在
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		resp.Error(c, "角色不存在")
		return
	case err != nil:
		resp.Error(c, err.Error())
		return
	}

	// 2.查询角色名称是否存在
	name := req.RoleName
	count, err := q.Where(r.ID.Neq(req.Id), r.RoleName.Eq(name)).Count()

	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	if count > 0 {
		resp.Error(c, "角色名称已存在")
		return
	}

	// 3.查询权限字符是否存在
	roleKey := req.RoleKey
	count, err = q.Where(r.ID.Neq(req.Id), r.RoleKey.Eq(roleKey)).Count()

	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	if count > 0 {
		resp.Error(c, "角色权限已存在")
		return
	}

	// 4.角色存在时,则直接更新角色
	user, _ := c.Get(mw.IdentityKey)
	updateTime := time.Now()
	sysRole := &model.SysRole{
		ID:         req.Id,               // 主键
		RoleName:   req.RoleName,         // 名称
		RoleKey:    req.RoleKey,          // 角色权限字符串
		DataScope:  req.DataScope,        // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
		Status:     req.Status,           // 状态(1:正常，0:禁用)
		Remark:     req.Remark,           // 备注
		CreateBy:   item.CreateBy,        // 创建者
		CreateTime: item.CreateTime,      // 创建时间
		UpdateBy:   user.(*mw.User).Name, // 更新者
		UpdateTime: &updateTime,          // 更新时间

	}

	// 2.角色信息存在时,则直接更新角色信息
	err = dal.DB.Model(&model.SysRole{}).WithContext(ctx).Where(r.ID.Eq(req.Id)).Save(sysRole).Error

	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	resp.Success(c, "更新角色信息成功")
}

// UpdateRoleStatus 角色信息状态
// @router /api/system/role/updateRoleStatus [POST]
func UpdateRoleStatus(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var req role.UpdateRoleStatusReq
	err := c.BindAndValidate(&req)
	if err != nil {
		resp.Error(c, err.Error())

		return
	}

	q := query.SysRole

	for _, roleId := range req.Ids {

		if roleId == 1 {
			resp.Error(c, "不允许操作超级管理员角色")
			return
		}
		count, err := q.WithContext(ctx).Where(q.ID.Eq(roleId)).Count()

		if err != nil {
			resp.Error(c, err.Error())
			return
		}

		if count == 0 {
			resp.Error(c, "角色不存在")
			return
		}
	}

	_, err = q.WithContext(ctx).Where(q.ID.In(req.Ids...)).Update(q.Status, req.Status)

	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	resp.Success(c, "更新角色信息状态成功")
}

// QueryRoleDetail 查询角色信息详情
// @router /api/system/role/queryRoleDetail [POST]
func QueryRoleDetail(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var req role.QueryRoleDetailReq
	err := c.BindAndValidate(&req)
	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	item, err := query.SysRole.WithContext(ctx).Where(query.SysRole.ID.Eq(req.Id)).First()
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		resp.Error(c, "角色信息不存在")
		return
	case err != nil:
		resp.Error(c, err.Error())
		return
	}

	data := &role.QueryRoleDetailResp{
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

	}

	resp.Success(c, data)
}

// QueryRoleList 查询角色信息列表
// @router /api/system/role/queryRoleList [POST]
func QueryRoleList(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var err error
	var req role.QueryRoleListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	r := query.SysRole
	q := r.WithContext(ctx)
	// 名称
	if len(req.RoleName) > 0 {
		q = q.Where(r.RoleName.Like("%" + req.RoleName + "%"))
	}
	// 角色权限字符串
	if len(req.RoleKey) > 0 {
		q = q.Where(r.RoleKey.Like("%" + req.RoleKey + "%"))
	}
	// 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
	if req.DataScope != 0 {
		q = q.Where(r.DataScope.Eq(req.DataScope))
	}
	// 状态(1:正常，0:禁用)
	if req.Status != 2 {
		q = q.Where(r.Status.Eq(req.Status))
	}

	result, count, err := q.FindByPage(int((req.PageNum-1)*req.PageSize), int(req.PageSize))

	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	var list []*role.QueryRoleListResp

	for _, item := range result {
		list = append(list, &role.QueryRoleListResp{
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
	}

	resp.SuccessPage(c, list, count)
}

// QueryRoleMenuList .
// @router /api/system/role/queryRoleMenu [POST]
func QueryRoleMenuList(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var err error
	var req role.QueryRoleMenuListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	// 1.查询所有菜单
	menus, err := query.SysMenu.WithContext(ctx).Find()
	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	var menuList []*role.RoleMenuListData
	var menuIds []int64

	for _, menu := range menus {
		menuList = append(menuList, &role.RoleMenuListData{
			Key:           strconv.FormatInt(menu.ID, 10),
			Title:         menu.MenuName,
			ParentId:      menu.ParentID,
			Id:            menu.ID,
			Label:         menu.MenuName,
			IsPenultimate: menu.ParentID == 2,
		})
		menuIds = append(menuIds, menu.ID)
	}

	// 2.如果角色不是admin则根据roleId查询菜单
	if req.RoleId != 1 {
		var ids []int64
		q := query.SysRoleMenu
		_ = q.WithContext(ctx).Select(q.MenuID).Where(q.RoleID.Eq(req.RoleId)).Scan(&ids)
		menuIds = ids
	}

	data := &role.QueryRoleMenuListResp{
		MenuList: menuList,
		MenuIds:  menuIds,
	}

	resp.Success(c, data)
}

// AddRoleMenu .
// @router /api/system/role/updateRoleMenu [POST]
func AddRoleMenu(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var err error
	var req role.AddRoleMenuReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	// role表id为1表示系统预留超级管理员角色,不用关联
	if req.RoleId == 1 {
		resp.Error(c, "不允许操作超级管理员角色")
		return
	}

	err = query.Q.Transaction(func(tx *query.Query) error {

		q := tx.SysRoleMenu
		// 1.删除角色与菜单的关联
		if _, err = q.WithContext(ctx).Where(q.RoleID.Eq(req.RoleId)).Delete(); err != nil {
			return err
		}

		var batches []*model.SysRoleMenu
		for _, menuId := range req.MenuIds {
			batches = append(batches, &model.SysRoleMenu{
				RoleID: req.RoleId,
				MenuID: menuId,
			})
		}

		// 2.添加角色与菜单的关联
		if err = q.WithContext(ctx).CreateInBatches(batches, len(batches)); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	resp.Success(c, "分配菜单成功")
}

// QueryAllocatedList .
// @router /api/system/role/queryAllocatedList [POST]
func QueryAllocatedList(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var err error
	var req role.QueryRoleUserListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	var result []model.SysUser

	ur := query.SysUserRole
	u := query.SysUser

	// 1.查询角色角色是否已存在
	count, err := query.SysRole.WithContext(ctx).Where(query.SysRole.ID.Eq(req.RoleId)).Count()

	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	if count == 0 {
		resp.Error(c, "角色不存在")
		return
	}

	q := u.WithContext(ctx).LeftJoin(ur, ur.UserID.EqCol(u.ID)).Select(u.ALL)
	q = q.Where(u.DelFlag.Eq(1))
	if len(req.Mobile) > 0 {
		q = q.Where(u.Mobile.Like("%" + req.Mobile + "%"))
	}
	if len(req.UserName) > 0 {
		q = q.Where(u.UserName.Like("%" + req.UserName + "%"))
	}

	q = q.Where(ur.RoleID.Eq(req.RoleId))

	offset := (req.PageNum - 1) * req.PageSize
	count, err = q.ScanByPage(&result, int(offset), int(req.PageSize))

	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	var list = make([]*role.QueryRoleUserListResp, 0, len(result))

	for _, item := range result {
		list = append(list, &role.QueryRoleUserListResp{
			Id:         item.ID,                             // 主键
			Mobile:     item.Mobile,                         // 手机号码
			UserName:   item.UserName,                       // 用户账号
			NickName:   item.NickName,                       // 用户昵称
			Avatar:     item.Avatar,                         // 头像路径
			Email:      item.Email,                          // 用户邮箱
			Status:     item.Status,                         // 状态(1:正常，0:禁用)
			DeptId:     item.DeptID,                         // 部门ID
			Remark:     item.Remark,                         // 备注
			CreateBy:   item.CreateBy,                       // 创建者
			CreateTime: utils.TimeToStr(item.CreateTime),    // 创建时间
			UpdateBy:   item.UpdateBy,                       // 更新者
			UpdateTime: utils.TimeToString(item.UpdateTime), // 更新时间

		})
	}

	resp.Success(c, list)
}

// QueryUnallocatedList .
// @router /api/system/role/queryUnallocatedList [POST]
func QueryUnallocatedList(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var err error
	var req role.QueryRoleUserListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	var result []model.SysUser

	ur := query.SysUserRole
	u := query.SysUser

	// 1.查询角色角色是否已存在
	count, err := query.SysRole.WithContext(ctx).Where(query.SysRole.ID.Eq(req.RoleId)).Count()

	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	if count == 0 {
		resp.Error(c, "角色不存在")
		return
	}

	q := u.WithContext(ctx).LeftJoin(ur, ur.UserID.EqCol(u.ID)).Select(u.ALL)
	q = q.Where(u.DelFlag.Eq(1))
	if len(req.Mobile) > 0 {
		q = q.Where(u.Mobile.Like("%" + req.Mobile + "%"))
	}
	if len(req.UserName) > 0 {
		q = q.Where(u.UserName.Like("%" + req.UserName + "%"))
	}

	q = q.Where(ur.WithContext(ctx).Where(ur.RoleID.IsNull()).Or(ur.RoleID.Neq(req.RoleId)))

	offset := (req.PageNum - 1) * req.PageSize
	count, err = q.ScanByPage(&result, int(offset), int(req.PageSize))

	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	var list = make([]*role.QueryRoleUserListResp, 0, len(result))

	for _, item := range result {
		list = append(list, &role.QueryRoleUserListResp{
			Id:         item.ID,                             // 主键
			Mobile:     item.Mobile,                         // 手机号码
			UserName:   item.UserName,                       // 用户账号
			NickName:   item.NickName,                       // 用户昵称
			Avatar:     item.Avatar,                         // 头像路径
			Email:      item.Email,                          // 用户邮箱
			Status:     item.Status,                         // 状态(1:正常，0:禁用)
			DeptId:     item.DeptID,                         // 部门ID
			Remark:     item.Remark,                         // 备注
			CreateBy:   item.CreateBy,                       // 创建者
			CreateTime: utils.TimeToStr(item.CreateTime),    // 创建时间
			UpdateBy:   item.UpdateBy,                       // 更新者
			UpdateTime: utils.TimeToString(item.UpdateTime), // 更新时间

		})
	}

	resp.Success(c, list)
}

// CancelAuthUser .
// @router /api/system/role/cancelAuthUser [POST]
func CancelAuthUser(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var err error
	var req role.CancelAuthUserReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	if req.RoleId == 1 {
		resp.Error(c, "不允许操作超级管理员角色")
		return
	}

	// 1.查询角色角色是否已存在
	count, err := query.SysRole.WithContext(ctx).Where(query.SysRole.ID.Eq(req.RoleId)).Count()

	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	if count == 0 {
		resp.Error(c, "角色不存在")
		return
	}

	userRole := query.SysUserRole
	q := userRole.WithContext(ctx)

	// 取消授权
	_, err = q.Where(userRole.RoleID.Eq(req.RoleId), userRole.UserID.Eq(req.UserId)).Delete()
	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	resp.Success(c, "取消授权成功")
}

// BatchCancelAuthUser .
// @router /api/system/role/batchCancelAuthUser [POST]
func BatchCancelAuthUser(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var err error
	var req role.CancelAuthUserAllReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	if req.RoleId == 1 {
		resp.Error(c, "不允许操作超级管理员角色")
		return
	}

	// 1.查询角色角色是否已存在
	count, err := query.SysRole.WithContext(ctx).Where(query.SysRole.ID.Eq(req.RoleId)).Count()

	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	if count == 0 {
		resp.Error(c, "角色不存在")
		return
	}

	userRole := query.SysUserRole
	q := userRole.WithContext(ctx)

	// 取消授权
	_, err = q.Where(userRole.RoleID.Eq(req.RoleId), userRole.UserID.In(req.UserIds...)).Delete()
	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	resp.Success(c, "取消授权成功")
}

// BatchAuthUser .
// @router /api/system/role/batchAuthUser [POST]
func BatchAuthUser(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var err error
	var req role.SelectAuthUserAllReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	if req.RoleId == 1 {
		resp.Error(c, "不允许操作超级管理员角色")
		return
	}

	// 1.查询角色角色是否已存在
	count, err := query.SysRole.WithContext(ctx).Where(query.SysRole.ID.Eq(req.RoleId)).Count()

	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	if count == 0 {
		resp.Error(c, "角色不存在")
		return
	}

	q := query.SysUserRole.WithContext(ctx)

	// 确认授权
	var userRoles []*model.SysUserRole
	for _, userId := range req.UserIds {
		sysUserRole := model.SysUserRole{
			RoleID: req.RoleId,
			UserID: userId,
		}
		userRoles = append(userRoles, &sysUserRole)
	}

	// 2.添加角色与用户的关联
	err = q.CreateInBatches(userRoles, len(userRoles))

	if err != nil {
		resp.Error(c, err.Error())

		return
	}

	resp.Success(c, "授权成功")
}
