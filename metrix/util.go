package metrix

import "time"

func benchmark(message string) func() {
	started := time.Now()
	logger.Printf("started  %s", message)
	return func() {
		logger.Printf("finished %s in %.06f", message, time.Since(started).Seconds())
	}
}
