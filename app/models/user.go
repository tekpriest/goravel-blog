package models

import (
	corm "github.com/goravel/framework/contracts/database/orm"
	"github.com/goravel/framework/database/orm"
	"github.com/goravel/framework/facades"
)

type User struct {
	orm.Timestamps
	orm.SoftDeletes
	ID       string         `gorm:"primaryKey" json:"id"`
	Name     string         `json:"name"`
	Email    string         `json:"email"`
	Password string         `json:"-" gorm:"-:all"`
	Username string         `json:"username" gorm:"-:all"`
	Avatar   string         `json:"avatar,omitempty" gorm:"-:all"`
	Posts    []*Post        `json:"posts,omitempty"`
	Data     map[string]any `json:"-" gorm:"data"`
}

// WARN: doesn't work
func (u *User) DispatchesEvents() map[corm.EventType]func(corm.Event) error {
	return map[corm.EventType]func(corm.Event) error{
		corm.EventCreating: func(e corm.Event) error {
			password, _ := facades.Hash().Make(u.Password)
			e.SetAttribute("password", password)
			return nil
		},
		corm.EventSaving: func(e corm.Event) error {
			password, _ := facades.Hash().Make(u.Password)
			u.Password = password
			return nil
		},
	}
}

func (u *User) SetPassword(password string) {
	u.Password, _ = facades.Hash().Make(u.Password)
}
