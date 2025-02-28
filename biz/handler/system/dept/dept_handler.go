package system

import (
	"context"
	"errors"
	"fmt"
	"github.com/feihua/hertz-admin/biz/dal"
	"github.com/feihua/hertz-admin/biz/model/system/dept"
	"github.com/feihua/hertz-admin/biz/pkg/utils"
	"github.com/feihua/hertz-admin/gen/model"
	"github.com/feihua/hertz-admin/gen/query"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
)

// AddDept 添加部门表
// @router /api/system/dept/addDept [POST]
func AddDept(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var req dept.AddDeptReq
	err := c.BindAndValidate(&req)
	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	sysDept := query.SysDept
	q := sysDept.WithContext(ctx)

	// 1.根据部门名称查询部门是否已存在
	name := req.DeptName
	count, err := q.Where(sysDept.DeptName.Eq(name), sysDept.ParentID.Eq(req.ParentId)).Count()

	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	// 2.如果部门已存在,则直接返回
	if count > 0 {
		resp.Error(c, "部门名称已存在")
		return
	}

	// 3.如果父节点不为正常状态,则不允许新增子节点
	parentDept, err := q.Where(sysDept.ID.Eq(req.ParentId)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		resp.Error(c, "添加部门失败,上级部门不存在")
		return
	case err != nil:
		resp.Error(c, err.Error())
		return
	}

	if parentDept.Status != 1 {
		resp.Error(c, "部门停用，不允许新增")
		return
	}

	ancestors := fmt.Sprintf("%s,%d", parentDept.Ancestors, parentDept.ID)
	// 4.部门不存在时,则直接添加部门
	createBy := ctx.Value("userName").(string)
	item := &model.SysDept{
		ParentID:  req.ParentId, // 父部门id
		Ancestors: ancestors,    // 祖级列表
		DeptName:  req.DeptName, // 部门名称
		Sort:      req.Sort,     // 显示顺序
		Leader:    req.Leader,   // 负责人
		Phone:     req.Phone,    // 联系电话
		Email:     req.Email,    // 邮箱
		Status:    req.Status,   // 部门状态（0：停用，1:正常）
		CreateBy:  createBy,     // 创建者
	}

	err = q.WithContext(ctx).Create(item)

	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	resp.Success(c, "添加部门表成功")
}

