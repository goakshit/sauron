package newcmd

import (
	"context"
	"fmt"

	"github.com/goakshit/sauron/internal/persistence"
	"github.com/goakshit/sauron/internal/svc/transaction"
	"github.com/spf13/cobra"
)

func getTransactionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "txn",
		Short: "Add a new transaction",
		Long:  `This command allows you to add a new transaction`,
		Run: func(cmd *cobra.Command, args []string) {
			createTransaction(args)
		},
	}
}

// createTransaction - creates transaction from data passed
func createTransaction(args []string) {

	txnRepo := transaction.NewRepository(persistence.GetGormClient())
	txnSVC := transaction.NewTxnService(txnRepo)
	err := txnSVC.CreateTransaction(context.Background(), args)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Successfully added new transaction")
}
