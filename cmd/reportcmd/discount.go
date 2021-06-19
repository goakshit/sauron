package reportcmd

import (
	"context"
	"fmt"

	"github.com/goakshit/sauron/internal/constants"
	"github.com/goakshit/sauron/internal/persistence"
	"github.com/goakshit/sauron/internal/svc/report"
	"github.com/spf13/cobra"
)

func getMerchantDiscountCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "discount",
		Short: "Get discount offered from merchant",
		Long:  `This command allows you to get discounts from a merchant on all successful transactions`,
		Run: func(cmd *cobra.Command, args []string) {
			generateMerchantDiscountReport(args)
		},
	}
}

// generateMerchantDiscountReport - generates merchant discounts report
func generateMerchantDiscountReport(args []string) {

	if len(args) != 1 {
		fmt.Println(constants.ReportInvalidParamErr)
		return
	}

	reportSVC := report.NewReportService(persistence.GetGormClient())
	discount, err := reportSVC.GetMerchantDiscount(context.Background(), args[0])
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Merchant(%s) discount for all transactions: %.2f", args[0], discount)
}
