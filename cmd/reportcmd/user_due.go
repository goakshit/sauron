package reportcmd

import (
	"context"
	"fmt"

	"github.com/goakshit/sauron/internal/constants"
	"github.com/goakshit/sauron/internal/persistence"
	"github.com/goakshit/sauron/internal/svc/report"
	"github.com/spf13/cobra"
)

func getUserDueCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "dues",
		Short: "Get user's pending dues",
		Long:  `This command allows you to get user's pending dues if any`,
		Run: func(cmd *cobra.Command, args []string) {
			generateUserPendingDueReport(args)
		},
	}
}

// generateUserPendingDueReport - generates user's pending due report
func generateUserPendingDueReport(args []string) {

	if len(args) != 1 {
		fmt.Println(constants.ReportInvalidParamErr)
		return
	}

	reportSVC := report.NewReportService(persistence.GetGormClient())
	dueAmount, err := reportSVC.GetUserDues(context.Background(), args[0])
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Pending due amount for user '%s': %.2f", args[0], dueAmount)
}
