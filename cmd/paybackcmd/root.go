package paybackcmd

import (
	"context"
	"fmt"
	"strconv"

	"github.com/goakshit/sauron/internal/constants"
	"github.com/goakshit/sauron/internal/persistence"
	"github.com/goakshit/sauron/internal/svc/payback"
	"github.com/goakshit/sauron/internal/types"
	"github.com/spf13/cobra"
)

func GetPaybackCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "payback",
		Short: "Create payback for an user",
		Long:  `This command allows you to create payback for an user`,
		Run: func(cmd *cobra.Command, args []string) {
			createPayback(args)
		},
	}
}

// createPayback - creates payback for an user from data passed
func createPayback(args []string) {

	var (
		err            error
		paybackDetails types.PaybackDetails
	)

	// Road block ahead, I mean validations
	if len(args) != 2 {
		fmt.Println(constants.CreatePaybackInvalidParamsErr)
		return
	}

	paybackAmount, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		fmt.Println(constants.CreatePaybackInvalidAmountErr)
		return
	}

	if paybackAmount <= 0 {
		fmt.Println(constants.CreatePaybackInvalidAmountErr)
		return
	}

	paybackDetails.UserName = args[0]
	paybackDetails.Amount = paybackAmount

	paybackSVC := payback.NewPaybackService(persistence.GetGormClient())
	err = paybackSVC.CreatePayback(context.Background(), paybackDetails)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Successfully added new payback")
}
