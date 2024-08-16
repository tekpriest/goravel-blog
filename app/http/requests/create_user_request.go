package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type CreateUserRequest struct {
	Name     string `form:"name" json:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

func (r *CreateUserRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *CreateUserRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"name":     "required|max_len:255",
		"email":    "required|email",
		"password": "required|string",
		"username": "required|string|alpha_dash",
	}
}

func (r *CreateUserRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *CreateUserRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *CreateUserRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
