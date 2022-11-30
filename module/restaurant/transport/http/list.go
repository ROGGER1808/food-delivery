package httprestaurant

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/genson1808/food-delivery/common"
	"gitlab.com/genson1808/food-delivery/component/appctx"
	restaurantbiz "gitlab.com/genson1808/food-delivery/module/restaurant/business"
	restaurantmodel "gitlab.com/genson1808/food-delivery/module/restaurant/model"
	restaurantrepo "gitlab.com/genson1808/food-delivery/module/restaurant/repository"
	restaurantstorage "gitlab.com/genson1808/food-delivery/module/restaurant/storage"
	restaurantlikestore "gitlab.com/genson1808/food-delivery/module/restaurantlike/storage"
	"net/http"
)

func List(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		var filter restaurantmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		store := restaurantstorage.NewStore(db)
		likeStore := restaurantlikestore.NewStore(db)
		repo := restaurantrepo.NewListRestaurantRepo(store, likeStore)
		biz := restaurantbiz.NewListRestaurantBiz(repo)

		result, err := biz.List(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(true)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))

	}

}
