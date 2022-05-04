package interactive

import (
	"fmt"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
	"github.com/hodl-repos/ready2go/internal/cli/helper"
)

func getNumberPrompt(title string) (*int, error) {
	for {
		numberString := ""

		prompt := &survey.Input{
			Message: title,
		}

		err := survey.AskOne(prompt, &numberString)
		if err != nil {
			helper.CheckTerminalInterrupt(err)
			return nil, err
		}

		val, err := strconv.ParseInt(numberString, 10, 64)

		if err != nil {
			fmt.Println("entered wrong type, number expected")
		} else {
			valueInt := int(val)
			return &valueInt, nil
		}
	}
}

func getNumberPromptOptional(title string) (*int, error) {
	for {
		numberString := ""

		prompt := &survey.Input{
			Message: title,
		}

		err := survey.AskOne(prompt, &numberString)
		if err != nil {
			helper.CheckTerminalInterrupt(err)
			return nil, err
		}

		if len(numberString) < 1 {
			return nil, nil
		}

		val, err := strconv.ParseInt(numberString, 10, 64)

		if err != nil {
			fmt.Println("entered wrong type, number expected")
		} else {
			valueInt := int(val)
			return &valueInt, nil
		}
	}
}
