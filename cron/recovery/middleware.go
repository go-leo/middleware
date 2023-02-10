package recovery

import (
	"github.com/robfig/cron/v3"
)

func Middleware(r func(any)) cron.JobWrapper {
	return func(j cron.Job) cron.Job {
		return cron.FuncJob(func() {
			defer func() {
				if p := recover(); p != nil {
					r(p)
				}
			}()
			j.Run()
		})
	}
}
