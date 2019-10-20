package worker

import (
	"github.com/tpphu/visual-job/worker/job"
)

// Worker is the interface
type Worker struct {
	Job chan job.Job
}

func (w Worker) Run() error {
	for {
		select {
		case j := <-w.Job:
			go j.WaitCancel()
			j.Process()
		}
	}
	return nil
}
