package main

import (
	"os"

	"go-microservice/cmd/app"

	"github.com/urfave/cli/v2"
)

func main() {
	app.Init()
	defer app.ShutDown()

	cliApp := cli.NewApp()
	cliApp.Name = "Zalango BFF - app wrapper for Zalango "
	cliApp.Version = "1.0.0"
	cliApp.Usage = ""

	cliApp.Commands = cli.Commands{
		{
			Name:  "server",
			Usage: "Start server",
			Action: func(c *cli.Context) error {
				app.StartServer()
				return nil
			},
		},
	}

	if err := cliApp.Run(os.Args); err != nil {
		panic(err)
	}

}
