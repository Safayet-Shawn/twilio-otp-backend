package main

import (
	"time"

	"github.com/safayet-shawn/twilio-otp-app/model"
	"github.com/safayet-shawn/twilio-otp-app/queue"
	"github.com/safayet-shawn/twilio-otp-app/worker"
)

func main() {
	q := queue.NewQueue(10)
	worker.StartWork(q.Job)

	// Add OTP jobs
	q.AddTask(model.Task{Phone: "+8801813587080", Otp: "123456"})
	q.AddTask(model.Task{Phone: "+8801813587080", Otp: "654321"})

	time.Sleep(10 * time.Second) // wait for all jobs
}
