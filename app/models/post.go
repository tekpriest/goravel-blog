package models

import (
	"github.com/goravel/framework/database/orm"
)

type Post struct {
	orm.Model
	orm.Timestamps
	User *User
	orm.SoftDeletes
	ID          string `gorm:"primaryKey" json:"id"`
	Title       string
	Body        string
	UserID      string
	IsPublished bool
}
