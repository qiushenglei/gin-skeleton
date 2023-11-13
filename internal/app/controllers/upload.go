package controller

import (
	"fmt"
	"github.com/anguloc/zet/pkg/safe"
	"github.com/gin-gonic/gin"
	"github.com/qiushenglei/gin-skeleton/internal/app/entity"
	"github.com/qiushenglei/gin-skeleton/internal/app/global/utils"
	"github.com/qiushenglei/gin-skeleton/pkg/validatorx"
	"io"
	"os"
)

func UploadImg(c *gin.Context) {
	body := &entity.UploadImgRequest{}
	if err := validatorx.Validate(c, body); err != nil {
		panic(err)
	}

	file, err := body.File.Open()
	if err != nil {
		utils.Response(c, nil, err)
		return
	}
	content, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	//tmp, err := os.Create(safe.Path("/storage/img/" + body.File.Filename))
	//if _, err = tmp.Write(content); err != nil {
	//	panic(err)
	//}

	tmp, err := os.OpenFile(safe.Path("/storage/img/"+body.File.Filename), os.O_CREATE, 0644)
	if _, err = tmp.Write(content); err != nil {
		panic(err)
	}
	utils.Response(c, "success", nil)
}

func GetImg(c *gin.Context) {
	body := &entity.GetImgRequest{}
	if err := validatorx.Validate(c, body); err != nil {
		panic(err)
	}

	fmt.Println(body.ID)
}
