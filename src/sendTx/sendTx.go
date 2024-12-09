package sendTx

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Init() *cobra.Command {
	return &cobra.Command{
		Use:     "sendTx",
		Short:   "Send transaction",
		Example: "evm-cli sendTx",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("This is sendTx command")
		},
	}
}
