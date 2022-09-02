package httprestaurantlike

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gitlab.com/genson1808/food-delivery/common"
	"gitlab.com/genson1808/food-delivery/foundation/appctx"
	restaurantlikebusiness "gitlab.com/genson1808/food-delivery/module/restaurantlike/business"
	restaurantlikemodel "gitlab.com/genson1808/food-delivery/module/restaurantlike/model"
	restaurantlikestore "gitlab.com/genson1808/food-delivery/module/restaurantlike/storage"
	"net/http"
)

func GetUserLikedRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		restaurantUID, err := common.FromBase58(c.Param("id"))

		filter := restaurantlikemodel.Filter{RestaurantId: int(restaurantUID.GetLocalID())}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if paging.FakeCursor != "" && !common.ValidBase582Int(paging.FakeCursor) {
			panic(common.ErrInvalidRequest(errors.New("cursor invalid")))
		}

		paging.Fulfill()

		store := restaurantlikestore.NewStore(db)
		biz := restaurantlikebusiness.NewRestaurantLikedBiz(store)

		result, err := biz.GetUserLikedRestaurant(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(true)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))

	}
}
