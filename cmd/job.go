package cmd

import (
	"github.com/tpphu/visual-job/worker"
	"github.com/urfave/cli"
)

// jobAction starts command and init DI
func jobAction(appContext *cli.Context) {
	db := newDB(appContext)
	defer db.Close()

	pool := worker.CreateWorkerPool(3, db)
	pool.Start()

}

// Job is a definition of cli.Command used to start http server
var Job = cli.Command{
	Name:  "job",
	Usage: "Run Job Application",
	Flags: []cli.Flag{},
	Action: func(appContext *cli.Context) error {
		jobAction(appContext)
		return nil
	},
}
