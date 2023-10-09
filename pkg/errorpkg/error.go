package errorpkg

import "fmt"

// 错误类型
const (
	cateSys = "sys"
	cateIO  = "io"
	cateBiz = "biz"
)

type Errx struct {
	code int    // 错误码
	Msg  string // 错误信息
	cate string // 错误类型
}

func (e Errx) Code() int {
	return e.code
}

func (e Errx) Error() string {
	return fmt.Sprintf("failed to %s,code:%d,cate:%s", e.Msg, e.code, e.cate)
}

func (e Errx) SetMsg(msg string) Errx {
	e.Msg = msg
	return e
}

func NewSysErrx(code int, msg string) Errx {
	return Errx{
		code: code,
		Msg:  msg,
		cate: cateSys,
	}
}

func NewBizErrx(code int, msg string) Errx {
	return Errx{
		code: code,
		Msg:  msg,
		cate: cateBiz,
	}
}

func NewIOErrx(code int, msg string) Errx {
	return Errx{
		code: code,
		Msg:  msg,
		cate: cateIO,
	}
}

const (
	codeSystem = 1001 // 系统错误

	codeLogic = 10001 // 逻辑异常
	codeParam = 10002 // 参数错误

	CodeFalse = 10086
)

var (
	ErrSystem = NewSysErrx(codeSystem, "system err")
	ErrLogic  = NewBizErrx(codeLogic, "logic err")
	ErrParam  = NewBizErrx(codeParam, "param invalid")
)
