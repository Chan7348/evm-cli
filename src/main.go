package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "evm-cli",
	Short: "A cli tool for interaction with EVM chains",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		rpc := os.Getenv("BASE_RPC")
		privateKey := os.Getenv("near1")
		network := args[0]

		if network == "" {
			fmt.Println("Error: Missing network")
			os.Exit(1)
		}

		fmt.Println("Environment variables loaded successfully!")

		fmt.Println("rpc:", rpc)
		fmt.Println("key:", privateKey)
		fmt.Println("network:", network)
	},
}

func main() {

	rootCmd.AddCommand(&cobra.Command{
		Use:   "run",
		Short: "Run the Script",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Running........")
		},
	})

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
