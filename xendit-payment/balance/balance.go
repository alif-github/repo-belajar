package balance

import (
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/balance"
	"log"
)

func GetBalance(accountType string) (result float64) {
	data := balance.GetParams{
		AccountType: xendit.BalanceAccountTypeEnum(accountType),
	}

	resp, err := balance.Get(&data)
	if err != nil {
		log.Fatal(err)
	}

	result = resp.Balance
	return
}
