package helper

import (
	"os"

	"github.com/AlecAivazis/survey/v2/terminal"
)

func CheckTerminalInterrupt(err error) {
	if err == terminal.InterruptErr {
		os.Exit(0)
	}
}
