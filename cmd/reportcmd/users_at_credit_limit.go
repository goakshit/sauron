package reportcmd

import (
	"context"
	"fmt"

	"github.com/goakshit/sauron/internal/constants"
	"github.com/goakshit/sauron/internal/persistence"
	"github.com/goakshit/sauron/internal/svc/report"
	"github.com/spf13/cobra"
)

func getUsersAtCreditLimitCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "users-at-credit-limit",
		Short: "Gets all users that have reached their credit limit",
		Long:  `This command allows you to gets all users that have reached their credit limit`,
		Run: func(cmd *cobra.Command, args []string) {
			generateUsersAtCreditLimitReport(args)
		},
	}
}

// generateUsersAtCreditLimitReport - generates users at credit limit report
func generateUsersAtCreditLimitReport(args []string) {

	reportSVC := report.NewReportService(persistence.GetGormClient())
	users, err := reportSVC.GetUsersAtCreditLimit(context.Background())
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if len(users) == 0 {
		fmt.Println(constants.ReportUACLNoRecordsErr)
		return
	}
	fmt.Println("Users at credit limit are:")
	for _, u := range users {
		fmt.Println(u)
	}
}
