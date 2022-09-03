package main

import (
	"context"
	"errors"
	"gitlab.com/genson1808/food-delivery/component/asyncjob"
	"log"
)

func main() {
	job1 := asyncjob.NewJob(func(ctx context.Context) error {
		log.Println("Job 1")
		return errors.New("error for job 1")
	})

	job2 := asyncjob.NewJob(func(ctx context.Context) error {
		log.Println("Job 2")
		return nil
	})

	job3 := asyncjob.NewJob(func(ctx context.Context) error {
		log.Println("Job 3")
		return errors.New("error for job 3")
	})

	g := asyncjob.NewGroup(true, job1, job2, job3)
	if err := g.Run(context.Background()); err != nil {
		log.Println(err)
	}

}
