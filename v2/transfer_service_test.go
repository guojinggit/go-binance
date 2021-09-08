package binance

import (
	"context"
	"fmt"
	"testing"
)

func TestTransferService(t *testing.T) {
	client := NewClient("", "")
	client.Debug = true
	s := client.NewTransferService().
		FromEmail("cn18380434045_mobileuser@binance.com").
		ToEmail("binance_arb2@protonmail.com").
		FromAccountType("SPOT").
		ToAccountType("USDT_FUTURE").
		Asset("BUSD").
		Amount("0.5")

	id, err := s.Do(context.Background())
	fmt.Println(id)
	fmt.Println("err:", err)
	//r, err := client.NewGetAccountService().Do(context.Background())
	//fmt.Println(r)
	//fmt.Println(err)
}


