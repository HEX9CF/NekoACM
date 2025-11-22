package vo

// 0 失败，1 成功，2 重新请求
type RespCode uint8

const (
	RespCodeError RespCode = 0
	RespCodeOk    RespCode = 1
)

func (c RespCode) String() string {
	switch c {
	case RespCodeError:
		return "错误"
	case RespCodeOk:
		return "成功"
	default:
		return "未知状态"
	}
}

// http响应体
type Resp struct {
	Code RespCode    `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func RespError(m string, d interface{}) Resp {
	return Resp{
		Code: RespCodeError,
		Msg:  m,
		Data: d,
	}
}

func RespOk(m string, d interface{}) Resp {
	return Resp{
		Code: RespCodeOk,
		Msg:  m,
		Data: d,
	}
}
