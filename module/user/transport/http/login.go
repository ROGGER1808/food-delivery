package httpuser

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/genson1808/food-delivery/common"
	"gitlab.com/genson1808/food-delivery/foundation/appctx"
	hasher "gitlab.com/genson1808/food-delivery/foundation/hasher"
	"gitlab.com/genson1808/food-delivery/foundation/tokenprovider/jwt"
	userbusiness "gitlab.com/genson1808/food-delivery/module/user/business"
	usermodel "gitlab.com/genson1808/food-delivery/module/user/model"
	userstorage "gitlab.com/genson1808/food-delivery/module/user/storage"
	"net/http"
)

func Login(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		var data usermodel.UserCredentials
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		provider := jwt.NewTokenJWTProvider(appCtx.SecretKey())
		md5 := hasher.NewMd5Hash()
		store := userstorage.NewStore(db)
		business := userbusiness.NewLoginBiz(store, md5, provider, 60*60*24*30)

		account, err := business.Login(c.Request.Context(), &data)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}
