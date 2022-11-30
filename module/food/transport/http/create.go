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

func Create(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		var data foodmodel.FoodCreate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		store := foodstorage.NewStore(db)
		biz := foodbiz.NewCreateFoodBiz(store)
		if err := biz.Create(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		data.Mask(false)

		c.JSON(http.StatusCreated, common.SimpleSuccessResponse(data.FakeId))
	}

}
