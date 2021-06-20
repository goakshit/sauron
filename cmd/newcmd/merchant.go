package newcmd

import (
	"context"
	"fmt"

	"github.com/goakshit/sauron/internal/constants"
	"github.com/goakshit/sauron/internal/persistence"
	"github.com/goakshit/sauron/internal/svc/merchant"
	"github.com/spf13/cobra"
)

func getMerchantCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "merchant",
		Short: "Add a new merchant",
		Long:  `This command allows you to add a new merchant`,
		Run: func(cmd *cobra.Command, args []string) {
			createMerchant(args)
		},
	}
}

// createMerchant - creates merchant from data passed
func createMerchant(args []string) {

	merchantRepo := merchant.NewRepository(persistence.GetGormClient())
	merchantSVC := merchant.NewMerchantService(merchantRepo)
	err := merchantSVC.CreateMerchant(context.Background(), args)
	if err != nil {
		switch err.Error() {
		case "ERROR: duplicate key value violates unique constraint \"merchant_pkey\" (SQLSTATE 23505)":
			fmt.Println(constants.CreateMerchantDuplicateIDErr)
		case "ERROR: duplicate key value violates unique constraint \"merchant_email_unique\" (SQLSTATE 23505)":
			fmt.Println(constants.CreateMerchantDuplicateEmailErr)
		default:
			fmt.Println(err.Error())
		}
		return
	}
	fmt.Println("Successfully added merchant with name: " + args[0])
}
