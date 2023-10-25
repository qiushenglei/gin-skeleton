package errorpkg

import (
	"fmt"
)

// 错误类型
const (
	cateSys = "sys"
	cateIO  = "io"
	cateBiz = "biz"
)

type Errx struct {
	code  int    // 错误码
	msg   string // 错误信息
	cate  string // 错误类型
	trace string
}

func (e Errx) Msg() string {
	return e.msg
}

func (e Errx) Cate() string {
	return e.cate
}

func (e Errx) Code() int {
	return e.code
}

func (e Errx) Error() string {
	return fmt.Sprintf("failed to %s,code:%d,cate:%s", e.Msg, e.code, e.cate)
}

func (e Errx) SetMsg(msg string) Errx {
	e.msg = msg
	return e
}

func newErrx(code int, msg string) {
	//pc := make([]uintptr, 10)
	//n := runtime.Callers(2, pc)
	//frames := runtime.CallersFrames(pc[:n])
}

func NewSysErrx(code int, msg string) Errx {
	return Errx{
		code: code,
		msg:  msg,
		cate: cateSys,
	}
}

func NewBizErrx(code int, msg string) Errx {
	return Errx{
		code: code,
		msg:  msg,
		cate: cateBiz,
	}
}

func NewIOErrx(code int, msg string) Errx {
	return Errx{
		code: code,
		msg:  msg,
		cate: cateIO,
	}
}

func NewNoGinContextErrx(code int, msg string) Errx {
	return Errx{
		code: code,
		msg:  msg,
		cate: cateIO,
	}
}

const (
	codeSystem = 1001 // 系统错误

	codeLogic = 10001 // 逻辑异常
	CodeParam = 10002 // 参数错误

	CodeFalse        = 10086
	CodeNoGinContext = 10087
	CodeNoLogin      = 10088

	CodeBodyBind = 10088
)

var (
	ErrSystem       = NewSysErrx(codeSystem, "system err")
	ErrLogic        = NewBizErrx(codeLogic, "logic err")
	ErrParam        = NewBizErrx(CodeParam, "param invalid")
	ErrNoGinContext = NewNoGinContextErrx(CodeNoGinContext, "no gin context")
	ErrNoLogin      = NewBizErrx(CodeNoLogin, "no login")
)
