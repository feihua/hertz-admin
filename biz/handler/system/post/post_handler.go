package system

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/feihua/hertz-admin/biz/dal"
	"github.com/feihua/hertz-admin/biz/model/system/post"
	"github.com/feihua/hertz-admin/biz/pkg/utils"
	"github.com/feihua/hertz-admin/gen/model"
	"github.com/feihua/hertz-admin/gen/query"
	"gorm.io/gorm"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
)

// AddPost 添加岗位信息表
// @router /api/system/post/addPost [POST]
func AddPost(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var req post.AddPostReq
	err := c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, resp.Error(err))
		return
	}

	p := query.SysPost
	q := p.WithContext(ctx)

	name := req.PostName
	count, err := q.Where(p.PostName.Eq(name)).Count()

	if err != nil {
		c.JSON(consts.StatusOK, resp.ErrorMsg("添加岗位信息失败"))
		return
	}

	if count > 0 {
		c.JSON(consts.StatusOK, resp.ErrorMsg(fmt.Sprintf("添加岗位信息失败,岗位名称：%s,已存在", name)))
		return
	}

	code := req.PostCode
	count, err = q.Where(p.PostCode.Eq(code)).Count()

	if err != nil {
		c.JSON(consts.StatusOK, resp.ErrorMsg("添加岗位信息失败"))
		return
	}

	if count > 0 {
		c.JSON(consts.StatusOK, resp.ErrorMsg(fmt.Sprintf("添加岗位信息失败，岗位编码：%s,已存在", code)))
		return
	}

	createBy := ctx.Value("userName").(string)
	item := &model.SysPost{
		PostCode: code,       // 岗位编码
		PostName: name,       // 岗位名称
		Sort:     req.Sort,   // 显示顺序
		Status:   req.Status, // 岗位状态（0：停用，1:正常）
		Remark:   req.Remark, // 备注
		CreateBy: createBy,   // 创建者

	}
	err = q.WithContext(ctx).Create(item)

	if err != nil {
		c.JSON(consts.StatusOK, resp.Error(err))
		return
	}

	c.JSON(consts.StatusOK, resp.Success("添加岗位信息表成功"))
}

// DeletePost 删除岗位信息表
// @router /api/system/post/deletePost [POST]
func DeletePost(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var req post.DeletePostReq
	err := c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, resp.Error(err))
		return
	}

	q := query.SysPost

	count, err := query.SysUserPost.WithContext(ctx).Where(query.SysUserPost.PostID.In(req.Ids...)).Count()
	if err != nil {
		c.JSON(consts.StatusOK, resp.ErrorMsg("查询岗位使用数量异常"))
		return
	}

	if count > 0 {
		c.JSON(consts.StatusOK, resp.ErrorMsg("删除岗位信息失败,已分配,不能删除"))
		return
	}

	_, err = q.WithContext(ctx).Where(q.ID.In(req.Ids...)).Delete()
	if err != nil {
		c.JSON(consts.StatusOK, resp.Error(err))
		return
	}

	c.JSON(consts.StatusOK, resp.Success("删除岗位信息表成功"))
}

// UpdatePost 更新岗位信息表
// @router /api/system/post/updatePost [POST]
func UpdatePost(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var req post.UpdatePostReq
	err := c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, resp.Error(err))
		return
	}

	p := query.SysPost
	q := p.WithContext(ctx)

	// 1.根据岗位信息表id查询岗位信息表是否已存在
	item, err := q.Where(p.ID.Eq(req.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		c.JSON(consts.StatusOK, resp.ErrorMsg("岗位不存在"))
		return
	case err != nil:
		c.JSON(consts.StatusOK, resp.ErrorMsg("查询岗位信息表异常"))
		return
	}

	// 2.查询postName是否被占用,如果被占用,则直接返回
	name := req.PostName
	count, err := q.Where(p.PostName.Eq(name), p.ID.Neq(req.Id)).Count()

	if err != nil {
		c.JSON(consts.StatusOK, resp.ErrorMsg("更新岗位信息失败"))
		return
	}

	if count > 0 {
		c.JSON(consts.StatusOK, resp.ErrorMsg(fmt.Sprintf("更新岗位信息失败,岗位名称：%s,已存在", name)))
		return
	}

	// 3.查询postCode是否被占用,如果被占用,则直接返回
	code := req.PostCode
	count, err = q.Where(p.PostCode.Eq(code), p.ID.Neq(req.Id)).Count()

	if err != nil {
		c.JSON(consts.StatusOK, resp.ErrorMsg("更新岗位信息失败"))
		return
	}

	if count > 0 {
		c.JSON(consts.StatusOK, resp.ErrorMsg(fmt.Sprintf("更新岗位信息失败，岗位编码：%s,已存在", code)))
		return
	}

	updateBy := ctx.Value("userName").(string)
	updateTime := time.Now()
	sysPost := &model.SysPost{
		ID:         req.Id,          // 岗位id
		PostCode:   req.PostCode,    // 岗位编码
		PostName:   req.PostName,    // 岗位名称
		Sort:       req.Sort,        // 显示顺序
		Status:     req.Status,      // 岗位状态（0：停用，1:正常）
		Remark:     req.Remark,      // 备注
		CreateBy:   item.CreateBy,   // 创建者
		CreateTime: item.CreateTime, // 创建时间
		UpdateBy:   updateBy,        // 更新者
		UpdateTime: &updateTime,     // 更新时间

	}

	// 4.岗位信息表存在时,则直接更新岗位信息表
	err = dal.DB.Model(&model.SysPost{}).WithContext(ctx).Where(p.ID.Eq(req.Id)).Save(sysPost).Error

	if err != nil {
		c.JSON(consts.StatusOK, resp.Error(err))
		return
	}

	c.JSON(consts.StatusOK, resp.Success("更新岗位信息表成功"))
}

