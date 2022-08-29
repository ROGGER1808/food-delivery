package main

import (
	"github.com/gin-gonic/gin"
	httprestaurant "gitlab.com/genson1808/food-delivery/module/restaurant/transport/http"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Restaurant struct {
	Id        int       `json:"id" gorm:"column:id;"`
	Name      string    `json:"name" gorm:"column:name;"`
	Addr      string    `json:"addr" gorm:"column:addr;"`
	Lat       float32   `json:"lat" gorm:"column:lat;"`
	Lng       float32   `json:"lng" gorm:"column:Lng;"`
	Status    int       `json:"status" gorm:"column:status;"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;"`
}

func (Restaurant) TableName() string { return "restaurants" }

type RestaurantUpdate struct {
	Name *string `json:"name" gorm:"column:name;"`
	Addr *string `json:"addr" gorm:"column:addr;"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

func main() {
	dsn := os.Getenv("MYSQL_DNS")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	var restaurants []Restaurant
	db.Find(&restaurants)
	log.Println(restaurants)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSONP(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/v1")

	v1.POST("/restaurants", httprestaurant.Create(db))

	v1.GET("/restaurants/:id", func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSONP(http.StatusBadRequest, gin.H{
				"error": err,
			})
		}

		var data Restaurant

		err = db.Where("id = ?", id).First(&data).Error
		if err != nil || err == gorm.ErrRecordNotFound {
			c.JSONP(http.StatusBadRequest, gin.H{
				"error": err,
			})
		} else {
			c.JSONP(http.StatusOK, gin.H{"data": data})
		}

	})

	v1.DELETE("/restaurants/:id", func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSONP(http.StatusBadRequest, gin.H{
				"error": err,
			})
		}

		var data Restaurant

		err = db.Table(data.TableName()).Where("id = ?", id).Delete(nil).Error
		if err != nil {
			c.JSONP(http.StatusBadRequest, gin.H{
				"error": err,
			})
		}

		c.JSONP(http.StatusOK, gin.H{"data": 1})

	})

	v1.PATCH("/restaurants/:id", func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSONP(http.StatusBadRequest, gin.H{
				"error": err,
			})
		}

		var data RestaurantUpdate

		if err := c.ShouldBind(&data); err != nil {
			c.JSONP(http.StatusBadRequest, gin.H{
				"error": err,
			})
		}

		db.Where("id = ?", id).Updates(&data)

		c.JSONP(http.StatusOK, gin.H{"data": 1})
	})

	v1.GET("/restaurants", func(c *gin.Context) {

		var data []Restaurant

		type Paging struct {
			Page  int `json:"page" form:"page"`
			Limit int `json:"limit" form:"limit"`
		}

		var paging Paging
		err := c.ShouldBind(&paging)
		if err != nil {
			c.JSONP(http.StatusBadRequest, gin.H{
				"error": err,
			})
		}

		if paging.Page <= 0 {
			paging.Page = 1
		}
		if paging.Limit <= 0 {
			paging.Limit = 5
		}

		err = db.Offset((paging.Page - 1) * paging.Limit).
			Order("id desc").
			Limit(paging.Limit).
			Find(&data).Error

		c.JSONP(http.StatusOK, gin.H{"data": data})

	})

	r.Run()
}
