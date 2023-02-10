package skip

import (
	"sync"
	"time"

	"github.com/robfig/cron/v3"
)

// Middleware serializes jobs, delaying subsequent runs until the
// previous one is complete. Jobs running after a delay of more than a minute
// have the delay logged at Info.
func Middleware(delay func(dur time.Duration)) cron.JobWrapper {
	if delay == nil {
		delay = func(time.Duration) {}
	}
	return func(j cron.Job) cron.Job {
		var mu sync.Mutex
		return cron.FuncJob(func() {
			start := time.Now()
			mu.Lock()
			defer mu.Unlock()
			if dur := time.Since(start); dur > time.Minute {
				delay(dur)
			}
			j.Run()
		})
	}
}
