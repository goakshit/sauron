package newcmd

import (
	"context"
	"fmt"
	"strconv"

	"github.com/goakshit/sauron/internal/constants"
	"github.com/goakshit/sauron/internal/persistence"
	"github.com/goakshit/sauron/internal/svc/merchant"
	"github.com/goakshit/sauron/internal/types"
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

	var (
		err             error
		merchantDetails types.MerchantDetails
	)

	// Road block ahead, I mean validations

	if len(args) != 3 {
		fmt.Println(constants.CreateMerchantInvalidParamsErr)
		return
	}

	perc, err := strconv.ParseFloat(args[2], 64)
	if err != nil {
		fmt.Println(constants.CreateMerchantPercErr)
		return
	}

	if perc <= 0 || perc > 100 {
		fmt.Println(constants.UpdateMerchantInvalidDiscountErr)
		return
	}

	merchantDetails.Name = args[0]
	merchantDetails.Email = args[1]
	merchantDetails.Perc = perc

	merchantSVC := merchant.NewMerchantService(persistence.GetGormClient())
	err = merchantSVC.CreateMerchant(context.Background(), merchantDetails)
	if err != nil {
		switch err.Error() {
		case "ERROR: duplicate key value violates unique constraint \"merchant_pkey\" (SQLSTATE 23505)":
			fmt.Println(constants.CreateMerchantDuplicateIDErr)
		case "ERROR: duplicate key value violates unique constraint \"merchant_email_unique\" (SQLSTATE 23505)":
			fmt.Println(constants.CreateMerchantDuplicateEmailErr)
		}
		return
	}
	fmt.Println("Successfully added merchant with name: " + merchantDetails.Name)
}
