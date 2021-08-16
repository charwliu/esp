package main

import (
	"os"

	"github.com/urfave/cli/v2"

	"go.vixal.xyz/esp/internal/commands"
	"go.vixal.xyz/esp/internal/config"
	"go.vixal.xyz/esp/internal/event"

	_ "go.vixal.xyz/esp/openapi"
)

var version = "development"
var log = event.Log

// @title Gosh ESP API
// @version 1.0
// @description This is Gosh ESP REST-ish API Docs.
// @contact.name API Support
// @contact.email charwliu@ngboss.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @BasePath /api/v1
// @host localhost:8080
// @query.collection.format multi
func main() {
	app := &cli.App{
		Name:                 "Esp",
		Usage:                "E-Service platform",
		Version:              version,
		Copyright:            "© 2021 Charlie W. Liu <charwliu@gmail.com>",
		EnableBashCompletion: true,
		Flags:                config.GlobalFlags,
	}

	app.Commands = []*cli.Command{
		commands.StartCommand,
		commands.StopCommand,
		commands.ConfigCommand,
		commands.StatusCommand,
		commands.PasswdCommand,
		commands.MigrateCommand,
		commands.VersionCommand,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Error(err.Error())
	}
}
