package entity

import (
	"mime/multipart"
)

type UploadImgRequest struct {
	File *multipart.FileHeader `form:"file"`
}

type GetImgRequest struct {
	ID int `form:"id"`
}
