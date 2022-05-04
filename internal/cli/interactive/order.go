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

func orderMenu(c *cli.Context) error {
	return runOrderGet(c)
	// return menuSelectAndRun("select function to run:", c, discountMenus)
}

func runOrderGet(c *cli.Context) error {
	pagination, err := requestPaginationData()
	if err != nil {
		helper.CheckTerminalInterrupt(err)
		return err
	}

	//SELECT GROUP-BYs
	groupByModesStrings := make([]string, 0)
	prompt := &survey.MultiSelect{
		Message: "Select GroupBy:",
		Options: []string{
			string(r2o.GroupByTypeTable),
			string(r2o.GroupByTypeProduct),
		},
	}
	err = survey.AskOne(prompt, &groupByModesStrings)

	if err != nil {
		helper.CheckTerminalInterrupt(err)
		return err
	}

	groupByModes := make([]r2o.GroupByType, 0)

	for _, modeString := range groupByModesStrings {
		if len(modeString) > 0 {
			groupByModes = append(groupByModes, r2o.GroupByType(modeString))
		}
	}

	//SELECT IF TRAININGSMODE
	trainingModeString := ""
	promptT := &survey.Select{
		Message: "Trainingsmode:",
		Options: []string{
			"YES",
			"NO",
		},
	}
	err = survey.AskOne(promptT, &trainingModeString)

	if err != nil {
		helper.CheckTerminalInterrupt(err)
		return err
	}

	trainingMode := false

	if len(trainingModeString) > 0 {
		if trainingModeString == "YES" {
			trainingMode = true
		}
	}

	//request table-id
	table_id, err := getNumberPromptOptional("Enter Table-ID:")
	if err != nil {
		helper.CheckTerminalInterrupt(err)
		return err
	}

	//request product_id
	product_id, err := getNumberPromptOptional("Enter Product-ID:")
	if err != nil {
		helper.CheckTerminalInterrupt(err)
		return err
	}

	request := r2o.OrderRequest{
		Pagination:    *pagination,
		TableId:       table_id,
		ProductId:     product_id,
		GroupBy:       &groupByModes,
		TrainingsMode: &trainingMode,
	}

	accToken := config.AccountToken()
	clt := r2o.NewClient(&accToken, nil)

	data, err := clt.Order.GetOrders(c.Context, &request)

	if err != nil {
		msg := fmt.Sprintf("error while loading from r2o API: %v", err.Error())
		fmt.Println(msg)
		return errors.New(msg)
	}

	formattedBytes, _ := json.MarshalIndent(data, "", "    ")
	fmt.Printf("RESULT:\n\n%v\n\n", string(formattedBytes))

	return nil
}
