package interactive

import (
	"github.com/urfave/cli/v2"
)

var webhookMenus = map[string]func(c *cli.Context) error{
	"get url":                    runWebhookGetUrl,
	"set url":                    runWebhookSetUrl,
	"get subscripted events":     runWebhookGetEvents,
	"add new subscription event": runWebhookAddSubscriptionEvent,
	"remove subscription event":  runWebhookRemoveSubscriptionEvent,
}

func webhookMenu(c *cli.Context) error {
	return menuSelectAndRun("select function to run:", c, webhookMenus)
}
