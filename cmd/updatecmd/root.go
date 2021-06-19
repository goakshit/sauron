package updatecmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var updateCmd *cobra.Command

func GetUpdateCmd() *cobra.Command {
	return updateCmd
}

func init() {
	updateCmd = &cobra.Command{
		Use:   "update",
		Short: "Updates existing instances",
		Long:  `Allows us to update existing instances of merchants`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Update Cmd")
		},
	}
	// Initialises all the sub commands
	addSubCmd()
}

func addSubCmd() {
	updateCmd.AddCommand(
		getMerchantCmd(),
		getUserCmd(),
	)
}
