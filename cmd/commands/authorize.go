package commands

import (
	"fmt"

	"github.com/hodl-repos/ready2go/r2o/authorization"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var authorizeCmd = &cobra.Command{
	Use:   "authorize",
	Short: "Starts the authorization flow to link the app with an account",
	Long:  `Uses the Developer Token to create a new Session and forwards the user to the Auth-Website`,
	Run: func(cmd *cobra.Command, args []string) {
		prompt := promptui.Prompt{
			Label:    "Developer API Token",
			Validate: validateJwtToken,
		}

		result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		fmt.Println("Starting Auth-Workflow")

		uri, err := authorization.GetAccountAccessToken(result)

		if err != nil {
			fmt.Printf("Cannot get URI from r2o: %v\n", err)
			return
		}

		fmt.Printf("VISIT:\n\t%v\n\n", *uri)

		fmt.Println("You will be redirected to a small page showing the URL Query-Params containing the wanted Account API Token.")
	},
}
