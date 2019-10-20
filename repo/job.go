package repo

import (
	"github.com/jinzhu/gorm"
	"github.com/tpphu/visual-job/model"
)

// JobRepo interface
type JobRepo interface {
	Find(uint) (*model.Job, error)
	Create(model.Job) (*model.Job, error)
	List(limit int) ([]model.Job, error)
}

// JobRepoImpl struct
type JobRepoImpl struct {
	DB *gorm.DB
}

// Find a job
func (jobRepo JobRepoImpl) Find(id uint) (*model.Job, error) {
	job := &model.Job{}
	err := jobRepo.DB.Find(job, id).Error
	return job, err
}

// Create returns a job
func (jobRepo JobRepoImpl) Create(job model.Job) (*model.Job, error) {
	err := jobRepo.DB.Create(&job).Error
	return &job, err
}

func (jobRepo JobRepoImpl) List(limit int) ([]model.Job, error) {
	if limit == 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}
	jobs := []model.Job{}
	err := jobRepo.DB.
		Where("state = ?", model.JOB_STATE_IDLE).
		Where("status = ?", model.JOB_STATUS_NONE).
		Limit(limit).
		Find(&jobs).
		Error
	return jobs, err
}
