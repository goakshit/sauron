package updatecmd

import (
	"context"
	"fmt"

	"github.com/goakshit/sauron/internal/persistence"
	"github.com/goakshit/sauron/internal/svc/merchant"
	"github.com/spf13/cobra"
)

func getMerchantCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "merchant",
		Short: "Updates existing merchant",
		Long:  `This command allows you to update merchant discount`,
		Run: func(cmd *cobra.Command, args []string) {
			updateMerchantDiscount(args)
		},
	}
}

// updateMerchantDiscount - updates merchant discount from data passed
func updateMerchantDiscount(args []string) {

	merchantRepo := merchant.NewRepository(persistence.GetGormClient())
	merchantSVC := merchant.NewMerchantService(merchantRepo)
	err := merchantSVC.UpdateMerchantDiscount(context.Background(), args)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Successfully updated the merchant '%s' discount to %s", args[0], args[1])
}
