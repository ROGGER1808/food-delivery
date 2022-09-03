package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/genson1808/food-delivery/component/appctx"
	"gitlab.com/genson1808/food-delivery/component/logger"
	"gitlab.com/genson1808/food-delivery/component/pubsub/pblocal"
	"gitlab.com/genson1808/food-delivery/component/uploadprovider"
	"gitlab.com/genson1808/food-delivery/middleware"
	"gitlab.com/genson1808/food-delivery/subscriber"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	dsn := os.Getenv("MYSQL_DNS")

	s3BucketName := os.Getenv("S3_BUCKET_NAME")
	s3Region := os.Getenv("S3_REGION")
	s3APIKey := os.Getenv("S3_API_KEY")
	s3SecretKey := os.Getenv("S3_SECRET_KEY")
	s3Domain := os.Getenv("S3_DOMAIN")
	secretKey := os.Getenv("SYSTEM_SECRET")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	db = db.Debug()

	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)
	ps := pblocal.NewPubSub()
	log, err := logger.New("foody")
	if err != nil {
		log.Fatalln(err)
	}

	appContext := appctx.NewAppContext(db, s3Provider, secretKey, ps, log)
	psEngine := subscriber.NewEngine(appContext)
	_ = psEngine.Start()

	r := gin.Default()
	r.Use(middleware.Recover(appContext))
	v1 := r.Group("/v1")

	setupRoutes(appContext, v1)

	r.Run()
}
