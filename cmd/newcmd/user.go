package newcmd

import (
	"context"
	"fmt"

	"github.com/goakshit/sauron/internal/constants"
	"github.com/goakshit/sauron/internal/persistence"
	"github.com/goakshit/sauron/internal/svc/user"
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

	userRepo := user.NewRepository(persistence.GetGormClient())
	userSVC := user.NewUserService(userRepo)
	err := userSVC.CreateUser(context.Background(), args)
	if err != nil {
		switch err.Error() {
		case "ERROR: duplicate key value violates unique constraint \"user_pkey\" (SQLSTATE 23505)":
			fmt.Println(constants.CreateUserDuplicateIDErr)
		case "ERROR: duplicate key value violates unique constraint \"user_email_unique\" (SQLSTATE 23505)":
			fmt.Println(constants.CreateUserDuplicateEmailErr)
		}
		return
	}
	fmt.Println("Successfully added user with name: " + args[0])
}
