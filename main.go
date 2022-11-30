package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
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

	//s3BucketName := os.Getenv("S3_BUCKET_NAME")
	//s3Region := os.Getenv("S3_REGION")
	//s3APIKey := os.Getenv("S3_API_KEY")
	//s3SecretKey := os.Getenv("S3_SECRET_KEY")
	//s3Domain := os.Getenv("S3_DOMAIN")
	secretKey := os.Getenv("SYSTEM_SECRET")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	db = db.Debug()

	//s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)
	ps := pblocal.NewPubSub()
	log, err := logger.New("foody")
	if err != nil {
		log.Fatalln(err)
	}

	cldProvider := uploadprovider.NewCloudinaryProvider(
		"genson1808",
		"786551593633328",
		"b33rVC3vCkwn2ZyxO3HPAImI9y4",
		"food-app",
	)

	//appContext := appctx.NewAppContext(db, s3Provider, secretKey, ps, log)
	appContext := appctx.NewAppContext(db, cldProvider, secretKey, ps, log)
	psEngine := subscriber.NewEngine(appContext)
	_ = psEngine.Start()

	r := gin.Default()
	r.Use(middleware.Recover(appContext))
	v1 := r.Group("/v1")
	setupRoutes(appContext, v1)

	startSocketIOServer(r, appContext)

	r.Run(":8080")
}

func startSocketIOServer(engine *gin.Engine, appCtx appctx.AppContext) {
	server, _ := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{websocket.Default},
	})

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID(), " IP:", s.RemoteAddr())

		//go func() {
		//	i := 0
		//	for {
		//		i++
		//		s.Emit("test", i)
		//		time.Sleep(time.Second)
		//	}
		//}()
		return nil
	})

	server.OnEvent("/", "authenticate", func(s socketio.Conn, token string) {
		// Validate token
		// If false: s.Close(), and return

		// If true
		// => UserId
		// Fetch db find user by Id
		// Here: s belongs to who? (user_id)
		// We need a map[user_id][]socketio.Conn
		log.Println(s.ID(), token)
	})

	type A struct {
		Age int `json:"age"`
	}

	server.OnEvent("/", "notice", func(s socketio.Conn, msg A) {
		log.Println("notice:", msg.Age)
		s.Emit("reply", msg)
	})

	server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		return "recv " + msg
	})

	server.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
		// Remove socket from socket engine (from app context)
	})

	go server.Serve()

	engine.GET("/socket.io/*any", gin.WrapH(server))
	engine.POST("/socket.io/*any", gin.WrapH(server))

	engine.StaticFile("/demo/", "./demo.html")
}
