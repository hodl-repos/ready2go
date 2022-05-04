package interactive

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/hodl-repos/ready2go/internal/cli/helper"
	"github.com/urfave/cli/v2"
)

var menus = map[string]func(c *cli.Context) error{
	"webhook":           webhookMenu,
	"billType":          billTypeMenu,
	"country":           countryMenu,
	"coupon":            couponMenu,
	"couponCategory":    couponCategoryMenu,
	"currency":          currencyMenu,
	"customer":          customerMenu,
	"customerCategory":  customerCategoryMenu,
	"discount":          discountMenu,
	"discountGroup":     discountGroupMenu,
	"language":          languageMenu,
	"legalForm":         legalFormMenu,
	"order":             orderMenu,
	"paymentMethod":     paymentMethodMenu,
	"paymentMethodType": paymentMethodTypeMenu,
	"printJob":          printJobMenu,
	"printLog":          printLogMenu,
	"printer":           printerMenu,
	"productGroup":      productGroupMenu,
	"storno":            stornoMenu,
}

func menuSelectAndRun(title string, c *cli.Context, menu map[string]func(c *cli.Context) error) error {
	menuOptions := make([]string, 0)
	for k := range menu {
		menuOptions = append(menuOptions, k)
	}

	option := ""
	prompt := &survey.Select{
		Message: title,
		Options: menuOptions,
	}
	err := survey.AskOne(prompt, &option)

	if err != nil {
		helper.CheckTerminalInterrupt(err)
		return err
	}

	selMenu := menu[option]

	err = selMenu(c)

	if err != nil {
		helper.CheckTerminalInterrupt(err)
		return err
	}

	return nil
}
