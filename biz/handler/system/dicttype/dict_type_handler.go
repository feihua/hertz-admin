package system

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/feihua/hertz-admin/biz/dal"
	dictType "github.com/feihua/hertz-admin/biz/model/system/dictType"
	"github.com/feihua/hertz-admin/biz/pkg/utils"
	"github.com/feihua/hertz-admin/gen/model"
	"github.com/feihua/hertz-admin/gen/query"
	"gorm.io/gorm"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
)

// AddDictType 添加字典类型表
// @router /api/system/dictType/addDictType [POST]
func AddDictType(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var req dictType.AddDictTypeReq
	err := c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, resp.Error(err))
		return
	}

	q := query.SysDictType

	// 1.查询字典名称是否已存在,如果字典名称已存在,则直接返回
	count, err := q.WithContext(ctx).Where(q.DictName.Eq(req.DictName)).Count()

	if err != nil {
		c.JSON(consts.StatusOK, resp.ErrorMsg("添加字典信息失败"))
		return
	}

	if count > 0 {
		c.JSON(consts.StatusOK, resp.ErrorMsg("添加字典类型失败,字典名称已存在"))
		return
	}

	// 2.查询字典类型是否已存在,如果字典类型已存在,则直接返回
	count, err = q.WithContext(ctx).Where(q.DictType.Eq(req.DictType)).Count()

	if err != nil {
		c.JSON(consts.StatusOK, resp.ErrorMsg("查询字典信息失败"))
		return
	}

	if count > 0 {
		c.JSON(consts.StatusOK, resp.ErrorMsg("添加字典类型失败,字典类型已存在"))
		return
	}

	createBy := ctx.Value("userName").(string)
	item := &model.SysDictType{
		DictName: req.DictName, // 字典名称
		DictType: req.DictType, // 字典类型
		Status:   req.Status,   // 状态（0：停用，1:正常）
		Remark:   req.Remark,   // 备注
		CreateBy: createBy,     // 创建者
	}

	err = q.WithContext(ctx).Create(item)

	if err != nil {
		c.JSON(consts.StatusOK, resp.Error(err))
		return
	}

	c.JSON(consts.StatusOK, resp.Success("添加字典类型表成功"))
}

// DeleteDictType 删除字典类型表
// @router /api/system/dictType/deleteDictType [POST]
func DeleteDictType(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var req dictType.DeleteDictTypeReq
	err := c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, resp.Error(err))
		return
	}

	q := query.SysDictType
	q1 := query.SysDictData

	ids := req.Ids
	for _, id := range ids {
		// 1.查询字典是否存在
		d, err1 := q.WithContext(ctx).Where(q.ID.Eq(id)).First()
		if err1 != nil {
			c.JSON(consts.StatusOK, resp.ErrorMsg("删除字典类型表失败"))
			return
		}

		// 2.字典类型下是否有字典数据
		count, err1 := q1.WithContext(ctx).Where(q1.DictType.Eq(d.DictType)).Count()
		if err1 != nil {
			c.JSON(consts.StatusOK, resp.ErrorMsg("删除字典类型表失败"))
			return
		}
		if count > 0 {
			c.JSON(consts.StatusOK, resp.ErrorMsg("字典类型下有字典数据，不允许删除"))
			return
		}
	}

	_, err = q.WithContext(ctx).Where(q.ID.In(ids...)).Delete()
	if err != nil {
		c.JSON(consts.StatusOK, resp.Error(err))
		return
	}

	c.JSON(consts.StatusOK, resp.Success("删除字典类型表成功"))
}

// UpdateDictType 更新字典类型表
// @router /api/system/dictType/updateDictType [POST]
func UpdateDictType(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var req dictType.UpdateDictTypeReq
	err := c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, resp.Error(err))
		return
	}

	d := query.SysDictType
	q := d.WithContext(ctx)

	// 1.根据字典id查询字典是否已存在
	res, err := query.SysDictType.WithContext(ctx).Where(query.SysDictType.ID.Eq(req.Id)).First()

	// 1.判断字典类型是否存在
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		c.JSON(consts.StatusOK, resp.ErrorMsg("字典类型不存在"))
		return
	case err != nil:
		c.JSON(consts.StatusOK, resp.ErrorMsg("查询字典类型异常"))
		return
	}

	// 2.查询字典名称是否已存在,如果字典名称已存在,则直接返回
	count, err := q.WithContext(ctx).Where(d.ID.Neq(req.Id), d.DictName.Eq(req.DictName)).Count()

	if err != nil {
		c.JSON(consts.StatusOK, resp.ErrorMsg("更新字典信息失败"))
		return
	}

	if count > 0 {
		c.JSON(consts.StatusOK, resp.ErrorMsg("更新字典类型失败,字典名称已存在"))
		return
	}

	// 3.查询字典类型是否已存在,如果字典类型已存在,则直接返回
	count, err = q.WithContext(ctx).Where(d.ID.Neq(req.Id), d.DictType.Eq(req.DictType)).Count()

	if err != nil {
		c.JSON(consts.StatusOK, resp.ErrorMsg("查询字典信息失败"))
		return
	}

	if count > 0 {
		c.JSON(consts.StatusOK, resp.ErrorMsg("更新字典类型失败,字典类型已存在"))
		return
	}

	// 4.字典存在时,则直接更新字典
	updateBy := ctx.Value("userName").(string)
	now := time.Now()
	dict := &model.SysDictType{
		ID:         req.Id,         // 编号
		DictName:   req.DictName,   // 字典名称
		DictType:   req.DictType,   // 字典类型
		Status:     req.Status,     // 字典状态
		Remark:     req.Remark,     // 备注信息
		CreateBy:   res.CreateBy,   // 创建者
		CreateTime: res.CreateTime, // 创建时间
		UpdateBy:   updateBy,       // 更新者
		UpdateTime: &now,           // 更新时间
	}
	err = dal.DB.Model(&model.SysDictType{}).WithContext(ctx).Where(d.ID.Eq(req.Id)).Save(dict).Error

	if err != nil {
		c.JSON(consts.StatusOK, resp.Error(err))
		return
	}

	c.JSON(consts.StatusOK, resp.Success("更新字典类型表成功"))
}

