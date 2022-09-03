package httprestaurantlike

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/genson1808/food-delivery/common"
	"gitlab.com/genson1808/food-delivery/component/appctx"
	restaurantstorage "gitlab.com/genson1808/food-delivery/module/restaurant/storage"
	restaurantlikebusiness "gitlab.com/genson1808/food-delivery/module/restaurantlike/business"
	restaurantlikemodel "gitlab.com/genson1808/food-delivery/module/restaurantlike/model"
	restaurantlikestore "gitlab.com/genson1808/food-delivery/module/restaurantlike/storage"
	"net/http"
)

func UserLikeRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		data := restaurantlikemodel.Like{
			RestaurantId: int(uid.GetLocalID()),
			UserId:       requester.GetUserId(),
		}

		db := appCtx.GetMainDBConnection()

		store := restaurantlikestore.NewStore(db)
		restaurantStore := restaurantstorage.NewStore(db)
		biz := restaurantlikebusiness.NewUserLikeRestaurantBiz(store, restaurantStore)

		if err := biz.LikeRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
