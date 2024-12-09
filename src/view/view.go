package view

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/Chan7348/evm-cli/src/config"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/sha3"
)

func Init() *cobra.Command {
	return &cobra.Command{
		Use:     "view",
		Short:   "call a view functions of a Smart Contract",
		Example: "evm-cli view",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("This is view command")
			var weth common.Address

			if config.Network == "base" {
				weth = common.HexToAddress("0x4200000000000000000000000000000000000006")
			} else if config.Network == "ethereum" {
				weth = common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")
			}

			CallMethod(weth, "totalSupply()", nil)
		},
	}
}

func CallMethod(addr common.Address, method string, parameters []string) {
	if config.Network == "" {
		log.Fatal("Settings not done.")
	}
	rpc := os.Getenv(fmt.Sprintf("%s_RPC", strings.ToUpper(config.Network)))
	if rpc == "" {
		log.Fatalf("RPC for %s not found", config.Network)
	}

	client, err := ethclient.Dial(rpc)
	if err != nil {
		log.Fatalf("Error connecting the RPC: %v", err)
	}

	selector := getMethodSelector(method)

	callMsg := ethereum.CallMsg{
		To:   &addr,
		Data: append(selector, []byte(strings.Join(parameters, ""))...),
	}

	result, err := client.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		log.Fatalf("contract call failed: %s", err)
	}

	output := new(big.Int).SetBytes(result)
	fmt.Printf("Result of %s: %s\n", method, output)
}

func getMethodSelector(method string) []byte {
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write([]byte(method))
	return hasher.Sum(nil)[:4]
}
