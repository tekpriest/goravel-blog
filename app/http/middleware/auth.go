package middleware

import (
	"errors"
	"fmt"
	"goravel/app/models"
	"net/http"

	"github.com/goravel/framework/auth"
	chttp "github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

func Auth() chttp.Middleware {
	var user models.User

	return func(ctx chttp.Context) {
		token := ctx.Request().Header("Authorization", "")
		if token == "" {
			ctx.Request().AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if _, err := facades.Auth(ctx).Parse(token); err != nil {
			if errors.Is(err, auth.ErrorTokenExpired) {
				token, err = facades.Auth(ctx).Refresh()
				if err != nil {
					ctx.Request().AbortWithStatus(http.StatusUnauthorized)
					return
				}

				ctx.Request().Header("Authorization", fmt.Sprintf("Bearer %s", token))
			} else {
				ctx.Request().AbortWithStatus(http.StatusUnauthorized)
				return
			}
		}

		if err := facades.Auth(ctx).User(&user); err != nil {
			ctx.Request().AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.WithValue("user", user)
		ctx.Request().Next()
	}
}
