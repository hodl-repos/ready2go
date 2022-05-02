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

var countryMenus = map[string]func(c *cli.Context) error{
	"get all":   runCountryGetAll,
	"get by id": runCountryGet,
}

func countryMenu(c *cli.Context) error {
	return menuSelectAndRun("select function to run:", c, countryMenus)
}

func runCountryGetAll(c *cli.Context) error {
	accToken := config.AccountToken()
	clt := r2o.NewClient(&accToken, nil)

	data, err := clt.Country.GetCountries(c.Context)

	if err != nil {
		msg := fmt.Sprintf("error while loading from r2o API: %v", err.Error())
		fmt.Println(msg)
		return errors.New(msg)
	}

	formattedBytes, _ := json.MarshalIndent(data, "", "    ")
	fmt.Printf("RESULT:\n\n%v\n\n", string(formattedBytes))

	return nil
}

func runCountryGet(c *cli.Context) error {

	var countryId int

	for {
		countryIdString := ""

		prompt := &survey.Input{
			Message: "enter country-id",
		}

		err := survey.AskOne(prompt, &countryIdString)
		if err != nil {
			helper.CheckTerminalInterrupt(err)
			return err
		}

		val, err := strconv.ParseInt(countryIdString, 10, 64)

		if err != nil {
			fmt.Println("entered wrong type, number expected")
		} else {
			countryId = int(val)
			break
		}
	}

	accToken := config.AccountToken()
	clt := r2o.NewClient(&accToken, nil)

	data, err := clt.Country.GetCountry(c.Context, &countryId)

	if err != nil {
		msg := fmt.Sprintf("error while loading from r2o API: %v", err.Error())
		fmt.Println(msg)
		return errors.New(msg)
	}

	formattedBytes, _ := json.MarshalIndent(data, "", "    ")
	fmt.Printf("RESULT:\n\n%v\n\n", string(formattedBytes))

	return nil
}
