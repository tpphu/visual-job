package cmd

import (
	"github.com/tpphu/visual-job/model"

	"github.com/urfave/cli"
)

func migrateAction(appContext *cli.Context) {
	db := newDB(appContext)
	defer db.Close()
	db.AutoMigrate(&model.Job{}, &model.Product{})
}

// Migrate is a definition of cli.Command used to migrate schema to database
var Migrate = cli.Command{
	Name:  "migrate",
	Usage: "Migrate db",
	Action: func(appContext *cli.Context) error {
		migrateAction(appContext)
		return nil
	},
}