// UpdateDictTypeStatus 字典类型表状态
// @router /api/system/dictType/updateDictTypeStatus [POST]
func UpdateDictTypeStatus(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var req dictType.UpdateDictTypeStatusReq
	err := c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, resp.Error(err))
		return
	}

	q := query.SysDictType

	_, err = q.WithContext(ctx).Where(q.ID.In(req.Ids...)).Update(q.Status, req.Status)

	if err != nil {
		c.JSON(consts.StatusOK, resp.Error(err))
		return
	}

	c.JSON(consts.StatusOK, resp.Success("更新字典类型表状态成功"))
}

// QueryDictTypeDetail 查询字典类型表详情
// @router /api/system/dictType/queryDictTypeDetail [POST]
func QueryDictTypeDetail(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var req dictType.QueryDictTypeDetailReq
	err := c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, resp.Error(err))
		return
	}

	item, err := query.SysDictType.WithContext(ctx).Where(query.SysDictType.ID.Eq(req.Id)).First()
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		c.JSON(consts.StatusOK, resp.ErrorMsg("字典类型表不存在"))
		return
	case err != nil:
		c.JSON(consts.StatusOK, resp.ErrorMsg("查询字典类型表异常"))
		return
	}

	data := &dictType.QueryDictTypeDetailResp{
		Id:         item.ID,                             // 字典主键
		DictName:   item.DictName,                       // 字典名称
		DictType:   item.DictType,                       // 字典类型
		Status:     item.Status,                         // 状态（0：停用，1:正常）
		Remark:     item.Remark,                         // 备注
		CreateBy:   item.CreateBy,                       // 创建者
		CreateTime: utils.TimeToStr(item.CreateTime),    // 创建时间
		UpdateBy:   item.UpdateBy,                       // 更新者
		UpdateTime: utils.TimeToString(item.UpdateTime), // 更新时间

	}

	c.JSON(consts.StatusOK, resp.Success(data))
}

// QueryDictTypeList 查询字典类型表列表
// @router /api/system/dictType/queryDictTypeList [POST]
func QueryDictTypeList(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var err error
	var req dictType.QueryDictTypeListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, resp.Error(err))
		return
	}

	d := query.SysDictType
	q := d.WithContext(ctx)
	// 字典名称
	if len(req.DictName) > 0 {
		q = q.Where(d.DictName.Like("%" + req.DictName + "%"))
	}
	// 字典类型
	if len(req.DictType) > 0 {
		q = q.Where(d.DictType.Like("%" + req.DictType + "%"))
	}
	// 状态（0：停用，1:正常）
	if req.Status != 2 {
		q = q.Where(d.Status.Eq(req.Status))
	}

	result, count, err := q.FindByPage(int((req.PageNum-1)*req.PageSize), int(req.PageSize))

	if err != nil {
		c.JSON(consts.StatusOK, resp.ErrorMsg("查询字典类型表列表失败"))
		return
	}

	var list []*dictType.QueryDictTypeListResp

	for _, item := range result {
		list = append(list, &dictType.QueryDictTypeListResp{
			Id:         item.ID,                             // 字典主键
			DictName:   item.DictName,                       // 字典名称
			DictType:   item.DictType,                       // 字典类型
			Status:     item.Status,                         // 状态（0：停用，1:正常）
			Remark:     item.Remark,                         // 备注
			CreateBy:   item.CreateBy,                       // 创建者
			CreateTime: utils.TimeToStr(item.CreateTime),    // 创建时间
			UpdateBy:   item.UpdateBy,                       // 更新者
			UpdateTime: utils.TimeToString(item.UpdateTime), // 更新时间
		})
	}

	c.JSON(consts.StatusOK, resp.SuccessPage(list, count))
}
