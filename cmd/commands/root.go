package commands

import (
	"fmt"
	"os"

	"github.com/hodl-repos/ready2go/r2o"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "r2o",
	Short: "Ready2Go CLI is a wrapper for the public Ready2Order API",
	Long: `A wrapper for the public Ready2Order API, written in Go.
                Documentation from the API is available at https://app.swaggerhub.com/apis-docs/ready2ordergmbh/ready2order-api-production`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("RUN HELP OR DO SOMETHING")
	},
	TraverseChildren: true,
}

func Execute() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(accountCmd)
	rootCmd.AddCommand(authorizeCmd)
	rootCmd.AddCommand(webhookCmd)
	webhookCmd.AddCommand(webhookEventsCmd)

	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func parseJwtInitializeClient() (*r2o.Client, error) {
	prompt := promptui.Prompt{
		Label:    "Account API Token",
		Validate: validateJwtToken,
	}

	result, err := prompt.Run()

	if err != nil {
		return nil, err
	}

	client := r2o.NewClient(&result, nil)

	return client, nil
}

func validateJwtToken(input string) error {
	//VALIDATE JWT FORMAT
	return nil
}