// DeleteDept 删除部门表
// @router /api/system/dept/deleteDept [POST]
func DeleteDept(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var req dept.DeleteDeptReq
	err := c.BindAndValidate(&req)
	if err != nil {
		resp.Error(c, err.Error())
		return
	}
	q := query.SysDept

	id := req.Id
	if id == 1 {
		resp.Error(c, "一级部门,不允许删除")
		return
	}

	// 1.查询部门是否存在
	record, err := q.WithContext(ctx).Where(q.ID.Eq(id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		resp.Error(c, "部门不存在")
		return
	case err != nil:
		resp.Error(c, err.Error())
		return
	}

	// 2.判断部门状态是否为启用
	if record.Status == 1 {
		resp.Error(c, "部门状态为启用,不允许删除")
		return
	}

	// 3.查询是否有下级部门
	count, err := q.WithContext(ctx).Where(q.ParentID.Eq(id)).Count()
	if err != nil {
		resp.Error(c, err.Error())
		return
	}
	if count > 0 {
		resp.Error(c, "存在下级部门,不允许删除")
		return
	}

	// 4.查询部门是否存在用户
	count, err = query.SysUser.WithContext(ctx).Where(query.SysUser.DeptID.Eq(id)).Count()
	if err != nil {
		resp.Error(c, err.Error())
		return
	}
	if count > 0 {
		resp.Error(c, "部门存在用户,不允许删除")
		return
	}

	resp.Success(c, "删除部门表成功")
}

// UpdateDept 更新部门表
// @router /api/system/dept/updateDept [POST]
func UpdateDept(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var req dept.UpdateDeptReq
	err := c.BindAndValidate(&req)
	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	d := query.SysDept
	q := d.WithContext(ctx)

	if req.ParentId == req.Id {
		resp.Error(c, "上级部门不能为当前部门")
		return
	}

	// 1.根据部门id查询部门是否已存在
	oldDept, err := q.Where(d.ID.Eq(req.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		resp.Error(c, "部门不存在")
		return
	case err != nil:
		resp.Error(c, err.Error())
		return
	}

	// 2.根据部门parentId查询部门是否已存在
	parentDept, err := q.Where(d.ID.Eq(req.ParentId)).First()

	// 1.判断上级部门是否存在
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		resp.Error(c, "上级部门不存在")
		return
	case err != nil:
		resp.Error(c, err.Error())
		return
	}

	// 3.根据部门名称查询部门是否已存在
	deptName := req.DeptName
	count, err := q.Where(d.ID.Neq(req.Id), d.DeptName.Eq(deptName), d.ParentID.Eq(parentDept.ID)).Count()

	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	if count > 0 {
		resp.Error(c, "部门信息已存在")
		return
	}

	// 4.查询是否有下级部门
	sql := "select count(*) from sys_dept where status = 1 and del_flag = 1 and find_in_set(?, 'ancestors')"
	err = dal.DB.Raw(sql, req.Id).Count(&count).Error
	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	if count > 0 && req.Status == 0 {
		resp.Error(c, "该部门包含未停用的子部门")
		return
	}

	sql = "select * from sys_dept where find_in_set(?, 'ancestors')"
	list := make([]model.SysDept, 10)
	err = dal.DB.Model(&model.SysDept{}).Raw(sql, req.Id).Scan(list).Error
	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	// 5.修改下级部门祖级
	ancestors := fmt.Sprintf("%s,%d", parentDept.Ancestors, parentDept.ID)
	if count > 0 {
		for _, dept1 := range list {
			parentIdStr := strings.Replace(dept1.Ancestors, oldDept.Ancestors, ancestors, -1)
			_, err = q.Where(d.ID.Eq(dept1.ID)).Update(d.Ancestors, parentIdStr)
			if err != nil {
				resp.Error(c, err.Error())
				return
			}
		}
	}

	updateBy := ctx.Value("userName").(string)
	now := time.Now()
	sysDept := &model.SysDept{
		ID:         req.Id,             // 部门id
		ParentID:   req.ParentId,       // 父部门id
		Ancestors:  ancestors,          // 祖级列表
		DeptName:   req.DeptName,       // 部门名称
		Sort:       req.Sort,           // 显示顺序
		Leader:     req.Leader,         // 负责人
		Phone:      req.Phone,          // 联系电话
		Email:      req.Email,          // 邮箱
		Status:     req.Status,         // 部门状态（0：停用，1:正常）
		CreateBy:   oldDept.CreateBy,   // 创建者
		CreateTime: oldDept.CreateTime, // 创建时间
		UpdateBy:   updateBy,           // 更新者
		UpdateTime: &now,               // 更新时间
	}

	// 6.部门存在时,则直接更新部门
	err = dal.DB.Model(&model.SysDept{}).WithContext(ctx).Where(d.ID.Eq(req.Id)).Save(sysDept).Error

	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	// 7.如果该部门是启用状态，则启用该部门的所有上级部门
	split := strings.Split(ancestors, ",")
	var parentIds []int64

	for _, str := range split {
		num, _ := strconv.ParseInt(str, 10, 64)
		parentIds = append(parentIds, num)
	}

	if req.Status == 1 {
		_, err = q.Where(d.ID.In(parentIds...)).Update(d.Status, req.Status)
		if err != nil {
			resp.Error(c, err.Error())
			return
		}

	}

	resp.Success(c, "更新部门表成功")
}

// UpdateDeptStatus 部门表状态
// @router /api/system/dept/updateDeptStatus [POST]
func UpdateDeptStatus(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var req dept.UpdateDeptStatusReq
	err1 := c.BindAndValidate(&req)
	if err1 != nil {
		resp.Error(c, err1.Error())
		return
	}

	status := req.Status // 1.判断状态是否正确

	q := query.SysDept

	for _, id := range req.Ids {
		d, err := q.WithContext(ctx).Where(q.ID.Eq(id)).First()

		// 1.判断部门是否存在
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			resp.Error(c, "部门不存在")
			return
		case err != nil:
			resp.Error(c, err.Error())
			return
		}

		// 2.查询是否有下级部门
		var count int64
		sql := "select count(*) from sys_dept where status = 1 and del_flag = 1 and find_in_set(?, 'parentIds')"
		err = dal.DB.Raw(sql, id).Count(&count).Error
		if err != nil {
			resp.Error(c, err.Error())
			return
		}

		if count > 0 && req.Status == 0 {
			resp.Error(c, "该部门包含未停用的子部门")
			return
		}

		if status == 1 {
			split := strings.Split(d.Ancestors, ",")
			var parentIds []int64

			for _, str := range split {
				num, _ := strconv.ParseInt(str, 10, 64)
				parentIds = append(parentIds, num)
			}
			_, err = query.SysDept.WithContext(ctx).Where(q.ID.In(parentIds...)).Update(q.Status, status)
			if err != nil {
				resp.Error(c, err.Error())
				return
			}
		}

	}

	_, err := q.WithContext(ctx).Where(q.ID.In(req.Ids...)).Update(q.Status, status)

	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	resp.Success(c, "更新部门状态成功")
}

