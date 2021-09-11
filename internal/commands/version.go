package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"go.vixal.xyz/esp/internal/config"
)

// VersionCommand registers the version cli command.
var VersionCommand = &cli.Command{
	Name:   "version",
	Usage:  "Shows version information",
	Action: versionAction,
}

// versionAction prints the current version
func versionAction(ctx *cli.Context) error {
	conf := config.NewConfig(ctx)

	defer func() {
		conf.LoggerClose()
	}()

	fmt.Println(conf.Version())

	return nil
}
