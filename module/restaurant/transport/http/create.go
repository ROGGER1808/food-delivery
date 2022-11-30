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

func Create(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		var data restaurantmodel.RestaurantCreate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		data.OwnerId = requester.GetUserId()

		store := restaurantstorage.NewStore(db)
		biz := restaurantbiz.NewCreateRestaurantBiz(store)
		if err := biz.Create(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		data.Mask(false)

		c.JSON(http.StatusCreated, common.SimpleSuccessResponse(data.FakeId))
	}

}
