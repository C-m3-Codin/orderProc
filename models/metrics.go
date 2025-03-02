package models

import "time"

type Metrics struct {
	PendingCount          int64
	Proccessed            int64
	Completed             int64
	TotalCount            int64
	AverageProcessingTime time.Duration
}
