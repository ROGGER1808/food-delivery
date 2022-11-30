package httprestaurant

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/genson1808/food-delivery/common"
	"gitlab.com/genson1808/food-delivery/component/appctx"
	restaurantbiz "gitlab.com/genson1808/food-delivery/module/restaurant/business"
	restaurantstorage "gitlab.com/genson1808/food-delivery/module/restaurant/storage"
	"net/http"
)

func Delete(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		store := restaurantstorage.NewStore(db)
		biz := restaurantbiz.NewDeleteRestaurantBiz(store, requester)
		err = biz.Delete(c.Request.Context(), int(uid.GetLocalID()))
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))

	}

}
