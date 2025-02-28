package system

import (
	"context"
	"errors"
	"fmt"
	"github.com/feihua/hertz-admin/biz/dal"
	dictData "github.com/feihua/hertz-admin/biz/model/system/dictData"
	"github.com/feihua/hertz-admin/biz/pkg/utils"
	"github.com/feihua/hertz-admin/gen/model"
	"github.com/feihua/hertz-admin/gen/query"
	"gorm.io/gorm"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
)

// AddDictData 添加字典数据表
// @router /api/system/dictData/addDictData [POST]
func AddDictData(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var req dictData.AddDictDataReq
	err := c.BindAndValidate(&req)
	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	q := query.SysDictData

	dictType := req.DictType

	// 1.根据字典类型查询字典是否已存在
	count, err := query.SysDictType.WithContext(ctx).Where(query.SysDictType.DictType.Eq(dictType)).Count()
	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	if count == 0 {
		resp.Error(c, fmt.Sprintf("新增字典数据失败,字典类型：%s,不存在", dictType))
		return
	}

	// 2.查询字典标签是否已存在,如果字典标签已存在,则直接返回
	count, err = q.WithContext(ctx).Where(q.DictLabel.Eq(req.DictLabel), q.DictType.Eq(dictType)).Count()

	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	if count > 0 {
		resp.Error(c, fmt.Sprintf("新增字典数据失败,字典标签：%s,已存在", req.DictLabel))
		return
	}

	// 3.查询字典键值是否已存在,如果字典键值已存在,则直接返回
	count, err = q.WithContext(ctx).Where(q.DictValue.Eq(req.DictValue), q.DictType.Eq(dictType)).Count()

	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	if count > 0 {
		resp.Error(c, fmt.Sprintf("新增字典数据失败,字典键值：%s,已存在", req.DictValue))
		return
	}

	// 4.如果新增字典数据是默认,则修改其他选项为非默认状态
	if req.IsDefault == "Y" {
		_, err = q.WithContext(ctx).Where(q.DictType.Eq(dictType)).Where(q.IsDefault.Eq("Y")).Update(q.IsDefault, "N")
		if err != nil {
			resp.Error(c, err.Error())
			return
		}
	}

	// 5.字典数据不存在时,则直接添加字典数据
	createBy := ctx.Value("userName").(string)
	item := &model.SysDictData{
		DictSort:  req.DictSort,  // 字典排序
		DictLabel: req.DictLabel, // 字典标签
		DictValue: req.DictValue, // 字典键值
		DictType:  req.DictType,  // 字典类型
		CSSClass:  req.CssClass,  // 样式属性（其他样式扩展）
		ListClass: req.ListClass, // 表格回显样式
		IsDefault: req.IsDefault, // 是否默认（Y是 N否）
		Status:    req.Status,    // 状态（0：停用，1:正常）
		Remark:    req.Remark,    // 备注
		CreateBy:  createBy,      // 创建者

	}

	err = q.WithContext(ctx).Create(item)

	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	resp.Success(c, "添加字典数据表成功")
}

// DeleteDictData 删除字典数据表
// @router /api/system/dictData/deleteDictData [POST]
func DeleteDictData(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var req dictData.DeleteDictDataReq
	err := c.BindAndValidate(&req)
	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	q := query.SysDictData

	_, err = q.WithContext(ctx).Where(q.ID.In(req.Ids...)).Delete()
	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	resp.Success(c, "删除字典数据表成功")
}

// UpdateDictData 更新字典数据表
// @router /api/system/dictData/updateDictData [POST]
func UpdateDictData(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var req dictData.UpdateDictDataReq
	err := c.BindAndValidate(&req)
	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	q := query.SysDictData

	dictType := req.DictType

	// 1.判断字典数据是否存在
	dictItem, err := q.WithContext(ctx).Where(query.SysDictData.ID.Eq(req.Id)).First()

	// 1.判断字典数据是否存在
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		resp.Error(c, "字典数据不存在")
		return
	case err != nil:
		resp.Error(c, err.Error())
		return
	}

	// 2.判断字典类型是否存在
	count, err := query.SysDictType.WithContext(ctx).Where(query.SysDictType.DictType.Eq(dictType)).Count()
	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	if count == 0 {
		resp.Error(c, fmt.Sprintf("更新字典数据失败,字典类型：%s,不存在", dictType))
		return
	}

	// 3.查询字典标签是否已存在,如果字典标签已存在,则直接返回
	count, err = q.WithContext(ctx).Where(q.ID.Neq(req.Id), q.DictLabel.Eq(req.DictLabel), q.DictType.Eq(dictType)).Count()

	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	if count > 0 {
		resp.Error(c, fmt.Sprintf("更新字典数据失败,字典标签：%s,已存在", req.DictLabel))
		return
	}

	// 4.查询字典键值是否已存在,如果字典键值已存在,则直接返回
	count, err = q.WithContext(ctx).Where(q.ID.Neq(req.Id), q.DictValue.Eq(req.DictValue), q.DictType.Eq(dictType)).Count()

	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	if count > 0 {
		resp.Error(c, fmt.Sprintf("更新字典数据失败,字典键值：%s,已存在", req.DictValue))
		return
	}

	// 5.如果更新字典数据是默认,则修改其他选项为非默认状态
	if req.IsDefault == "Y" {
		_, err = q.WithContext(ctx).Where(q.ID.Neq(req.Id)).Where(q.DictType.Eq(dictType)).Where(q.IsDefault.Eq("Y")).Update(q.IsDefault, "N")
		if err != nil {
			resp.Error(c, err.Error())
			return
		}
	}

	updateBy := ctx.Value("userName").(string)
	now := time.Now()
	data := &model.SysDictData{
		ID:         req.Id,              // 编号
		DictType:   dictType,            // 字典类型
		DictLabel:  req.DictLabel,       // 字典标签
		DictValue:  req.DictValue,       // 字典键值
		Status:     req.Status,          // 字典状态
		DictSort:   req.DictSort,        // 排序
		Remark:     req.Remark,          // 备注信息
		IsDefault:  req.IsDefault,       // 是否默认  0：否  1：是
		CreateBy:   dictItem.CreateBy,   // 创建者
		CreateTime: dictItem.CreateTime, // 创建时间
		UpdateBy:   updateBy,            // 更新者
		UpdateTime: &now,                // 更新时间
	}

	// 6.字典数据存在时,则直接更新字典数据
	err = dal.DB.Model(&model.SysDictData{}).WithContext(ctx).Where(q.ID.Eq(req.Id)).Save(data).Error
	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	resp.Success(c, "更新字典数据表成功")
}

