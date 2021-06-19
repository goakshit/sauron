package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func GetNewCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "new",
		Short: "A brief description of your command",
		Long:  `A longer description that spans multiple lines and likely contains examples and usage of using your command.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("this is a new command 1")
		},
	}
}
