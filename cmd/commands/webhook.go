package commands

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var webhookCmd = &cobra.Command{
	Use:   "webhook",
	Short: "Finds the webhook from the current user",
	Long:  `Finds the webhook from the current user`,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := parseJwtInitializeClient()

		if err != nil {
			fmt.Printf("Cannot create r2o-client with given access token: %v", err.Error())
			return
		}

		webhook, err := client.Webhook.GetUrl(cmd.Context())

		if err != nil {
			fmt.Printf("error while loading from r2o API: %v", err.Error())
			return
		}

		formattedBytes, _ := json.MarshalIndent(webhook, "", "    ")
		fmt.Printf("RESULT:\n\n%v\n\n", string(formattedBytes))
	},
}

var webhookEventsCmd = &cobra.Command{
	Use:   "events",
	Short: "Finds the webhook-events from the current user",
	Long:  `Finds the webhook-events from the current user`,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := parseJwtInitializeClient()

		if err != nil {
			fmt.Printf("Cannot create r2o-client with given access token: %v", err.Error())
			return
		}

		events, err := client.Webhook.FindEvents(cmd.Context())

		if err != nil {
			fmt.Printf("error while loading from r2o API: %v", err.Error())
			return
		}

		formattedBytes, _ := json.MarshalIndent(events, "", "    ")
		fmt.Printf("RESULT:\n\n%v\n\n", string(formattedBytes))
	},
}
