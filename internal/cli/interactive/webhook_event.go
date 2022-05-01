package interactive

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/hodl-repos/ready2go/internal/cli/config"
	"github.com/hodl-repos/ready2go/internal/cli/helper"
	"github.com/hodl-repos/ready2go/r2o"
	"github.com/urfave/cli/v2"
)

func runWebhookGetEvents(c *cli.Context) error {
	accToken := config.AccountToken()
	clt := r2o.NewClient(&accToken, nil)

	data, err := clt.Webhook.FindEvents(c.Context)

	if err != nil {
		msg := fmt.Sprintf("error while loading from r2o API: %v", err.Error())
		fmt.Println(msg)
		return errors.New(msg)
	}

	formattedBytes, _ := json.MarshalIndent(data, "", "    ")
	fmt.Printf("RESULT:\n\n%v\n\n", string(formattedBytes))

	return nil
}

func webhookEventSelect() (*r2o.WebhookEvent, error) {
	menuOptions := make([]string, 0)

	menuOptions = append(menuOptions, string(r2o.WebhookEventProductCreated))
	menuOptions = append(menuOptions, string(r2o.WebhookEventProductUpdated))
	menuOptions = append(menuOptions, string(r2o.WebhookEventProductDeleted))
	menuOptions = append(menuOptions, string(r2o.WebhookEventProductGroupCreated))
	menuOptions = append(menuOptions, string(r2o.WebhookEventProductGroupUpdated))
	menuOptions = append(menuOptions, string(r2o.WebhookEventProductGroupDeleted))
	menuOptions = append(menuOptions, string(r2o.WebhookEventInvoiceCreated))
	menuOptions = append(menuOptions, string(r2o.WebhookEventOrderItemCreated))

	option := ""
	prompt := &survey.Select{
		Message: "select event:",
		Options: menuOptions,
	}
	err := survey.AskOne(prompt, &option)

	if err != nil {
		helper.CheckTerminalInterrupt(err)
		return nil, err
	}

	optionType := r2o.WebhookEvent(option)

	return &optionType, nil
}

func runWebhookAddSubscriptionEvent(c *cli.Context) error {
	selectedEvent, err := webhookEventSelect()
	if err != nil {
		return err
	}

	accToken := config.AccountToken()
	clt := r2o.NewClient(&accToken, nil)

	data, err := clt.Webhook.AddEvent(c.Context, selectedEvent)

	if err != nil {
		msg := fmt.Sprintf("error while loading from r2o API: %v", err.Error())
		fmt.Println(msg)
		return errors.New(msg)
	}

	formattedBytes, _ := json.MarshalIndent(data, "", "    ")
	fmt.Printf("RESULT:\n\n%v\n\n", string(formattedBytes))

	return nil
}

func runWebhookRemoveSubscriptionEvent(c *cli.Context) error {
	selectedEvent, err := webhookEventSelect()
	if err != nil {
		return err
	}

	accToken := config.AccountToken()
	clt := r2o.NewClient(&accToken, nil)

	data, err := clt.Webhook.RemoveEvent(c.Context, selectedEvent)

	if err != nil {
		msg := fmt.Sprintf("error while loading from r2o API: %v", err.Error())
		fmt.Println(msg)
		return errors.New(msg)
	}

	formattedBytes, _ := json.MarshalIndent(data, "", "    ")
	fmt.Printf("RESULT:\n\n%v\n\n", string(formattedBytes))

	return nil
}
