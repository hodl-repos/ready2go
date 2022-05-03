package interactive

import (
	"fmt"

	"github.com/hodl-repos/ready2go/internal/cli/config"
	"github.com/urfave/cli/v2"
)

var Command = &cli.Command{
	Name:  "interactive",
	Usage: "get interactive with ready2order API",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "token",
			Aliases: []string{"t"},
			Usage:   "used value for the account-token",
		},
	},
	Action: authorizeAndRun,
}

func authorizeAndRun(c *cli.Context) error {
	accToken := ""
	if len(c.String("token")) > 0 {
		accToken = c.String("token")
	} else {
		token, err := tokenPrompt("account-token: ")
		if err != nil {
			return err
		}

		accToken = token
	}

	config.SetAccountToken(accToken)

	for {
		err := menuSelectAndRun("Which client do you want to use?", c, menus)

		if err != nil {
			return err
		}

		fmt.Println("--- starting over with same account token")
	}
}
