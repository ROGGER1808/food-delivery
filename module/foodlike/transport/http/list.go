package http

//
//func ListUserLikedFood(appCtx appctx.AppContext) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		restaurantId, err := strconv.Atoi(c.Param("id"))
//		if err != nil {
//			panic(common.ErrInvalidRequest(err))
//		}
//		db := appCtx.GetMainDBConnection()
//
//		var filter foodmodel.Filter
//		if err := c.ShouldBind(&filter); err != nil {
//			panic(common.ErrInvalidRequest(err))
//		}
//
//		var paging common.Paging
//		if err := c.ShouldBind(&paging); err != nil {
//			panic(common.ErrInvalidRequest(err))
//		}
//
//		paging.Fulfill()
//
//		foodStore := foodstorage.NewStore(db)
//		foodLikeStore := foodlikestorage.NewStore(db)
//
//		biz := foodbiz.NewListFoodBiz(foodStore, foodLikeStore)
//
//		userId := c.MustGet(common.CurrentUser).(common.Requester).GetUserId()
//
//		result, err := biz.ListFoodOfRestaurant(c.Request.Context(), restaurantId, userId, &filter, &paging)
//		if err != nil {
//			panic(err)
//		}
//
//		for i := range result {
//			result[i].Mask(true)
//		}
//
//		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
//
//	}
//
//}
