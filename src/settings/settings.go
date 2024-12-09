package settings

import (
	"fmt"
	"log"

	cobra "github.com/spf13/cobra"
)

var (
	Network string
)

func Init() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "evm-cli",
		Short: "CLI tool for query and send Tx from evm chain.",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			fmt.Println("Running evm-cli...")
			if Network == "" {
				log.Fatal("Error: Network flag is required")
			}
			fmt.Printf("Using network: %s\n", Network)
		},
	}
	cmd.PersistentFlags().StringVar(&Network, "network", "", "which blockchain.")

	show := &cobra.Command{
		Use:     "show",
		Short:   "show something.",
		Example: "evm-cli show",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("This is show command")
		},
	}

	view := &cobra.Command{
		Use:     "view",
		Short:   "Query the view functions of a Smart Contract",
		Example: "evm-cli view",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("This is view command")
		},
	}

	sendTx := &cobra.Command{
		Use:     "sendTx",
		Short:   "Send transaction",
		Example: "evm-cli sendTx",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("This is sendTx command")
		},
	}

	cmd.AddCommand(show, view, sendTx)
	return cmd
}
