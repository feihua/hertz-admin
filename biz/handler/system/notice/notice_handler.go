package system

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/feihua/hertz-admin/biz/dal"
	"github.com/feihua/hertz-admin/biz/model/system/notice"
	"github.com/feihua/hertz-admin/biz/pkg/utils"
	"github.com/feihua/hertz-admin/gen/model"
	"github.com/feihua/hertz-admin/gen/query"
	"gorm.io/gorm"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
)

// AddNotice 添加通知公告表
// @router /api/system/notice/addNotice [POST]
func AddNotice(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var req notice.AddNoticeReq
	err := c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, resp.Error(err))
		return
	}

	title := req.NoticeTitle

	q := query.SysNotice

	count, err := q.WithContext(ctx).Where(q.NoticeTitle.Eq(title)).Count()
	if err != nil {
		c.JSON(consts.StatusOK, resp.ErrorMsg("添加通知公告表失败"))
		return
	}

	if count > 0 {
		c.JSON(consts.StatusOK, resp.ErrorMsg("公告标题已存在"))
		return
	}

	createBy := ctx.Value("userName").(string)
	item := &model.SysNotice{
		NoticeTitle:   title,             // 公告标题
		NoticeType:    req.NoticeType,    // 公告类型（1:通知,2:公告）
		NoticeContent: req.NoticeContent, // 公告内容
		Status:        req.Status,        // 公告状态（0:关闭,1:正常 ）
		Remark:        req.Remark,        // 备注
		CreateBy:      createBy,          // 创建者
	}
	err = q.WithContext(ctx).Create(item)

	if err != nil {
		c.JSON(consts.StatusOK, resp.Error(err))
		return
	}

	c.JSON(consts.StatusOK, resp.Success("添加通知公告表成功"))
}

// DeleteNotice 删除通知公告表
// @router /api/system/notice/deleteNotice [POST]
func DeleteNotice(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var req notice.DeleteNoticeReq
	err := c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, resp.Error(err))
		return
	}

	q := query.SysNotice

	_, err = q.WithContext(ctx).Where(q.ID.In(req.Ids...)).Delete()
	if err != nil {
		c.JSON(consts.StatusOK, resp.Error(err))
		return
	}

	c.JSON(consts.StatusOK, resp.Success("删除通知公告表成功"))
}

// UpdateNotice 更新通知公告表
// @router /api/system/notice/updateNotice [POST]
func UpdateNotice(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var req notice.UpdateNoticeReq
	err := c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, resp.Error(err))
		return
	}

	title := req.NoticeTitle

	q := query.SysNotice

	// 1.根据通知公告表id查询通知公告表是否已存在
	item, err := q.WithContext(ctx).Where(query.SysNotice.ID.Eq(req.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		c.JSON(consts.StatusOK, resp.ErrorMsg("通知公告不存在"))
		return
	case err != nil:
		c.JSON(consts.StatusOK, resp.ErrorMsg("查询通知公告表异常"))
		return
	}

	count, err := q.WithContext(ctx).Where(q.NoticeTitle.Eq(title), q.ID.Neq(req.Id)).Count()
	if err != nil {
		c.JSON(consts.StatusOK, resp.ErrorMsg("更新通知公告表失败"))
		return
	}

	if count > 0 {
		c.JSON(consts.StatusOK, resp.ErrorMsg("公告标题已存在"))
		return
	}

	updateBy := ctx.Value("userName").(string)
	updateTime := time.Now()
	sysNotice := &model.SysNotice{
		ID:            req.Id,            // 公告ID
		NoticeTitle:   title,             // 公告标题
		NoticeType:    req.NoticeType,    // 公告类型（1:通知,2:公告）
		NoticeContent: req.NoticeContent, // 公告内容
		Status:        req.Status,        // 公告状态（0:关闭,1:正常 ）
		Remark:        req.Remark,        // 备注
		CreateBy:      item.CreateBy,     // 创建者
		CreateTime:    item.CreateTime,   // 创建时间
		UpdateBy:      updateBy,          // 更新者
		UpdateTime:    &updateTime,       // 更新时间
	}

	// 2.通知公告表存在时,则直接更新通知公告表
	err = dal.DB.Model(&model.SysNotice{}).WithContext(ctx).Where(q.ID.Eq(req.Id)).Save(sysNotice).Error

	if err != nil {
		c.JSON(consts.StatusOK, resp.Error(err))
		return
	}

	c.JSON(consts.StatusOK, resp.Success("更新通知公告表成功"))
}

