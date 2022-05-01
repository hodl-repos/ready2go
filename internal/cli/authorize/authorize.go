package authorize

import (
	"fmt"

	"github.com/hodl-repos/ready2go/internal/cli/config"
	"github.com/hodl-repos/ready2go/r2o/authorization"
	"github.com/urfave/cli/v2"
)

var Command = &cli.Command{
	Name:   "authorize",
	Usage:  "Authorize with ready2order API",
	Action: authorize,
}

func authorize(c *cli.Context) error {
	devToken, err := tokenPrompt("developer-token: ")
	if err != nil {
		return err
	}

	fmt.Println("Starting Auth-Workflow")
	uri, err := authorization.GetAccountAccessToken(devToken)
	if err != nil {
		fmt.Printf("Cannot get URI from r2o: %v\n", err)
	}

	fmt.Printf("VISIT:\n\t%v\n\n", *uri)

	accToken, err := tokenPrompt("account-token: ")
	if err != nil {
		return err
	}

	config.SetAccountToken(accToken)

	return nil
}
