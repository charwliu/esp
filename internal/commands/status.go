package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"
	"github.com/valyala/fasthttp"

	"go.vixal.xyz/esp/internal/config"
)

// StatusCommand performs a server health check.
var StatusCommand = &cli.Command{
	Name:   "status",
	Usage:  "Performs a server health check",
	Action: statusAction,
}

// statusAction shows the server health status
func statusAction(ctx *cli.Context) error {
	conf := config.NewConfig(ctx)
	client := fasthttp.Client{}

	url := fmt.Sprintf("http://%s:%d/api/v1/status", conf.HttpHost(), conf.HttpPort())

	var status []byte

	statusCode, status, err := client.Get(status, url)

	if err != nil {
		return fmt.Errorf("can't connect to %s:%d", conf.HttpHost(), conf.HttpPort())
	} else if statusCode != 200 {
		return fmt.Errorf("server running at %s:%d, bad status %d\n", conf.HttpHost(), conf.HttpPort(), statusCode)
	}

	fmt.Println(string(status))

	return nil
}
