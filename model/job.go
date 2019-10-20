package model

import "github.com/jinzhu/gorm"

type JobState int
type JobStatus int

const (
	JOB_STATE_IDLE JobState = iota
	JOB_STATE_RUNNING
	JOB_STATE_PAUSED
	JOB_STATE_STOPPED
)

const (
	JOB_STATUS_NONE JobStatus = iota
	JOB_STATUS_SUCCESS
	JOB_STATUS_ERROR
)

type Job struct {
	gorm.Model
	Name   string
	Offset int
	Limit  int
	Total  int
	State  JobState
	Status JobStatus
}
