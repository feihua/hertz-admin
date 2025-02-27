package system

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	dept "github.com/feihua/hertz-admin/biz/model/system/dept"
	"github.com/feihua/hertz-admin/biz/pkg/utils"
	"github.com/feihua/hertz-admin/gen/model"
	"github.com/feihua/hertz-admin/gen/query"
	"gorm.io/gorm"

	"github.com/cloudwego/hertz/pkg/app"
)

// AddDept 添加部门表
// @router /addDept [POST]
func AddDept(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var req dept.AddDeptReq
	err := c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, resp.Error(err))
		return
	}

	// 添加部门表
	q := query.SysDept

	item := &model.SysDept{
		ParentId:  req.ParentId,  // 父部门id
		Ancestors: req.Ancestors, // 祖级列表
		DeptName:  req.DeptName,  // 部门名称
		Sort:      req.Sort,      // 显示顺序
		Leader:    req.Leader,    // 负责人
		Phone:     req.Phone,     // 联系电话
		Email:     req.Email,     // 邮箱
		Status:    req.Status,    // 部门状态（0：停用，1:正常）
		DelFlag:   req.DelFlag,   // 删除标志（0代表删除 1代表存在）
		CreateBy:  "",            // 创建者
	}

	err = q.WithContext(ctx).Create(item)

	if err != nil {
		c.JSON(consts.StatusOK, resp.Error(err))
		return
	}

	c.JSON(consts.StatusOK, resp.Success("添加部门表成功"))
}

// DeleteDept 删除部门表
// @router /deleteDept [POST]
func DeleteDept(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var req dept.DeleteDeptReq
	err := c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, resp.Error(err))
		return
	}

	q := query.SysDept

	_, err = q.WithContext(ctx).Where(q.ID.In(req.Ids...)).Delete()
	if err != nil {
		c.JSON(consts.StatusOK, resp.Error(err))
		return
	}

	c.JSON(consts.StatusOK, resp.Success("删除部门表成功"))
}

// UpdateDept 更新部门表
// @router /updateDept [POST]
func UpdateDept(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var req dept.UpdateDeptReq
	err := c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, resp.Error(err))
		return
	}

	q := query.SysDept.WithContext(ctx)

	// 1.根据部门表id查询部门表是否已存在
	_, err = q.Where(query.SysDept.ID.Eq(req.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		c.JSON(consts.StatusOK, resp.ErrorMsg("部门表不存在"))
		return
	case err != nil:
		c.JSON(consts.StatusOK, resp.ErrorMsg("查询部门表异常"))
		return
	}

	item := &model.SysDept{
		Id:        req.Id,        // 部门id
		ParentId:  req.ParentId,  // 父部门id
		Ancestors: req.Ancestors, // 祖级列表
		DeptName:  req.DeptName,  // 部门名称
		Sort:      req.Sort,      // 显示顺序
		Leader:    req.Leader,    // 负责人
		Phone:     req.Phone,     // 联系电话
		Email:     req.Email,     // 邮箱
		Status:    req.Status,    // 部门状态（0：停用，1:正常）
		DelFlag:   req.DelFlag,   // 删除标志（0代表删除 1代表存在）
		UpdateBy:  "",            // 更新者
	}

	// 2.部门表存在时,则直接更新部门表
	_, err = q.Updates(item)

	if err != nil {
		c.JSON(consts.StatusOK, resp.Error(err))
		return
	}

	c.JSON(consts.StatusOK, resp.Success("更新部门表成功"))
}

// UpdateDeptStatus 部门表状态
// @router /updateDeptStatus [POST]
func UpdateDeptStatus(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var req dept.UpdateDeptStatusReq
	err := c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, resp.Error(err))
		return
	}

	q := query.SysDept

	_, err = q.WithContext(ctx).Where(q.ID.In(req.Ids...)).Update(q.Status, req.Status)

	if err != nil {
		c.JSON(consts.StatusOK, resp.Error(err))
		return
	}

	c.JSON(consts.StatusOK, resp.Success("更新部门表状态成功"))
}