// UpdatePostStatus 岗位信息表状态
// @router /api/system/post/updatePostStatus [POST]
func UpdatePostStatus(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var req post.UpdatePostStatusReq
	err := c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, resp.Error(err))
		return
	}

	q := query.SysPost

	_, err = q.WithContext(ctx).Where(q.ID.In(req.Ids...)).Update(q.Status, req.Status)

	if err != nil {
		c.JSON(consts.StatusOK, resp.Error(err))
		return
	}

	c.JSON(consts.StatusOK, resp.Success("更新岗位信息表状态成功"))
}

// QueryPostDetail 查询岗位信息表详情
// @router /api/system/post/queryPostDetail [POST]
func QueryPostDetail(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var req post.QueryPostDetailReq
	err := c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, resp.Error(err))
		return
	}

	item, err := query.SysPost.WithContext(ctx).Where(query.SysPost.ID.Eq(req.Id)).First()
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		c.JSON(consts.StatusOK, resp.ErrorMsg("岗位信息表不存在"))
		return
	case err != nil:
		c.JSON(consts.StatusOK, resp.ErrorMsg("查询岗位信息表异常"))
		return
	}

	data := &post.QueryPostDetailResp{
		Id:         item.ID,                             // 岗位id
		PostCode:   item.PostCode,                       // 岗位编码
		PostName:   item.PostName,                       // 岗位名称
		Sort:       item.Sort,                           // 显示顺序
		Status:     item.Status,                         // 岗位状态（0：停用，1:正常）
		Remark:     item.Remark,                         // 备注
		CreateBy:   item.CreateBy,                       // 创建者
		CreateTime: utils.TimeToStr(item.CreateTime),    // 创建时间
		UpdateBy:   item.UpdateBy,                       // 更新者
		UpdateTime: utils.TimeToString(item.UpdateTime), // 更新时间

	}

	c.JSON(consts.StatusOK, resp.Success(data))
}

// QueryPostList 查询岗位信息表列表
// @router /api/system/post/queryPostList [POST]
func QueryPostList(ctx context.Context, c *app.RequestContext) {
	resp := utils.BaseResponse{}

	var err error
	var req post.QueryPostListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, resp.Error(err))
		return
	}

	p := query.SysPost
	q := p.WithContext(ctx)
	// 岗位编码
	if len(req.PostCode) > 0 {
		q = q.Where(p.PostCode.Like("%" + req.PostCode + "%"))
	}
	// 岗位名称
	if len(req.PostName) > 0 {
		q = q.Where(p.PostName.Like("%" + req.PostName + "%"))
	}
	// 岗位状态（0：停用，1:正常）
	if req.Status != 2 {
		q = q.Where(p.Status.Eq(req.Status))
	}

	result, count, err := q.FindByPage(int((req.PageNum-1)*req.PageSize), int(req.PageSize))

	if err != nil {
		c.JSON(consts.StatusOK, resp.Error(err))
		return
	}

	var list []*post.QueryPostListResp

	for _, item := range result {
		list = append(list, &post.QueryPostListResp{
			Id:         item.ID,                             // 岗位id
			PostCode:   item.PostCode,                       // 岗位编码
			PostName:   item.PostName,                       // 岗位名称
			Sort:       item.Sort,                           // 显示顺序
			Status:     item.Status,                         // 岗位状态（0：停用，1:正常）
			Remark:     item.Remark,                         // 备注
			CreateBy:   item.CreateBy,                       // 创建者
			CreateTime: utils.TimeToStr(item.CreateTime),    // 创建时间
			UpdateBy:   item.UpdateBy,                       // 更新者
			UpdateTime: utils.TimeToString(item.UpdateTime), // 更新时间

		})
	}

	c.JSON(consts.StatusOK, resp.SuccessPage(list, count))
}
