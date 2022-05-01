package interactive

import (
	"github.com/hodl-repos/ready2go/internal/cli/config"
	"github.com/urfave/cli/v2"
)

var Command = &cli.Command{
	Name:   "interactive",
	Usage:  "get interactive with ready2order API",
	Action: authorizeAndRun,
}

func authorizeAndRun(c *cli.Context) error {
	accToken, err := tokenPrompt("account-token: ")
	if err != nil {
		return err
	}

	config.SetAccountToken(accToken)

	for {
		err := menuSelectAndRun("Which client do you want to use?", c, menus)

		if err != nil {
			return err
		}
	}
}
