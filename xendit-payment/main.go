package main

import (
	"fmt"
	balanceData "github.com/alif-github/xendit-payment/balance"
	"github.com/xendit/xendit-go"
)

func main() {
	xendit.Opt.SecretKey = "xnd_development_TiSz3z0w7vPJtasmOvszojdHrIN21tYkq9oqTLHGbxkluLCQm2sDlGzhFPePh"

	//--- Get Balance
	balance := balanceData.GetBalance(string(xendit.BalanceAccountTypeCash))
	fmt.Printf("balance: %v\n", balance)
}
