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

func Update(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var data categorymodel.CategoryUpdate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := categorystorage.NewStore(db)
		biz := categorybiz.NewUpdateRestaurantBiz(store)
		err = biz.Update(c.Request.Context(), int(uid.GetLocalID()), &data)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}

}