// QueryDeptDetail 查询部门表详情
// @router /api/system/dept/queryDeptDetail [POST]
func QueryDeptDetail(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var req dept.QueryDeptDetailReq
	err := c.BindAndValidate(&req)
	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	item, err := query.SysDept.WithContext(ctx).Where(query.SysDept.ID.Eq(req.Id)).First()
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		resp.Error(c, "部门不存在")
		return
	case err != nil:
		resp.Error(c, err.Error())
		return
	}

	data := &dept.QueryDeptDetailResp{
		Id:         item.ID,                             // 部门id
		ParentId:   item.ParentID,                       // 父部门id
		Ancestors:  item.Ancestors,                      // 祖级列表
		DeptName:   item.DeptName,                       // 部门名称
		Sort:       item.Sort,                           // 显示顺序
		Leader:     item.Leader,                         // 负责人
		Phone:      item.Phone,                          // 联系电话
		Email:      item.Email,                          // 邮箱
		Status:     item.Status,                         // 部门状态（0：停用，1:正常）
		CreateBy:   item.CreateBy,                       // 创建者
		CreateTime: utils.TimeToStr(item.CreateTime),    // 创建时间
		UpdateBy:   item.UpdateBy,                       // 更新者
		UpdateTime: utils.TimeToString(item.UpdateTime), // 更新时间

	}

	resp.Success(c, data)
}

// QueryDeptList 查询部门表列表
// @router /api/system/dept/queryDeptList [POST]
func QueryDeptList(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var err error
	var req dept.QueryDeptListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	d := query.SysDept
	q := d.WithContext(ctx)

	// 部门名称
	if len(req.DeptName) > 0 {
		q = q.Where(d.DeptName.Like("%" + req.DeptName + "%"))
	}
	// 部门状态（0：停用，1:正常）
	if req.Status != 2 {
		q = q.Where(d.Status.Eq(req.Status))
	}

	result, count, err := q.FindByPage(int((req.PageNum-1)*req.PageSize), int(req.PageSize))

	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	var list []*dept.QueryDeptListResp

	for _, item := range result {
		list = append(list, &dept.QueryDeptListResp{
			Id:         item.ID,                             // 部门id
			ParentId:   item.ParentID,                       // 父部门id
			Ancestors:  item.Ancestors,                      // 祖级列表
			DeptName:   item.DeptName,                       // 部门名称
			Sort:       item.Sort,                           // 显示顺序
			Leader:     item.Leader,                         // 负责人
			Phone:      item.Phone,                          // 联系电话
			Email:      item.Email,                          // 邮箱
			Status:     item.Status,                         // 部门状态（0：停用，1:正常）
			CreateBy:   item.CreateBy,                       // 创建者
			CreateTime: utils.TimeToStr(item.CreateTime),    // 创建时间
			UpdateBy:   item.UpdateBy,                       // 更新者
			UpdateTime: utils.TimeToString(item.UpdateTime), // 更新时间
		})
	}

	resp.SuccessPage(c, list, count)
}
