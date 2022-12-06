package errcode

var (
	Success                   = NewError(0, "成功")
	ServerError               = NewError(1000001, "服务内部错误")
	InvalidParams             = NewError(1000002, "入参错误")
	NotFound                  = NewError(1000003, "找不到")
	UnauthorizedAuthNotExist  = NewError(1000004, "鉴权失败,找不到对应的appKey和appSecret")
	UnauthorizedTokenError    = NewError(1000005, "鉴权失败,Token错误")
	UnauthorizedTokenTimeout  = NewError(1000006, "鉴权失败,Token超时")
	UnauthorizedTokenGenerate = NewError(1000007, "鉴权失败,Token生成失败")
	TooManyRequests           = NewError(1000008, "请求过多")
	UnauthorizedError         = NewError(1000009, "Jwt鉴权失败")

	LoginError = NewError(2000001, "登录失败")
)
