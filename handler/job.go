package handler

import (
	"log"

	"github.com/kataras/iris"
	"github.com/tpphu/visual-job/model"
	"github.com/tpphu/visual-job/repo"
)

type jobHandlerImpl struct {
	jobRepo repo.JobRepo
	log     *log.Logger
}

func (handler jobHandlerImpl) inject(app *iris.Application) {
	group := app.Party("/job")
	group.Post("", handler.create)
	group.Get("/{id:uint}", handler.get)
}

func (handler jobHandlerImpl) get(c iris.Context) {
	id := c.Params().GetUintDefault("id", 0)
	job, err := handler.jobRepo.Find(id)
	if err != nil {
		simpleReturnHandler(c, job, NewNotFoundErr(err))
		return
	}
	simpleReturnHandler(c, job, nil)
}

func (handler jobHandlerImpl) create(c iris.Context) {
	job := model.Job{}
	err := c.ReadJSON(&job)
	if err != nil {
		simpleReturnHandler(c, job, NewValidateErr(err))
		return
	}
	result, err := handler.jobRepo.Create(job)
	simpleReturnHandler(c, result, nil)
}
