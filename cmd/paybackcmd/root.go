package paybackcmd

import (
	"context"
	"fmt"

	"github.com/goakshit/sauron/internal/persistence"
	"github.com/goakshit/sauron/internal/svc/payback"
	"github.com/goakshit/sauron/internal/svc/transaction"
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

	paybackRepo := payback.NewRepository(persistence.GetGormClient())
	txnRepo := transaction.NewRepository(persistence.GetGormClient())
	paybackSVC := payback.NewPaybackService(paybackRepo, txnRepo)
	err := paybackSVC.CreatePayback(context.Background(), args)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Successfully added payback")
}
