package routes

import (
	"goravel/app/http/controllers"
	"goravel/app/http/middleware"

	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"
)

func Api() {
	facades.Route().Prefix("auth").Group(func(router route.Router) {
		auth := controllers.NewAuthController()

		router.Post("login", auth.Login)
		router.Post("register", auth.Register)
	})

	facades.Route().Middleware(middleware.Auth()).Prefix("user").Group(func(router route.Router) {
		user := controllers.NewUserController()
		router.Get("profile", user.Profile)
	})
}
