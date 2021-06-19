package reportcmd

import (
	"github.com/spf13/cobra"
)

var reportCmd *cobra.Command

func GetReportCmd() *cobra.Command {
	return reportCmd
}

func init() {
	reportCmd = &cobra.Command{
		Use:   "report",
		Short: "Generates reports",
		Long:  `Allows us to generate different reports`,
		// Run: func(cmd *cobra.Command, args []string) {},
	}
	// Initialises all the sub commands
	addSubCmd()
}

func addSubCmd() {
	reportCmd.AddCommand(
		getUsersAtCreditLimitCmd(),
		getUsersDuesCmd(),
		getUserDueCmd(),
		getMerchantDiscountCmd(),
	)
}
