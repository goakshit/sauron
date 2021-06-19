package newcmd

import (
	"fmt"

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
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("New Cmd")
		},
	}
	// Initialises all the sub commands
	addSubCmd()
}

func addSubCmd() {
	// Adds merchant command
	newCmd.AddCommand(getMerchantCmd())
}
