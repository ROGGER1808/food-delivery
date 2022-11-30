package httpcategory

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/genson1808/food-delivery/common"
	"gitlab.com/genson1808/food-delivery/component/appctx"
	categorybiz "gitlab.com/genson1808/food-delivery/module/category/business"
	categorystorage "gitlab.com/genson1808/food-delivery/module/category/storage"
	"net/http"
)

func Get(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := categorystorage.NewStore(db)
		biz := categorybiz.NewFindCategoryBiz(store)
		result, err := biz.GetById(c.Request.Context(), int(uid.GetLocalID()))
		if err != nil {
			panic(err)
		}

		result.Mask(true)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))

	}

}
