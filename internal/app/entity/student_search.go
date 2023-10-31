package entity

import (
	"github.com/qiushenglei/gin-skeleton/pkg/localtime"
)

type StudentSetDataRequest struct {
	Id         int                  `json:"id"`
	Username   string               `json:"username" validate:"required,gte=1,lte=5"`
	Label      []string             `json:"label" validate:"gte=0,lte=3"`
	ClassId    int                  `json:"class_id"`
	StudentId  string               `json:"student_id"`
	AddTime    *localtime.LocalTime `json:"add_time" `
	UpdateTime *localtime.LocalTime `json:"update_time"  `
	IsDeleted  int                  `json:"is_deleted" validate:"oneof=0 1"`
	Content    string               `json:"content"`
	Hobby      []string             `json:"hobby"`
	Address    []string             `json:"address"`
	ClassInfo  *ClassInfo           `json:"class_info" validate:"required" `
	ScoreInfo  []*Score             `json:"score_info"`
}

type StudentSetDataResponse struct {
	Id int `json:"id"`
}

type ClassInfo struct {
	ClassId    int                  `json:"class_id"`
	ClassName  string               `json:"class_name" validate:"required"`
	Grade      int                  `json:"grade" validate:"number" `
	AddTime    *localtime.LocalTime `json:"add_time" `
	UpdateTime *localtime.LocalTime `json:"update_time" `
}

type Score struct {
	Id          int                  `json:"id"`
	StudentId   string               `json:"student_id"`
	SubjectId   int                  `json:"subject_id"`
	SubjectInfo SubjectInfo          `json:"subject_info"`
	Score       int                  `json:"score"`
	AddTime     *localtime.LocalTime `json:"add_time" `
	UpdateTime  *localtime.LocalTime `json:"update_time" `
}

type SubjectInfo struct {
	SubjectId   int    `json:"subject_id"`
	SubjectName string `json:"subject_name"`
}

type SearchCond struct {
	ClassId   int       `json:"class_id"`
	ClassName string    `json:"class_name"`
	Username  string    `json:"username"`
	StudentId string    `json:"student_id"`
	Label     []string  `json:"label"`
	ScoreCond ScoreCond `json:"score_cond"`
	Content   string    `json:"content"`
	Address   string    `json:"address"`
	Page      int       `json:"page"`
	PageSize  int       `json:"page_size"`
}

type ScoreCond struct {
	SubjectName string `json:"subject_name"`
	Score       *int   `json:"score"` //可能查询0分的同学
}

type DelayRequest struct {
	Expire  int    `json:"expire"`
	Content string `json:"content"`
}
