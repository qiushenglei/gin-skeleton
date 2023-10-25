// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameOrder1 = "order1"

// Order1 mapped from table <order1>
type Order1 struct {
	ID         uint32     `gorm:"column:id;type:int(10) unsigned;primaryKey;autoIncrement:true" json:"id"`
	OrderID    string     `gorm:"column:order_id;type:varchar(255);not null;uniqueIndex:idx_orderid,priority:1" json:"order_id"`
	AppID      string     `gorm:"column:app_id;type:varchar(255);not null;index:idx_appid,priority:1" json:"app_id"`
	Fee        int32      `gorm:"column:fee;type:int(11);not null" json:"fee"`
	AddTime    time.Time  `gorm:"column:add_time;type:datetime;not null" json:"add_time"`
	UpdateTime *time.Time `gorm:"column:update_time;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"update_time"`
	IsDeleted  int32      `gorm:"column:is_deleted;type:tinyint(4);not null" json:"is_deleted"`
}

// TableName Order1's table name
func (*Order1) TableName() string {
	return TableNameOrder1
}
