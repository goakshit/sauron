package updatecmd

import (
	"context"
	"fmt"
	"strconv"

	"github.com/goakshit/sauron/internal/constants"
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

	var (
		err error
	)

	// Road block ahead, I mean validations

	if len(args) != 2 {
		fmt.Println(constants.UpdateMerchantInvalidParamsErr)
		return
	}

	perc, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		fmt.Println(constants.UpdateMerchantPercErr)
		return
	}

	if perc <= 0 {
		fmt.Println(constants.UpdateMerchantInvalidDiscountErr)
		return
	}

	merchantSVC := merchant.NewMerchantService(persistence.GetGormClient())
	err = merchantSVC.UpdateMerchantDiscount(context.Background(), args[0], perc)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Successfully updated the merchant '%s' discount to %f", args[0], perc)
}
