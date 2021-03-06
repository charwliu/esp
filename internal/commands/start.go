package commands

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/sevlyar/go-daemon"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"

	"go.vixal.xyz/esp/internal/config"
	"go.vixal.xyz/esp/internal/service"
	"go.vixal.xyz/esp/pkg/fs"
	"go.vixal.xyz/esp/pkg/txt"
	"go.vixal.xyz/esp/server"
)

// StartCommand registers the start cli command.
var StartCommand = &cli.Command{
	Name:    "start",
	Aliases: []string{"up"},
	Usage:   "Starts web server",
	Flags:   startFlags,
	Action:  startAction,
}

var startFlags = []cli.Flag{
	&cli.BoolFlag{
		Name:    "detach-server",
		Aliases: []string{"d"},
		Usage:   "detach from the console (daemon mode)",
		EnvVars: []string{"GOSH_DETACH_SERVER"},
	},
	&cli.BoolFlag{
		Name:    "config",
		Aliases: []string{"c"},
		Usage:   "show config",
	},
}

// startAction start the web server and initializes the daemon
func startAction(ctx *cli.Context) error {
	conf := config.NewConfig(ctx)
	defer func() {
		conf.LoggerClose()
	}()
	service.SetConfig(conf)

	if ctx.IsSet("config") {
		fmt.Printf("NAME                  VALUE\n")
		fmt.Printf("detach-server         %t\n", conf.DetachServer())

		fmt.Printf("http-host             %s\n", conf.HttpHost())
		fmt.Printf("http-port             %d\n", conf.HttpPort())
		fmt.Printf("http-mode             %s\n", conf.HttpMode())

		return nil
	}

	if conf.HttpPort() < 1 || conf.HttpPort() > 65535 {
		zap.S().Fatal("server port must be a number between 1 and 65535")
	}

	// pass this context down the chain
	cctx, cancel := context.WithCancel(context.Background())

	if err := conf.Init(); err != nil {
		zap.S().Fatal(err)
	}

	// initialize the database
	conf.InitDB()

	// check if daemon is running, if not initialize the daemon
	dctx := new(daemon.Context)
	dctx.LogFileName = conf.LogFilename()
	dctx.PidFileName = conf.PIDFilename()
	dctx.Args = ctx.Args().Tail()

	if !daemon.WasReborn() && conf.DetachServer() {
		conf.Shutdown()
		cancel()

		if pid, ok := childAlreadyRunning(conf.PIDFilename()); ok {
			zap.S().Infof("daemon already running with process id %v\n", pid)
			return nil
		}

		child, err := dctx.Reborn()
		if err != nil {
			zap.S().Fatal(err)
		}

		if child != nil {
			if !fs.Overwrite(conf.PIDFilename(), []byte(strconv.Itoa(child.Pid))) {
				zap.S().Fatalf("failed writing process id to %s", txt.Quote(conf.PIDFilename()))
			}

			zap.S().Infof("daemon started with process id %v\n", child.Pid)
			return nil
		}
	}

	if conf.ReadOnly() {
		zap.S().Infof("start: read-only mode enabled")
	}

	// start web server
	go server.Start(cctx, conf)

	// set up proper shutdown of daemon and web server
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	//stop share & sync workers
	//workers.Stop()
	//auto.Stop()

	zap.S().Info("shutting down...")
	conf.Shutdown()
	cancel()
	err := dctx.Release()

	if err != nil {
		zap.S().Error(err)
	}

	time.Sleep(3 * time.Second)

	return nil
}
