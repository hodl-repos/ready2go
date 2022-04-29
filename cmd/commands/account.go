package commands

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var accountCmd = &cobra.Command{
	Use:   "account",
	Short: "Loads the Account info for the current user",
	Long:  `Loads the Account info for the current user`,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := parseJwtInitializeClient()

		if err != nil {
			fmt.Printf("Cannot create r2o-client with given access token: %v", err.Error())
			return
		}

		account, err := client.Account.GetAccountInfo(cmd.Context())

		if err != nil {
			fmt.Printf("error while loading from r2o API: %v", err.Error())
			return
		}

		formattedBytes, _ := json.MarshalIndent(account, "", "    ")
		fmt.Printf("RESULT:\n\n%v\n\n", string(formattedBytes))
	},
}
