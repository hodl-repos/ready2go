package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Ready2Go CLI",
	Long:  `All software has versions. This is Ready2Go CLI's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Ready2Go CLI v0.1")
	},
}
