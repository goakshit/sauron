package reportcmd

import (
	"context"
	"fmt"

	"github.com/goakshit/sauron/internal/constants"
	"github.com/goakshit/sauron/internal/persistence"
	"github.com/goakshit/sauron/internal/svc/report"
	"github.com/spf13/cobra"
)

func getUsersDuesCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "total-dues",
		Short: "Gets all users that have pending dues",
		Long:  `This command allows you to gets all users that have pending dues`,
		Run: func(cmd *cobra.Command, args []string) {
			generateUsersDuesReport(args)
		},
	}
}

// generateUsersDuesReport - generates users with pending dues report
func generateUsersDuesReport(args []string) {

	reportRepo := report.NewRepository(persistence.GetGormClient())
	users, err := reportRepo.GetTotalDues(context.Background())
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if len(users) == 0 {
		fmt.Println(constants.ReportUserDuesNoRecordsErr)
		return
	}
	fmt.Println("Users with pending dues are:")
	var totalDues float64
	for _, u := range users {
		totalDues += u.DueAmount
		fmt.Printf("%s: %.2f", u.Name, u.DueAmount)
		fmt.Println()
	}
	fmt.Printf("Total Dues: %.2f", totalDues)
}
