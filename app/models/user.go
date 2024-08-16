package models

import (
	"fmt"

	corm "github.com/goravel/framework/contracts/database/orm"
	"github.com/goravel/framework/database/orm"
	"github.com/goravel/framework/facades"
)

type User struct {
	orm.Timestamps
	orm.SoftDeletes
	ID       string  `gorm:"primaryKey" json:"id"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Password string  `json:"-"`
	Username string  `json:"username"`
	Avatar   string  `json:"avatar,omitempty"`
	Posts    []*Post `json:"posts,omitempty"`
}

func (u *User) DispatchesEvents() map[corm.EventType]func(corm.Event) error {
	return map[corm.EventType]func(corm.Event) error{
		corm.EventCreating: func(e corm.Event) error {
			password, _ := facades.Hash().Make(u.Password)
			e.SetAttribute("password", password)
			return nil
		},
		corm.EventSaving: func(e corm.Event) error {
			fmt.Println(e)
			password, _ := facades.Hash().Make(u.Password)
			u.Password = password
			return nil
		},
	}
}
