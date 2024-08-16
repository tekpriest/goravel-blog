package controllers

import (
	"goravel/app/models"

	"github.com/goravel/framework/contracts/database/orm"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

type UserController struct {
	// Dependent services
	q orm.Query
}

func NewUserController() *UserController {
	return &UserController{
		q: facades.Orm().Query(),
		// Inject services
	}
}

func (c *UserController) Profile(ctx http.Context) http.Response {
	user := ctx.Value("user").(models.User)

	return ctx.Response().Success().Json(http.Json{"data": http.Json{
		"user": &user,
	}})
}

// func (c *UserController) Update(ctx http.Context) http.Response {
// 	_ := ctx.Value("user").(models.User)
//
// 	// c.q.Model(&models.User{}).Where("id", user.ID)
// }
