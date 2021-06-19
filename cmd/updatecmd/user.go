package updatecmd

import (
	"context"
	"fmt"
	"strconv"

	"github.com/goakshit/sauron/internal/constants"
	"github.com/goakshit/sauron/internal/persistence"
	"github.com/goakshit/sauron/internal/svc/user"
	"github.com/spf13/cobra"
)

func getUserCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "user",
		Short: "Updates existing user",
		Long:  `This command allows you to update user's credit limit`,
		Run: func(cmd *cobra.Command, args []string) {
			updateUserCreditLimit(args)
		},
	}
}

// updateUserCreditLimit - updates user credit limit from data passed
func updateUserCreditLimit(args []string) {

	var (
		err error
	)

	// Road block ahead, I mean validations

	if len(args) != 2 {
		fmt.Println(constants.UpdateUserInvalidParamsErr)
		return
	}

	creditLimit, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		fmt.Println(constants.UpdateUserInvalidCreditLimitErr)
		return
	}

	if creditLimit < 0 {
		fmt.Println(constants.UpdateUserInvalidCreditLimitErr)
		return
	}

	userSVC := user.NewUserService(persistence.GetGormClient())
	err = userSVC.UpdateUserCreditLimit(context.Background(), args[0], creditLimit)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Successfully updated the users '%s' credit limit to %.2f", args[0], creditLimit)
}
