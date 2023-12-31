// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameClass = "class"

// Class mapped from table <class>
type Class struct {
	ID         uint32     `gorm:"column:id;type:int(10) unsigned;primaryKey;autoIncrement:true" json:"id"`
	ClassName  string     `gorm:"column:class_name;type:varchar(255);not null" json:"class_name"`
	Grade      string     `gorm:"column:grade;type:varchar(255);not null" json:"grade"`
	AddTime    time.Time  `gorm:"column:add_time;type:datetime;not null" json:"add_time"`
	UpdateTime *time.Time `gorm:"column:update_time;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"update_time"`
}

// TableName Class's table name
func (*Class) TableName() string {
	return TableNameClass
}
