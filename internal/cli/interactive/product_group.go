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

var productGroupMenus = map[string]func(c *cli.Context) error{
	"get all":   runProductGroupGetAll,
	"get by id": runProductGroupGetById,
}

func productGroupMenu(c *cli.Context) error {
	return menuSelectAndRun("select function to run:", c, productGroupMenus)
}

func runProductGroupGetAll(c *cli.Context) error {
	pagination, err := requestPaginationData()

	if err != nil {
		msg := fmt.Sprintf("error while requesting pagintion-data from user: %v", err.Error())
		fmt.Println(msg)
		return errors.New(msg)
	}

	accToken := config.AccountToken()
	clt := r2o.NewClient(&accToken, nil)

	data, err := clt.ProductGroup.GetProductGroups(c.Context, pagination)

	if err != nil {
		msg := fmt.Sprintf("error while loading from r2o API: %v", err.Error())
		fmt.Println(msg)
		return errors.New(msg)
	}

	formattedBytes, _ := json.MarshalIndent(data, "", "    ")
	fmt.Printf("RESULT:\n\n%v\n\n", string(formattedBytes))

	return nil
}

func runProductGroupGetById(c *cli.Context) error {
	id, err := getNumberPrompt("enter productGroup-id")
	if err != nil {
		helper.CheckTerminalInterrupt(err)
		return err
	}

	accToken := config.AccountToken()
	clt := r2o.NewClient(&accToken, nil)

	data, err := clt.ProductGroup.GetProductGroup(c.Context, id)

	if err != nil {
		msg := fmt.Sprintf("error while loading from r2o API: %v", err.Error())
		fmt.Println(msg)
		return errors.New(msg)
	}

	formattedBytes, _ := json.MarshalIndent(data, "", "    ")
	fmt.Printf("RESULT:\n\n%v\n\n", string(formattedBytes))

	return nil
}
