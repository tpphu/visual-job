package handler

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"github.com/tpphu/visual-job/repo"
	"github.com/tpphu/visual-job/service"
	"github.com/urfave/cli"
)

// BuildEngine returns a *iris.Application
func BuildEngine(appContext *cli.Context, logger *log.Logger, db *gorm.DB) *iris.Application {
	app := iris.Default()
	app.Logger().SetLevel(appContext.GlobalString("loglevel"))
	healthCheckHandler := healthCheckHandlerImpl{
		log: logger,
	}
	healthCheckHandler.inject(app)
	// Job handler
	jobHanler := jobHandlerImpl{
		jobRepo: repo.JobRepoImpl{
			DB: db,
		},
		log: logger,
	}
	jobHanler.inject(app)
	// HTTP ReqResIn service
	reqResInHanler := reqResInHandlerImpl{
		reqResService: service.NewReqResIn("https://reqres.in"),
		log:           logger,
	}
	reqResInHanler.inject(app)
	return app
}

func simpleReturnHandler(c iris.Context, result interface{}, err Error) {
	if err != nil {
		c.StatusCode(err.Status())
		c.JSON(iris.Map{
			"error": err.Error(),
		})
		return
	}
	c.JSON(result)
}
