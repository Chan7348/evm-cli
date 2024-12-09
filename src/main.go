package main

import (
	"fmt"

	"github.com/Chan7348/evm-cli/src/config"
	"github.com/Chan7348/evm-cli/src/sendTx"
	"github.com/Chan7348/evm-cli/src/show"
	"github.com/Chan7348/evm-cli/src/view"
	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:   "evm-cli",
		Short: "CLI tool for call and send Txs to evm chains.",

		// 执行命令前的检查
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			fmt.Println("Running evm-cli......")
			fmt.Printf("Using network: %s\n", config.Network)
		},
	}

	cmd.PersistentFlags().StringVar(&config.Network, "network", "", "Which blockchain to interact with")
	cmd.MarkPersistentFlagRequired("network")

	// load commands
	cmd.AddCommand(show.Init(), view.Init(), sendTx.Init())

	// Run
	cmd.Execute()
}
