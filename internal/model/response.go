package model

// 0 失败，1 成功，2 重新请求
type ResponseCode uint8

const (
	ResponseCodeError ResponseCode = 0
	ResponseCodeOk    ResponseCode = 1
)

func (c ResponseCode) String() string {
	switch c {
	case ResponseCodeError:
		return "错误"
	case ResponseCodeOk:
		return "成功"
	default:
		return "未知状态"
	}
}

// http响应体
type Response struct {
	Code ResponseCode `json:"code"`
	Msg  string       `json:"msg"`
	Data interface{}  `json:"data"`
}

func RespError(m string, d interface{}) Response {
	return Response{
		Code: ResponseCodeError,
		Msg:  m,
		Data: d,
	}
}

func RespOk(m string, d interface{}) Response {
	return Response{
		Code: ResponseCodeOk,
		Msg:  m,
		Data: d,
	}
}
