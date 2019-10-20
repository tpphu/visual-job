package worker

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/tpphu/visual-job/repo"
	"github.com/tpphu/visual-job/worker/job"
)

// Pool is a manager
type Pool struct {
	db        *gorm.DB
	maxWorker int
	workers   []Worker
	jobs      chan job.Job
	stopChan  chan bool
}

// CreateWorkerPool is function to create pool
func CreateWorkerPool(maxWorker int, db *gorm.DB) Pool {
	p := Pool{
		db:        db,
		maxWorker: maxWorker,
	}
	p.jobs = make(chan job.Job, maxWorker)
	p.workers = make([]Worker, maxWorker)
	return p
}

// Start the pool
func (p Pool) Start() {
	go p.StartWorker()
	go p.LoadJob()
	<-p.stopChan
}

func (p Pool) LoadJob() {
	jobRepo := repo.JobRepoImpl{
		DB: p.db,
	}
	productRepo := repo.ProductRepoImpl{
		DB: p.db,
	}
	for {
		listJobs, err := jobRepo.List(3)
		fmt.Println(listJobs)
		if err != nil {
			time.Sleep(time.Second * 3)
			continue
		}
		for _, singleJob := range listJobs {
			job := &job.ExportProduct{
				ProductRepo: productRepo,
				Data:        singleJob,
			}
			p.Dispatch(job)
		}
	}
}

func (p Pool) StartWorker() {
	for _, w := range p.workers {
		w.Job = p.jobs
		w.Run()
	}
}

func (p Pool) Dispatch(job job.Job) {
	p.jobs <- job
}
