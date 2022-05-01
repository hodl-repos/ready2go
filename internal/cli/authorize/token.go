package authorize

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/hodl-repos/ready2go/internal/cli/helper"
)

func tokenPrompt(promptMsg string) (string, error) {
	token := ""
	prompt := &survey.Input{
		Message: promptMsg,
	}

	err := survey.AskOne(prompt, &token)
	if err != nil {
		helper.CheckTerminalInterrupt(err)
		return "", err
	}

	return token, nil
}
