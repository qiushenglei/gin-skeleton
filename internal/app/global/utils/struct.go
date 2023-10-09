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
