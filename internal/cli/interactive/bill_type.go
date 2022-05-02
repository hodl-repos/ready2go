package interactive

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
	"github.com/hodl-repos/ready2go/internal/cli/config"
	"github.com/hodl-repos/ready2go/internal/cli/helper"
	"github.com/hodl-repos/ready2go/r2o"
	"github.com/urfave/cli/v2"
)

var billTypeMenus = map[string]func(c *cli.Context) error{
	"get all":   runBillTypesGetAll,
	"get by id": runBillTypesGet,
}

func billTypeMenu(c *cli.Context) error {
	return menuSelectAndRun("select function to run:", c, billTypeMenus)
}

func runBillTypesGetAll(c *cli.Context) error {
	accToken := config.AccountToken()
	clt := r2o.NewClient(&accToken, nil)

	data, err := clt.BillType.GetBillTypes(c.Context)

	if err != nil {
		msg := fmt.Sprintf("error while loading from r2o API: %v", err.Error())
		fmt.Println(msg)
		return errors.New(msg)
	}

	formattedBytes, _ := json.MarshalIndent(data, "", "    ")
	fmt.Printf("RESULT:\n\n%v\n\n", string(formattedBytes))

	return nil
}

func runBillTypesGet(c *cli.Context) error {

	var billTypeId int

	for {
		billTypeIdString := ""

		prompt := &survey.Input{
			Message: "enter billType-id",
		}

		err := survey.AskOne(prompt, &billTypeIdString)
		if err != nil {
			helper.CheckTerminalInterrupt(err)
			return err
		}

		val, err := strconv.ParseInt(billTypeIdString, 10, 64)

		if err != nil {
			fmt.Println("entered wrong type, number expected")
		} else {
			billTypeId = int(val)
			break
		}
	}

	accToken := config.AccountToken()
	clt := r2o.NewClient(&accToken, nil)

	data, err := clt.BillType.GetBillType(c.Context, &billTypeId)

	if err != nil {
		msg := fmt.Sprintf("error while loading from r2o API: %v", err.Error())
		fmt.Println(msg)
		return errors.New(msg)
	}

	formattedBytes, _ := json.MarshalIndent(data, "", "    ")
	fmt.Printf("RESULT:\n\n%v\n\n", string(formattedBytes))

	return nil
}
