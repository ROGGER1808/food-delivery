package httpuser

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/genson1808/food-delivery/common"
	"gitlab.com/genson1808/food-delivery/component/appctx"
	"gitlab.com/genson1808/food-delivery/component/hasher"
	userbusiness "gitlab.com/genson1808/food-delivery/module/user/business"
	usermodel "gitlab.com/genson1808/food-delivery/module/user/model"
	userstorage "gitlab.com/genson1808/food-delivery/module/user/storage"
	"net/http"
)

func Register(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		var data usermodel.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := userstorage.NewStore(db)
		md5 := hasher.NewMd5Hash()
		business := userbusiness.NewRegisterBiz(store, md5)

		if err := business.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId))
	}
}
