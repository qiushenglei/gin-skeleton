// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameScore = "score"

// Score mapped from table <score>
type Score struct {
	ID         uint32     `gorm:"column:id;type:int(10) unsigned;primaryKey;autoIncrement:true" json:"id"`
	StudentID  string     `gorm:"column:student_id;type:varchar(255);not null" json:"student_id"`
	SubjectID  uint32     `gorm:"column:subject_id;type:int(10) unsigned;not null" json:"subject_id"`
	Score      int32      `gorm:"column:score;type:int(11);not null" json:"score"`
	AddTime    time.Time  `gorm:"column:add_time;type:datetime;not null" json:"add_time"`
	UpdateTime *time.Time `gorm:"column:update_time;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"update_time"`
}

// TableName Score's table name
func (*Score) TableName() string {
	return TableNameScore
}
