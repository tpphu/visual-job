package mock

import (
	"github.com/stretchr/testify/mock"
	"github.com/tpphu/visual-job/model"
)

type JobRepoImpl struct {
	mock.Mock
}

func (self *JobRepoImpl) Find(id uint) (*model.Job, error) {
	args := self.Called(id)
	// Todo: args.Get(0) can be nil
	return args.Get(0).(*model.Job), args.Error(1)
}

func (self *JobRepoImpl) Create(job model.Job) (*model.Job, error) {
	args := self.Called(job)
	out := args.Get(0)
	if out == nil {
		return nil, args.Error(1)
	}
	return out.(*model.Job), args.Error(1)
}
