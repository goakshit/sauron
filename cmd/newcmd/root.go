package newcmd

import (
	"github.com/spf13/cobra"
)

var newCmd *cobra.Command

func GetNewCmd() *cobra.Command {
	return newCmd
}

func init() {
	newCmd = &cobra.Command{
		Use:   "new",
		Short: "Creates new instances",
		Long:  `Allows us to create new instance of users, transactions & merchants`,
	}
	// Initialises all the sub commands
	addSubCmd()
}

func addSubCmd() {
	// Adds merchant, user & transaction command
	newCmd.AddCommand(
		getUserCmd(),
		getMerchantCmd(),
		getTransactionCmd(),
	)
}
