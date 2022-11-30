package httpfood

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/genson1808/food-delivery/common"
	"gitlab.com/genson1808/food-delivery/component/appctx"
	foodbiz "gitlab.com/genson1808/food-delivery/module/food/business"
	foodmodel "gitlab.com/genson1808/food-delivery/module/food/model"
	foodstorage "gitlab.com/genson1808/food-delivery/module/food/storage"
	foodlikestorage "gitlab.com/genson1808/food-delivery/module/foodlike/storage"
	"net/http"
	"strconv"
)

func ListFoodOfRestaurantId(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		restaurantId, err := strconv.Atoi(c.Param("restaurantId"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		db := appCtx.GetMainDBConnection()

		var filter foodmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		foodStore := foodstorage.NewStore(db)
		foodLikeStore := foodlikestorage.NewStore(db)

		biz := foodbiz.NewListFoodBiz(foodStore, foodLikeStore)

		userId := c.MustGet(common.CurrentUser).(common.Requester).GetUserId()

		result, err := biz.ListFoodOfRestaurant(c.Request.Context(), restaurantId, userId, &filter, &paging)
		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(true)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))

	}

}

func ListAllFood(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		var filter foodmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		foodStore := foodstorage.NewStore(db)
		foodLikeStore := foodlikestorage.NewStore(db)

		biz := foodbiz.NewListFoodBiz(foodStore, foodLikeStore)

		userId := c.MustGet(common.CurrentUser).(common.Requester).GetUserId()

		result, err := biz.ListAllFood(c.Request.Context(), userId, &paging, &filter)
		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(true)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))

	}

}
