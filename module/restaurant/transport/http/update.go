package httprestaurant

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/genson1808/food-delivery/common"
	"gitlab.com/genson1808/food-delivery/component/appctx"
	restaurantbiz "gitlab.com/genson1808/food-delivery/module/restaurant/business"
	restaurantmodel "gitlab.com/genson1808/food-delivery/module/restaurant/model"
	restaurantstorage "gitlab.com/genson1808/food-delivery/module/restaurant/storage"
	"net/http"
)

func Update(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var data restaurantmodel.RestaurantUpdate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstorage.NewStore(db)
		biz := restaurantbiz.NewUpdateRestaurantBiz(store)
		err = biz.Update(c.Request.Context(), int(uid.GetLocalID()), &data)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}

}
