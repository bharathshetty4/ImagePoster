package cli

import (
	"fmt"

	cli "github.com/urfave/cli/v2"
)

func AddCommandStartServer() *cli.Command {
	return &cli.Command{
		Name:    "server",
		Aliases: []string{"s"},
		Usage:   "Start a server to do RESTful calls",
		Description: "Command to enable RESTful API calls which can be used to upload images, profile update etc." +
			"\n   Please check the api list available at <swagger_placeholder>",
		//TODO: BKS: insert seagger url here.
		UsageText: "image-poster server [command options]",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "port",
				DefaultText: "3000",
				Aliases:     []string{"p"},
				Usage:       "port on which the server should run",
			},
		},
		Action: startServer,
	}
}

func startServer(c *cli.Context) error {
	fmt.Printf("starting the server")
	return nil
}