// UpdateDictDataStatus 字典数据表状态
// @router /api/system/dictData/updateDictDataStatus [POST]
func UpdateDictDataStatus(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var req dictData.UpdateDictDataStatusReq
	err := c.BindAndValidate(&req)
	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	q := query.SysDictData

	_, err = q.WithContext(ctx).Where(q.ID.In(req.Ids...)).Update(q.Status, req.Status)

	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	resp.Success(c, "更新字典数据表状态成功")
}

// QueryDictDataDetail 查询字典数据表详情
// @router /api/system/dictData/queryDictDataDetail [POST]
func QueryDictDataDetail(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var req dictData.QueryDictDataDetailReq
	err := c.BindAndValidate(&req)
	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	q := query.SysDictData
	item, err := q.WithContext(ctx).Where(q.ID.Eq(req.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		resp.Error(c, "字典数据不存在")
		return
	case err != nil:
		resp.Error(c, err.Error())
		return
	}

	data := &dictData.QueryDictDataDetailResp{
		Id:         item.ID,                             // 字典编码
		DictSort:   item.DictSort,                       // 字典排序
		DictLabel:  item.DictLabel,                      // 字典标签
		DictValue:  item.DictValue,                      // 字典键值
		DictType:   item.DictType,                       // 字典类型
		CssClass:   item.CSSClass,                       // 样式属性（其他样式扩展）
		ListClass:  item.ListClass,                      // 表格回显样式
		IsDefault:  item.IsDefault,                      // 是否默认（Y是 N否）
		Status:     item.Status,                         // 状态（0：停用，1:正常）
		Remark:     item.Remark,                         // 备注
		CreateBy:   item.CreateBy,                       // 创建者
		CreateTime: utils.TimeToStr(item.CreateTime),    // 创建时间
		UpdateBy:   item.UpdateBy,                       // 更新者
		UpdateTime: utils.TimeToString(item.UpdateTime), // 更新时间

	}
	resp.Success(c, data)
}

// QueryDictDataList 查询字典数据表列表
// @router /api/system/dictData/queryDictDataList [POST]
func QueryDictDataList(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var err error
	var req dictData.QueryDictDataListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.Error(c, err.Error())
		return
	}

	d := query.SysDictData
	q := d.WithContext(ctx)
	// 字典标签
	if len(req.DictLabel) > 0 {
		q = q.Where(d.DictLabel.Like("%" + req.DictLabel + "%"))
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
		resp.Error(c, err.Error())
		return
	}

	var list []*dictData.QueryDictDataListResp

	for _, item := range result {
		list = append(list, &dictData.QueryDictDataListResp{
			Id:         item.ID,                             // 字典编码
			DictSort:   item.DictSort,                       // 字典排序
			DictLabel:  item.DictLabel,                      // 字典标签
			DictValue:  item.DictValue,                      // 字典键值
			DictType:   item.DictType,                       // 字典类型
			CssClass:   item.CSSClass,                       // 样式属性（其他样式扩展）
			ListClass:  item.ListClass,                      // 表格回显样式
			IsDefault:  item.IsDefault,                      // 是否默认（Y是 N否）
			Status:     item.Status,                         // 状态（0：停用，1:正常）
			Remark:     item.Remark,                         // 备注
			CreateBy:   item.CreateBy,                       // 创建者
			CreateTime: utils.TimeToStr(item.CreateTime),    // 创建时间
			UpdateBy:   item.UpdateBy,                       // 更新者
			UpdateTime: utils.TimeToString(item.UpdateTime), // 更新时间
		})
	}

	resp.SuccessPage(c, list, count)
}
