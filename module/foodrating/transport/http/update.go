package httpfood

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/genson1808/food-delivery/common"
	"gitlab.com/genson1808/food-delivery/component/appctx"
	foodbiz "gitlab.com/genson1808/food-delivery/module/food/business"
	foodmodel "gitlab.com/genson1808/food-delivery/module/food/model"
	foodstorage "gitlab.com/genson1808/food-delivery/module/food/storage"
	"net/http"
)

func Update(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var data foodmodel.FoodUpdate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := foodstorage.NewStore(db)
		biz := foodbiz.NewUpdateFoodBiz(store)
		err = biz.Update(c.Request.Context(), int(uid.GetLocalID()), &data)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}

}
