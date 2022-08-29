package httprestaurant

import (
	"github.com/gin-gonic/gin"
	restaurantbiz "gitlab.com/genson1808/food-delivery/module/restaurant/business"
	restaurantmodel "gitlab.com/genson1808/food-delivery/module/restaurant/model"
	restaurantstorage "gitlab.com/genson1808/food-delivery/module/restaurant/storage"
	"gorm.io/gorm"
	"net/http"
)

func Create(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantCreate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		store := restaurantstorage.NewStore(db)
		biz := restaurantbiz.NewCreateRestaurantBiz(store)
		err := biz.Create(c.Request.Context(), &data)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"data": data,
		})
	}

}
