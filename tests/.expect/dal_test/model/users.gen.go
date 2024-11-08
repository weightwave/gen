// Code generated by github.com/weightwave/gen. DO NOT EDIT.
// Code generated by github.com/weightwave/gen. DO NOT EDIT.
// Code generated by github.com/weightwave/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameUser = "users"

// User mapped from table <users>
type User struct {
	ID           int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
	Name         string    `gorm:"column:name;comment:oneline" json:"name"`
	Address      string    `gorm:"column:address" json:"address"`
	RegisterTime time.Time `gorm:"column:register_time" json:"register_time"`
	Alive      bool   `gorm:"column:alive;comment:multiline\nline1\nline2" json:"alive"`
	CompanyID  int64  `gorm:"column:company_id;default:666" json:"company_id"`
	PrivateURL string `gorm:"column:private_url;default:https://a.b.c" json:"private_url"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
