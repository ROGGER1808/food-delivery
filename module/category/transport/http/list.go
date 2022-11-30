package httpcategory

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/genson1808/food-delivery/common"
	"gitlab.com/genson1808/food-delivery/component/appctx"
	categorybiz "gitlab.com/genson1808/food-delivery/module/category/business"
	categorymodel "gitlab.com/genson1808/food-delivery/module/category/model"
	categorystorage "gitlab.com/genson1808/food-delivery/module/category/storage"
	"net/http"
)

func List(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		var filter categorymodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		store := categorystorage.NewStore(db)
		biz := categorybiz.NewListCategoryBiz(store)

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
