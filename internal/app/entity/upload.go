package entity

import "mime/multipart"

type UploadImgRequest struct {
	File *multipart.FileHeader `form:"file"`
}