// QueryDeptDetail 查询部门表详情
// @router /queryDeptDetail [POST]
func QueryDeptDetail(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var req dept.QueryDeptDetailReq
	err := c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, resp.Error(err))
		return
	}

	item, err := query.SysDept.WithContext(ctx).Where(query.SysDept.ID.Eq(req.Id)).First()
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		c.JSON(consts.StatusOK, resp.ErrorMsg("部门表不存在"))
		return
	case err != nil:
		c.JSON(consts.StatusOK, resp.ErrorMsg("查询部门表异常"))
		return
	}

	data := &dept.QueryDeptDetailResp{
		Id:         item.Id,                                       // 部门id
		ParentId:   item.ParentId,                                 // 父部门id
		Ancestors:  item.Ancestors,                                // 祖级列表
		DeptName:   item.DeptName,                                 // 部门名称
		Sort:       item.Sort,                                     // 显示顺序
		Leader:     item.Leader,                                   // 负责人
		Phone:      item.Phone,                                    // 联系电话
		Email:      item.Email,                                    // 邮箱
		Status:     item.Status,                                   // 部门状态（0：停用，1:正常）
		DelFlag:    item.DelFlag,                                  // 删除标志（0代表删除 1代表存在）
		CreateBy:   item.CreateBy,                                 // 创建者
		CreateTime: item.CreateTime.Format("2006-01-02 15:04:05"), // 创建时间
		UpdateBy:   item.UpdateBy,                                 // 更新者
		UpdateTime: item.UpdateTime.Format("2006-01-02 15:04:05"), // 更新时间

	}

	c.JSON(consts.StatusOK, resp.Success(data))
}

// QueryDeptList 查询部门表列表
// @router /queryDeptList [POST]
func QueryDeptList(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var err error
	var req dept.QueryDeptListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, resp.Error(err))
		return
	}

	Dept := query.SysDept
	q := Dept.WithContext(ctx)
	// 父部门id
	// q = q.Where(SysDept.ParentId.Like("%" + req.ParentId + "%"))
	// 祖级列表
	// q = q.Where(SysDept.Ancestors.Like("%" + req.Ancestors + "%"))
	// 部门名称
	// q = q.Where(SysDept.DeptName.Like("%" + req.DeptName + "%"))
	// 负责人
	// q = q.Where(SysDept.Leader.Like("%" + req.Leader + "%"))
	// 联系电话
	// q = q.Where(SysDept.Phone.Like("%" + req.Phone + "%"))
	// 邮箱
	// q = q.Where(SysDept.Email.Like("%" + req.Email + "%"))
	// 部门状态（0：停用，1:正常）
	// q = q.Where(SysDept.Status.Like("%" + req.Status + "%"))
	// 删除标志（0代表删除 1代表存在）
	// q = q.Where(SysDept.DelFlag.Like("%" + req.DelFlag + "%"))

	result, count, err := q.FindByPage(int((req.PageNum-1)*req.PageSize), int(req.PageSize))

	if err != nil {
		c.JSON(consts.StatusOK, resp.Error(err))
		return
	}

	var list []*dept.QueryDeptListResp

	for _, item := range result {
		list = append(list, &dept.QueryDeptListResp{
			Id:         item.Id,                                       // 部门id
			ParentId:   item.ParentId,                                 // 父部门id
			Ancestors:  item.Ancestors,                                // 祖级列表
			DeptName:   item.DeptName,                                 // 部门名称
			Sort:       item.Sort,                                     // 显示顺序
			Leader:     item.Leader,                                   // 负责人
			Phone:      item.Phone,                                    // 联系电话
			Email:      item.Email,                                    // 邮箱
			Status:     item.Status,                                   // 部门状态（0：停用，1:正常）
			DelFlag:    item.DelFlag,                                  // 删除标志（0代表删除 1代表存在）
			CreateBy:   item.CreateBy,                                 // 创建者
			CreateTime: item.CreateTime.Format("2006-01-02 15:04:05"), // 创建时间
			UpdateBy:   item.UpdateBy,                                 // 更新者
			UpdateTime: item.UpdateTime.Format("2006-01-02 15:04:05"), // 更新时间

		})
	}

	c.JSON(consts.StatusOK, resp.SuccessPage(list, count))
}