// UpdateNoticeStatus 通知公告表状态
// @router /api/system/notice/updateNoticeStatus [POST]
func UpdateNoticeStatus(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var req notice.UpdateNoticeStatusReq
	err := c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, resp.Error(err))
		return
	}

	q := query.SysNotice

	_, err = q.WithContext(ctx).Where(q.ID.In(req.Ids...)).Update(q.Status, req.Status)

	if err != nil {
		c.JSON(consts.StatusOK, resp.Error(err))
		return
	}

	c.JSON(consts.StatusOK, resp.Success("更新通知公告表状态成功"))
}

// QueryNoticeDetail 查询通知公告表详情
// @router /api/system/notice/queryNoticeDetail [POST]
func QueryNoticeDetail(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var req notice.QueryNoticeDetailReq
	err := c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, resp.Error(err))
		return
	}

	item, err := query.SysNotice.WithContext(ctx).Where(query.SysNotice.ID.Eq(req.Id)).First()
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		c.JSON(consts.StatusOK, resp.ErrorMsg("通知公告表不存在"))
		return
	case err != nil:
		c.JSON(consts.StatusOK, resp.ErrorMsg("查询通知公告表异常"))
		return
	}

	data := &notice.QueryNoticeDetailResp{
		Id:            item.ID,                             // 公告ID
		NoticeTitle:   item.NoticeTitle,                    // 公告标题
		NoticeType:    item.NoticeType,                     // 公告类型（1:通知,2:公告）
		NoticeContent: item.NoticeContent,                  // 公告内容
		Status:        item.Status,                         // 公告状态（0:关闭,1:正常 ）
		Remark:        item.Remark,                         // 备注
		CreateBy:      item.CreateBy,                       // 创建者
		CreateTime:    utils.TimeToStr(item.CreateTime),    // 创建时间
		UpdateBy:      item.UpdateBy,                       // 更新者
		UpdateTime:    utils.TimeToString(item.UpdateTime), // 更新时间

	}

	c.JSON(consts.StatusOK, resp.Success(data))
}

// QueryNoticeList 查询通知公告表列表
// @router /api/system/notice/queryNoticeList [POST]
func QueryNoticeList(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var err error
	var req notice.QueryNoticeListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, resp.Error(err))
		return
	}

	n := query.SysNotice
	q := n.WithContext(ctx)
	// 公告标题
	if len(req.NoticeTitle) > 0 {
		q = q.Where(n.NoticeTitle.Like("%" + req.NoticeTitle + "%"))
	}
	// 公告类型（1:通知,2:公告）
	if req.NoticeType != 0 {
		q = q.Where(n.NoticeType.Eq(req.NoticeType))
	}
	// 公告状态（0:关闭,1:正常 ）
	if req.Status != 2 {
		q = q.Where(n.Status.Eq(req.Status))
	}

	result, count, err := q.FindByPage(int((req.PageNum-1)*req.PageSize), int(req.PageSize))

	if err != nil {
		c.JSON(consts.StatusOK, resp.ErrorMsg("查询通知公告表列表失败"))
		return
	}

	var list []*notice.QueryNoticeListResp

	for _, item := range result {
		list = append(list, &notice.QueryNoticeListResp{
			Id:            item.ID,                             // 公告ID
			NoticeTitle:   item.NoticeTitle,                    // 公告标题
			NoticeType:    item.NoticeType,                     // 公告类型（1:通知,2:公告）
			NoticeContent: item.NoticeContent,                  // 公告内容
			Status:        item.Status,                         // 公告状态（0:关闭,1:正常 ）
			Remark:        item.Remark,                         // 备注
			CreateBy:      item.CreateBy,                       // 创建者
			CreateTime:    utils.TimeToStr(item.CreateTime),    // 创建时间
			UpdateBy:      item.UpdateBy,                       // 更新者
			UpdateTime:    utils.TimeToString(item.UpdateTime), // 更新时间
		})
	}

	c.JSON(consts.StatusOK, resp.SuccessPage(list, count))
}
