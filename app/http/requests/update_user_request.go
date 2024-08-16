package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type UpdateUserRequest struct {
	Name     string `form:"name" json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

func (r *UpdateUserRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *UpdateUserRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"name":     "max_len:255",
		"email":    "email",
		"password": "string",
		"username": "string|alpha_dash",
		"avatar":   "image",
	}
}

func (r *UpdateUserRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *UpdateUserRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *UpdateUserRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
