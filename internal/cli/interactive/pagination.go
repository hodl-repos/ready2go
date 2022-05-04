package interactive

import (
	"fmt"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
	"github.com/hodl-repos/ready2go/internal/cli/helper"
	"github.com/hodl-repos/ready2go/r2o"
)

func requestPaginationData() (*r2o.Pagination, error) {
	var page int
	var limit int

	for {
		pageString := ""

		prompt := &survey.Input{
			Message: "enter pagination page",
		}

		err := survey.AskOne(prompt, &pageString)
		if err != nil {
			helper.CheckTerminalInterrupt(err)
			return nil, err
		}

		val, err := strconv.ParseInt(pageString, 10, 64)

		if err != nil {
			fmt.Println("entered wrong type, number expected")
		} else {
			page = int(val)
			break
		}
	}

	for {
		limitString := ""

		prompt := &survey.Input{
			Message: "enter pagination limit",
		}

		err := survey.AskOne(prompt, &limitString)
		if err != nil {
			helper.CheckTerminalInterrupt(err)
			return nil, err
		}

		val, err := strconv.ParseInt(limitString, 10, 64)

		if err != nil {
			fmt.Println("entered wrong type, number expected")
		} else {
			limit = int(val)
			break
		}
	}

	return &r2o.Pagination{
		Page:  &page,
		Limit: &limit,
	}, nil
}
