package controllers

import (
	"goravel/app/http/requests"
	"goravel/app/models"

	"github.com/google/uuid"
	"github.com/goravel/framework/contracts/database/orm"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

type AuthController struct {
	q orm.Query
}

func NewAuthController() *AuthController {
	return &AuthController{
		q: facades.Orm().Query(),
	}
}

func (c *AuthController) Login(ctx http.Context) http.Response {
	return ctx.Response().Success().Json(http.Json{"action": "login"})
}

func (c *AuthController) Register(ctx http.Context) http.Response {
	var request requests.CreateUserRequest
	errs, _ := ctx.Request().ValidateRequest(&request)
	if errs != nil {
		return ctx.Response().Json(http.StatusBadRequest, http.Json{"errors": errs.All()})
	}
	user := &models.User{
		ID:       uuid.NewString(),
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
		Username: request.Username,
	}
	if err := c.q.Create(&user); err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{"error": err.Error()})
	}
	token, err := facades.Auth(ctx).Login(user)
	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{"error": err.Error()})
	}
	return ctx.Response().Success().Json(http.Json{"data": http.Json{
		"user":  &user,
		"token": token,
	}})
}
