package interactive

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hodl-repos/ready2go/internal/cli/config"
	"github.com/hodl-repos/ready2go/r2o"
	"github.com/urfave/cli/v2"
)

func printJobMenu(c *cli.Context) error {
	// return menuSelectAndRun("select function to run:", c, printJobMenus)
	return runPrintJobGetAll(c)
}

func runPrintJobGetAll(c *cli.Context) error {
	accToken := config.AccountToken()
	clt := r2o.NewClient(&accToken, nil)

	data, err := clt.PrintJob.GetPrintJobs(c.Context)

	if err != nil {
		msg := fmt.Sprintf("error while loading from r2o API: %v", err.Error())
		fmt.Println(msg)
		return errors.New(msg)
	}

	formattedBytes, _ := json.MarshalIndent(data, "", "    ")
	fmt.Printf("RESULT:\n\n%v\n\n", string(formattedBytes))

	return nil
}
