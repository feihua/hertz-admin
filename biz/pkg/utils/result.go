package utils

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type BaseResponse struct {
	Code  int32       `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
	Total int64       `json:"total"`
}

func (r BaseResponse) Error(c *app.RequestContext, msg string) {
	c.JSON(consts.StatusOK, &BaseResponse{
		Code: 1,
		Msg:  msg,
	})
}

func (r BaseResponse) Success(c *app.RequestContext, data interface{}) {
	c.JSON(consts.StatusOK, &BaseResponse{
		Code: 0,
		Msg:  "操作成功",
		Data: data,
	})

}

func (r BaseResponse) SuccessPage(c *app.RequestContext, data interface{}, total int64) {
	c.JSON(consts.StatusOK, &BaseResponse{
		Code:  0,
		Msg:   "操作成功",
		Data:  data,
		Total: total,
	})
}
