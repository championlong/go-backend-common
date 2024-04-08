package runner

import (
	"context"
	"go-backend-common/util"
)

type Job interface {
	init() error
	run(ctx context.Context)
}

type Jobs []Job

func RegisterJobs(js ...Job) Jobs {
	var jobs Jobs
	for _, j := range js {
		jobs = append(jobs, j)
	}
	return jobs
}

func (js Jobs) Start(ctx context.Context) error {
	for _, j := range js {
		if err := j.init(); err != nil {
			return err
		}
		go util.SafeGoroutineByContext(ctx, j.run)
	}
	return nil
}
