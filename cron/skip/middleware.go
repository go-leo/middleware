package skip

import (
	"github.com/robfig/cron/v3"
)

// Middleware skips an invocation of the Job if a previous invocation is
// still running. It logs skips to the given logger at Info level.
func Middleware(skip func()) cron.JobWrapper {
	if skip == nil {
		skip = func() {}
	}
	return func(j cron.Job) cron.Job {
		var ch = make(chan struct{}, 1)
		ch <- struct{}{}
		return cron.FuncJob(func() {
			select {
			case v := <-ch:
				j.Run()
				ch <- v
			default:
				skip()
			}
		})
	}
}
