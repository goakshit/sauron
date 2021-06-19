package newcmd

import (
	"context"
	"fmt"
	"strconv"

	"github.com/goakshit/sauron/internal/constants"
	"github.com/goakshit/sauron/internal/persistence"
	"github.com/goakshit/sauron/internal/svc/transaction"
	"github.com/goakshit/sauron/internal/types"
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

	var (
		err        error
		txnDetails types.TxnDetails
	)

	// Road block ahead, I mean validations

	if len(args) != 3 {
		fmt.Println(constants.CreateTxnInvalidParamsErr)
		return
	}

	txnAmount, err := strconv.ParseFloat(args[2], 64)
	if err != nil {
		fmt.Println(constants.CreateTxnInvalidAmountErr)
		return
	}

	if txnAmount <= 0 {
		fmt.Println(constants.CreateTxnInvalidAmountErr)
		return
	}

	txnDetails.UserName = args[0]
	txnDetails.MerchantName = args[1]
	txnDetails.Amount = txnAmount

	txnRepo := transaction.NewRepository(persistence.GetGormClient())
	txnSVC := transaction.NewTxnService(txnRepo)
	err = txnSVC.CreateTransaction(context.Background(), txnDetails)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Successfully added new transaction")
}
