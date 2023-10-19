package utils

import (
	"encoding/json"
	"github.com/qiushenglei/gin-skeleton/pkg/errorpkg"
)

func Interface2Struct(src, dst interface{}) {
	bytes, err := json.Marshal(src)
	if err != nil {
		panic(errorpkg.NewBizErrx(errorpkg.CodeFalse, "json marshal false"))
	}

	err = json.Unmarshal(bytes, dst)
	if err != nil {
		panic(errorpkg.NewBizErrx(errorpkg.CodeFalse, "json unmarshal false"))
	}
}

// JsonBytes2String RequestBody里都是raw数据，byte数组里有回车和tab键，可以用json marshal重新组装给没有回车和tab的string
func JsonBytes2String(src []byte, dst interface{}) []byte {

	err := json.Unmarshal(src, dst)
	if err != nil {
		panic(errorpkg.NewBizErrx(errorpkg.CodeFalse, "json unmarshal false"))
	}
	b, err := json.Marshal(dst)
	if err != nil {
		panic(errorpkg.NewBizErrx(errorpkg.CodeFalse, "json marshal false"))
	}
	return b
}
