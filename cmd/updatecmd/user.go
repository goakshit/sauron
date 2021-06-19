package updatecmd

import (
	"context"
	"fmt"

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

	userRepo := user.NewRepository(persistence.GetGormClient())
	userSVC := user.NewUserService(userRepo)

	err := userSVC.UpdateUserCreditLimit(context.Background(), args)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Successfully updated the users '%s' credit limit to %s", args[0], args[1])
}
