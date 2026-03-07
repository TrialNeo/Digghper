package job

import "time"

// NewCorn 每隔duration就会执行一次f
func NewCorn(duration time.Duration, f func()) {
	for {
		select {
		case <-time.After(duration):
			f()
		}
	}
}
