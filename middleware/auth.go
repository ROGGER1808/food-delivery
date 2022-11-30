package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gitlab.com/genson1808/food-delivery/common"
	"gitlab.com/genson1808/food-delivery/component/appctx"
	"gitlab.com/genson1808/food-delivery/component/tokenprovider/jwt"
	userstorage "gitlab.com/genson1808/food-delivery/module/user/storage"
	"strings"
)

func Authenticate(appCtx appctx.AppContext) gin.HandlerFunc {
	tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())
	return func(c *gin.Context) {
		token, err := extractTokenFromHeader(c.GetHeader("Authorization"))
		if err != nil {
			panic(err)
		}

		db := appCtx.GetMainDBConnection()
		store := userstorage.NewStore(db)

		payload, err := tokenProvider.Validate(token)
		if err != nil {
			panic(err)
		}

		user, err := store.Find(c.Request.Context(), map[string]any{"id": payload.UserId})
		if err != nil {
			panic(err)
		}

		if user.Status == 0 {
			panic(common.ErrNoPermission(err))
		}

		user.Mask(false)
		c.Set(common.CurrentUser, user)
		c.Next()
	}
}

func Authorize(appCtx appctx.AppContext, roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet(common.CurrentUser).(common.Requester)

		if !authorized(user.GetRole(), roles...) {
			panic(common.ErrNoPermission(errors.New("you are not unauthorized for that action")))
		}

		c.Next()
	}
}

func ErrWrongAuthHeader(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("wrong authentication header"),
		fmt.Sprintf("ErrWrongAuthHeader"),
	)
}

func extractTokenFromHeader(authStr string) (string, error) {
	// Parse the authorization header.
	parts := strings.Split(authStr, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		err := errors.New("expected authorization header format: bearer <token>")
		return "", ErrWrongAuthHeader(err)
	}
	return parts[1], nil
}

func authorized(has string, roles ...string) bool {
	for _, want := range roles {
		if has == want {
			return true
		}
	}

	return false
}
