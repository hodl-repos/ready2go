package interactive

import (
	"fmt"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
	"github.com/hodl-repos/ready2go/internal/cli/helper"
	"github.com/hodl-repos/ready2go/r2o"
)

func requestPaginationData() (*r2o.Pagination, error) {
	var offset int
	var limit int

	for {
		offsetString := ""

		prompt := &survey.Input{
			Message: "enter pagination offset",
		}

		err := survey.AskOne(prompt, &offsetString)
		if err != nil {
			helper.CheckTerminalInterrupt(err)
			return nil, err
		}

		val, err := strconv.ParseInt(offsetString, 10, 64)

		if err != nil {
			fmt.Println("entered wrong type, number expected")
		} else {
			offset = int(val)
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
		Offset: &offset,
		Limit:  &limit,
	}, nil
}
