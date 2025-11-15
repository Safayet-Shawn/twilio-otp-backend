package queue

import (
	"github.com/safayet-shawn/twilio-otp-app/model"
)

type queue struct {
	Job chan model.Task
}

func NewQueue(size int) *queue {
	return &queue{
		Job: make(chan model.Task, size),
	}
}
func (q *queue) AddTask(task model.Task) {
	q.Job <- task
}
