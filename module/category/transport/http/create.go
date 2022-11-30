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

func Create(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		var data categorymodel.CategoryCreate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		store := categorystorage.NewStore(db)
		biz := categorybiz.NewCreateCategoryBiz(store)
		if err := biz.Create(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		data.Mask(false)

		c.JSON(http.StatusCreated, common.SimpleSuccessResponse(data.FakeId))
	}

}
