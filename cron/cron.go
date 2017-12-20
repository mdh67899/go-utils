package cron

import (
	"time"
)

type CronScheduler struct {
	Ticker    *time.Ticker
	Quit      chan struct{}
	Quit_Done chan struct{}
}

func NewCronScheduler(duration time.Duration) *CronScheduler {
	return &CronScheduler{Ticker: time.NewTicker(duration),
		Quit:      make(chan struct{}, 1),
		Quit_Done: make(chan struct{}, 1),
	}
}

type JobScheduler struct {
	Quit      chan struct{}
	Quit_Done chan struct{}
}

func NewJobScheduler() *JobScheduler {
	return &JobScheduler{Quit: make(chan struct{}, 1),
		Quit_Done: make(chan struct{}, 1),
	}
}
