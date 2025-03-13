package scheduler

import "time"

type Scheduler struct {
	Operation string
	Frequency uint32 //
	Last      time.Time
	Result    uint8 // 1 - success, 2 - failure
}
