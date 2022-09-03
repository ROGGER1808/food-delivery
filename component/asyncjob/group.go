package asyncjob

import (
	"bytes"
	"context"
	"fmt"
	"gitlab.com/genson1808/food-delivery/common"
	"log"
	"sync"
)

type group struct {
	isParallel bool
	jobs       []Job
	wg         *sync.WaitGroup
}

type jobErr struct {
	errors map[int]error
}

func (e jobErr) Error() string {
	var buffer bytes.Buffer
	for i, err := range e.errors {
		buffer.WriteString(fmt.Sprintf("#%d: %s\n", i, err.Error()))
	}
	return buffer.String()
}

func NewGroup(isParallel bool, jobs ...Job) *group {
	return &group{
		isParallel: isParallel,
		jobs:       jobs,
		wg:         new(sync.WaitGroup),
	}
}

func (g *group) Run(ctx context.Context) error {
	g.wg.Add(len(g.jobs))

	errChan := make(chan error, len(g.jobs))

	for i, _ := range g.jobs {
		if g.isParallel {
			go func(j Job) {
				defer common.AppRecover()
				errChan <- g.runJob(ctx, j)
				g.wg.Done()
			}(g.jobs[i])
			continue
		}

		errChan <- g.runJob(ctx, g.jobs[i])
		g.wg.Done()
	}

	g.wg.Wait()

	var err error
	for i := 1; i <= len(g.jobs); i++ {
		if v := <-errChan; v != nil {
			err = v
		}
	}
	return err
}

func (g *group) runJob(ctx context.Context, j Job) error {
	if err := j.Execute(ctx); err != nil {
		for {
			log.Println(err)
			if j.State() == StateRetryFailed {
				return err
			}
			if j.Retry(ctx) == nil {
				return nil
			}
		}
	}
	return nil
}
