package appctx

import (
	"gitlab.com/genson1808/food-delivery/component/pubsub"
	"gitlab.com/genson1808/food-delivery/component/uploadprovider"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
	SecretKey() string
	GetPubsub() pubsub.PubSub
	Logger() *zap.SugaredLogger
}

type appCtx struct {
	db             *gorm.DB
	uploadProvider uploadprovider.UploadProvider
	secretKey      string
	ps             pubsub.PubSub
	logger         *zap.SugaredLogger
}

func NewAppContext(
	db *gorm.DB,
	uploadProvider uploadprovider.UploadProvider,
	secretKey string,
	ps pubsub.PubSub,
	logger *zap.SugaredLogger) *appCtx {
	return &appCtx{db: db, uploadProvider: uploadProvider, secretKey: secretKey, ps: ps, logger: logger}
}

func (ctx *appCtx) SecretKey() string {
	return ctx.secretKey
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) UploadProvider() uploadprovider.UploadProvider {
	return ctx.uploadProvider
}

func (ctx *appCtx) GetPubsub() pubsub.PubSub {
	return ctx.ps
}

func (ctx *appCtx) Logger() *zap.SugaredLogger {
	return ctx.logger
}
