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

var legalFormMenus = map[string]func(c *cli.Context) error{
	"get all":   runLegalFormGetAll,
	"get by id": runLegalFormGetById,
}

func legalFormMenu(c *cli.Context) error {
	return menuSelectAndRun("select function to run:", c, legalFormMenus)
}

func runLegalFormGetAll(c *cli.Context) error {
	accToken := config.AccountToken()
	clt := r2o.NewClient(&accToken, nil)

	data, err := clt.LegalForm.GetLegalForms(c.Context)

	if err != nil {
		msg := fmt.Sprintf("error while loading from r2o API: %v", err.Error())
		fmt.Println(msg)
		return errors.New(msg)
	}

	formattedBytes, _ := json.MarshalIndent(data, "", "    ")
	fmt.Printf("RESULT:\n\n%v\n\n", string(formattedBytes))

	return nil
}

func runLegalFormGetById(c *cli.Context) error {
	parsed_id, err := getNumberPrompt("enter legal-form-id")
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

	data, err := clt.LegalForm.GetLegalForm(c.Context, parsed_id)

	if err != nil {
		msg := fmt.Sprintf("error while loading from r2o API: %v", err.Error())
		fmt.Println(msg)
		return errors.New(msg)
	}

	formattedBytes, _ := json.MarshalIndent(data, "", "    ")
	fmt.Printf("RESULT:\n\n%v\n\n", string(formattedBytes))

	return nil
}
