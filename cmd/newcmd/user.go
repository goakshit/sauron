package newcmd

import (
	"context"
	"fmt"
	"strconv"

	"github.com/goakshit/sauron/internal/constants"
	"github.com/goakshit/sauron/internal/persistence"
	"github.com/goakshit/sauron/internal/svc/user"
	"github.com/goakshit/sauron/internal/types"
	"github.com/spf13/cobra"
)

func getUserCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "user",
		Short: "Add a new user",
		Long:  `This command allows you to add a new user`,
		Run: func(cmd *cobra.Command, args []string) {
			createUser(args)
		},
	}
}

// createUser - creates user from data passed
func createUser(args []string) {

	var (
		err         error
		userDetails types.UserDetails
	)

	// Road block ahead, I mean validations

	if len(args) != 3 {
		fmt.Println(constants.CreateUserInvalidParamsErr)
		return
	}

	creditLimit, err := strconv.ParseFloat(args[2], 64)
	if err != nil {
		fmt.Println(constants.CreateUserInvalidCreditLimitErr)
		return
	}

	if creditLimit <= 0 {
		fmt.Println(constants.CreateUserInvalidCreditLimitErr)
		return
	}

	userDetails.Name = args[0]
	userDetails.Email = args[1]
	userDetails.CreditLimit = creditLimit

	userSVC := user.NewUserService(persistence.GetGormClient())
	err = userSVC.CreateUser(context.Background(), userDetails)
	if err != nil {
		switch err.Error() {
		case "ERROR: duplicate key value violates unique constraint \"user_pkey\" (SQLSTATE 23505)":
			fmt.Println(constants.CreateUserDuplicateIDErr)
		case "ERROR: duplicate key value violates unique constraint \"user_email_unique\" (SQLSTATE 23505)":
			fmt.Println(constants.CreateUserDuplicateEmailErr)
		}
		return
	}
	fmt.Println("Successfully added user with name: " + userDetails.Name)
}
