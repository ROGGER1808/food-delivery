package httpfood

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/genson1808/food-delivery/common"
	"gitlab.com/genson1808/food-delivery/component/appctx"
	foodbiz "gitlab.com/genson1808/food-delivery/module/food/business"
	foodstorage "gitlab.com/genson1808/food-delivery/module/food/storage"
	"net/http"
)

func Get(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := foodstorage.NewStore(db)
		biz := foodbiz.NewFindFoodBiz(store)
		result, err := biz.GetById(c.Request.Context(), int(uid.GetLocalID()))
		if err != nil {
			panic(err)
		}

		result.Mask(true)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))

	}

}
