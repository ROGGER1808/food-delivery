package subscriber

import (
	"context"
	"gitlab.com/genson1808/food-delivery/common"
	"gitlab.com/genson1808/food-delivery/component/appctx"
	"gitlab.com/genson1808/food-delivery/component/asyncjob"
	"gitlab.com/genson1808/food-delivery/component/pubsub"
)

type consumerJob struct {
	Title   string
	Handler func(ctx context.Context, message *pubsub.Message) error
}

type consumerEngine struct {
	appCtx appctx.AppContext
}

func NewEngine(appCtx appctx.AppContext) *consumerEngine {
	return &consumerEngine{appCtx: appCtx}
}

func (engine *consumerEngine) Start() error {
	engine.startSubTopic(
		common.TopicUserLikeRestaurant,
		true,
		IncreaseLikeAfterUserLikeRestaurant(engine.appCtx))

	engine.startSubTopic(
		common.TopicUserDislikeRestaurant,
		true,
		DecreaseLikeAfterUserLikeRestaurant(engine.appCtx))

	return nil
}

type GroupJob interface {
	Run(ctx context.Context) error
}

func (engine *consumerEngine) startSubTopic(topic pubsub.Topic, isConcurrent bool, consumerJobs ...consumerJob) error {
	c, _ := engine.appCtx.GetPubsub().Subscribe(context.Background(), topic)

	for _, item := range consumerJobs {
		engine.appCtx.Logger().Infow("subscriber.startSubTopic", "Setup consumer", item)
	}

	getJobHandler := func(job *consumerJob, message *pubsub.Message) asyncjob.JobHandler {
		return func(ctx context.Context) error {
			engine.appCtx.Logger().Infow(
				"subscriber.startSubTopic",
				"running job", job.Title,
				"value", message.Data())
			return job.Handler(ctx, message)
		}
	}

	go func() {
		for {
			msg := <-c

			jobHandlers := make([]asyncjob.Job, len(consumerJobs))

			for i := range consumerJobs {
				jobHandler := getJobHandler(&consumerJobs[i], msg)
				jobHandlers[i] = asyncjob.NewJob(jobHandler)
			}

			group := asyncjob.NewGroup(isConcurrent, jobHandlers...)

			if err := group.Run(context.Background()); err != nil {
				engine.appCtx.Logger().Infow("subscriber.startSubTopic", "group.run", err)
			}
		}
	}()

	return nil

}
