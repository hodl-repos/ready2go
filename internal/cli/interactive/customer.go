package interactive

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hodl-repos/ready2go/internal/cli/config"
	"github.com/hodl-repos/ready2go/internal/cli/helper"
	"github.com/hodl-repos/ready2go/r2o"
	"github.com/urfave/cli/v2"
)

var customerMenus = map[string]func(c *cli.Context) error{
	"get all":   runCustomerGetAll,
	"get by id": runCustomerGetById,
}

func customerMenu(c *cli.Context) error {
	return menuSelectAndRun("select function to run:", c, customerMenus)
}

func runCustomerGetAll(c *cli.Context) error {
	pagination, err := requestPaginationData()

	if err != nil {
		msg := fmt.Sprintf("error while requesting pagintion-data from user: %v", err.Error())
		fmt.Println(msg)
		return errors.New(msg)
	}

	accToken := config.AccountToken()
	clt := r2o.NewClient(&accToken, nil)

	data, err := clt.Customer.GetCustomers(c.Context, pagination)

	if err != nil {
		msg := fmt.Sprintf("error while loading from r2o API: %v", err.Error())
		fmt.Println(msg)
		return errors.New(msg)
	}

	formattedBytes, _ := json.MarshalIndent(data, "", "    ")
	fmt.Printf("RESULT:\n\n%v\n\n", string(formattedBytes))

	return nil
}

func runCustomerGetById(c *cli.Context) error {
	couponId, err := getNumberPrompt("enter customer-id")
	if err != nil {
		helper.CheckTerminalInterrupt(err)
		return err
	}

	if err != nil {
		msg := fmt.Sprintf("error while requesting pagintion-data from user: %v", err.Error())
		fmt.Println(msg)
		return errors.New(msg)
	}

	accToken := config.AccountToken()
	clt := r2o.NewClient(&accToken, nil)

	data, err := clt.Customer.GetCustomer(c.Context, couponId)

	if err != nil {
		msg := fmt.Sprintf("error while loading from r2o API: %v", err.Error())
		fmt.Println(msg)
		return errors.New(msg)
	}

	formattedBytes, _ := json.MarshalIndent(data, "", "    ")
	fmt.Printf("RESULT:\n\n%v\n\n", string(formattedBytes))

	return nil
}
