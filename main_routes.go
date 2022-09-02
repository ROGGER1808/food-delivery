package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/genson1808/food-delivery/foundation/appctx"
	"gitlab.com/genson1808/food-delivery/middleware"
	httprestaurant "gitlab.com/genson1808/food-delivery/module/restaurant/transport/http"
	httprestaurantlike "gitlab.com/genson1808/food-delivery/module/restaurantlike/transport/http"
	httpupload "gitlab.com/genson1808/food-delivery/module/upload/transport/http"
	httpuser "gitlab.com/genson1808/food-delivery/module/user/transport/http"
	"net/http"
)

func setupRoutes(appContext appctx.AppContext, v1 *gin.RouterGroup) {

	v1.GET("/ping", func(c *gin.Context) {
		c.JSONP(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1.POST("/upload", httpupload.UploadImage(appContext))
	v1.POST("/authenticate", httpuser.Login(appContext))
	v1.POST("/register", httpuser.Register(appContext))
	v1.GET("/profile", httpuser.GetProfile(appContext))

	restaurants := v1.Group("/restaurants", middleware.Authenticate(appContext))
	restaurants.POST("/", httprestaurant.Create(appContext))
	restaurants.GET("/:id", httprestaurant.Get(appContext))
	restaurants.DELETE("/:id", httprestaurant.Delete(appContext))
	restaurants.PATCH("/:id", httprestaurant.Update(appContext))
	restaurants.GET("/", httprestaurant.List(appContext))
	restaurants.POST("/:id/like", httprestaurantlike.UserLikeRestaurant(appContext))
	restaurants.DELETE("/:id/dislike", httprestaurantlike.UserDislikeRestaurant(appContext))
	restaurants.GET("/:id/liked_users", httprestaurantlike.GetUserLikedRestaurant(appContext))
}
