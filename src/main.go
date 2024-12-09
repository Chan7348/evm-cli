package main

import (
	settings "github.com/Chan7348/evm-cli/src/settings"
	"github.com/Chan7348/evm-cli/src/view"
	common "github.com/ethereum/go-ethereum/common"
)

func main() {

	setting := settings.Init()
	setting.Execute()

	var weth common.Address

	if settings.Network == "base" {
		weth = common.HexToAddress("0x4200000000000000000000000000000000000006")
	} else if settings.Network == "ethereum" {
		weth = common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")
	}

	view.CallMethod(weth, "totalSupply()", nil)
}
