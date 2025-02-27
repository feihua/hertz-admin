package utils

type BaseResponse struct {
	Code  int32       `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
	Total int64       `json:"total"`
}

func (r BaseResponse) Error(msg error) *BaseResponse {

	return &BaseResponse{
		Code: 1,
		Msg:  msg.Error(),
	}
}

func (r BaseResponse) ErrorMsg(msg error) *BaseResponse {

	return &BaseResponse{
		Code: 1,
		Msg:  msg.Error(),
	}
}

func (r BaseResponse) Success(data interface{}) BaseResponse {

	return BaseResponse{
		Code: 0,
		Msg:  "操作成功",
		Data: data,
	}
}

func (r BaseResponse) SuccessPage(data interface{}, total int64) BaseResponse {

	return BaseResponse{
		Code:  0,
		Msg:   "操作成功",
		Data:  data,
		Total: total,
	}
}
