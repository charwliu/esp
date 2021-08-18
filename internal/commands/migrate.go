package commands

import (
	"context"
	"time"

	"github.com/urfave/cli/v2"

	"go.vixal.xyz/esp/internal/config"
)

// MigrateCommand registers the migrate cli command.
var MigrateCommand = &cli.Command{
	Name:   "migrate",
	Usage:  "Initializes the index database if needed",
	Action: migrateAction,
}

// migrateAction initializes and migrates the database.
func migrateAction(ctx *cli.Context) error {
	start := time.Now()

	conf := config.NewConfig(ctx)

	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := conf.Init(); err != nil {
		return err
	}

	log.Info("migrating database")

	conf.InitDB()

	elapsed := time.Since(start)

	log.Infof("database migration completed in %s", elapsed)

	conf.Shutdown()

	return nil

}