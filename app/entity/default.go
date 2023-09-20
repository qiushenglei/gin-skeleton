package entity

var (
	EmptyData = make([]interface{}, 0) // 返回呈现给 Node.js 前端的结果就是 data: [], 而不是 data: null
)

type DefaultRequest struct {
	AppID  string `json:"app_id"`
	UserID string `json:"user_id"`
}

type DefaultResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
