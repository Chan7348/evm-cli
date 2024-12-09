package show

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Init() *cobra.Command {
	return &cobra.Command{
		Use:     "show",
		Short:   "show something.",
		Example: "evm-cli show",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("This is show Command")
		},
	}
}
