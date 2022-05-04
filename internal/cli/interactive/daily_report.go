package interactive

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hodl-repos/ready2go/internal/cli/config"
	"github.com/hodl-repos/ready2go/r2o"
	"github.com/urfave/cli/v2"
)

var dailyReportMenus = map[string]func(c *cli.Context) error{
	"open day":   runDailyReportOpenDay,
	"close day":  runDailyReportCloseDay,
	"get status": runDailyReportGetStatus,
}

func runDailyReportMenu(c *cli.Context) error {
	return menuSelectAndRun("select function to run:", c, dailyReportMenus)
}

func runDailyReportOpenDay(c *cli.Context) error {
	accToken := config.AccountToken()
	clt := r2o.NewClient(&accToken, nil)

	data, err := clt.DailyReport.OpenDay(c.Context)

	if err != nil {
		msg := fmt.Sprintf("error while loading from r2o API: %v", err.Error())
		fmt.Println(msg)
		return errors.New(msg)
	}

	formattedBytes, _ := json.MarshalIndent(data, "", "    ")
	fmt.Printf("RESULT:\n\n%v\n\n", string(formattedBytes))

	return nil
}

func runDailyReportCloseDay(c *cli.Context) error {
	accToken := config.AccountToken()
	clt := r2o.NewClient(&accToken, nil)

	data, err := clt.DailyReport.CloseDay(c.Context)

	if err != nil {
		msg := fmt.Sprintf("error while loading from r2o API: %v", err.Error())
		fmt.Println(msg)
		return errors.New(msg)
	}

	formattedBytes, _ := json.MarshalIndent(data, "", "    ")
	fmt.Printf("RESULT:\n\n%v\n\n", string(formattedBytes))

	return nil
}

func runDailyReportGetStatus(c *cli.Context) error {
	accToken := config.AccountToken()
	clt := r2o.NewClient(&accToken, nil)

	data, err := clt.DailyReport.GetStatus(c.Context)

	if err != nil {
		msg := fmt.Sprintf("error while loading from r2o API: %v", err.Error())
		fmt.Println(msg)
		return errors.New(msg)
	}

	formattedBytes, _ := json.MarshalIndent(data, "", "    ")
	fmt.Printf("RESULT:\n\n%v\n\n", string(formattedBytes))

	return nil
}
