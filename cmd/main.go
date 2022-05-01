package main

import (
	"log"
	"os"

	"github.com/hodl-repos/ready2go/internal/cli/authorize"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{}
	app.EnableBashCompletion = true
	app.UseShortOptionHandling = true

	// global flag
	app.Flags = []cli.Flag{}

	// root commands
	app.Commands = []*cli.Command{
		authorize.Command,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
