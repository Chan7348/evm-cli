package view

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/Chan7348/evm-cli/src/settings"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
)

func CallMethod(addr common.Address, method string, parameters []string) {
	if settings.Network == "" {
		log.Fatal("Settings not done.")
	}
	rpc := os.Getenv(fmt.Sprintf("%s_RPC", strings.ToUpper(settings.Network)))
	if rpc == "" {
		log.Fatalf("RPC for %s not found", settings.Network)
	}

	client, err := ethclient.Dial(rpc)
	if err != nil {
		log.Fatalf("Error connecting the RPC: %v", err)
	}

	// blockNumber, err := client.BlockNumber(context.Background())
	// if err != nil {
	// 	log.Fatalf("Error getting the block number: %v", err)
	// }
	// log.Println("blockNumber:", blockNumber)

	result, err := callMethod(client, addr, method, parameters)
	if err != nil {
		log.Fatalf("Error calling contract method: %v", err)
	}

	fmt.Printf("Result of %s: %s\n", method, result)
}

func callMethod(client *ethclient.Client, contract common.Address, method string, parameters []string) (any, error) {
	selector := getMethodSelector(method)

	data := append(selector, []byte(strings.Join(parameters, ""))...)

	callMsg := ethereum.CallMsg{
		To:   &contract,
		Data: data,
	}

	result, err := client.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		return nil, fmt.Errorf("contract call failed: %w", err)
	}

	output := new(big.Int)
	output.SetBytes(result)

	return output, nil
}

func getMethodSelector(method string) []byte {
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write([]byte(method))
	return hasher.Sum(nil)[:4]
}
