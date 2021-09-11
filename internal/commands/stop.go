package commands

import (
	"syscall"

	"github.com/sevlyar/go-daemon"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"

	"go.vixal.xyz/esp/internal/config"
	"go.vixal.xyz/esp/pkg/txt"
)

// StopCommand registers the stop cli command.
var StopCommand = &cli.Command{
	Name:    "stop",
	Aliases: []string{"down"},
	Usage:   "Stops web server (only in daemon mode)",
	Action:  stopAction,
}

// stopAction stops the daemon if it is running.
func stopAction(ctx *cli.Context) error {
	conf := config.NewConfig(ctx)

	defer func() {
		conf.LoggerClose()
	}()

	zap.S().Infof("looking for pid in %s", txt.Quote(conf.PIDFilename()))

	dcxt := new(daemon.Context)
	dcxt.PidFileName = conf.PIDFilename()
	child, err := dcxt.Search()

	if err != nil {
		zap.S().Fatal(err)
	}

	err = child.Signal(syscall.SIGTERM)

	if err != nil {
		zap.S().Fatal(err)
	}

	st, err := child.Wait()

	if err != nil {
		zap.S().Info("daemon exited successfully")
		return nil
	}

	zap.S().Infof("daemon[%v] exited[%v]? successfully[%v]?\n", st.Pid(), st.Exited(), st.Success())

	return nil
}
