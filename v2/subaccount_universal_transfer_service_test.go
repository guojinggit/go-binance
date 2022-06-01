package binance

import (
	"context"
	"fmt"
	"testing"
)

func TestTransferService(t *testing.T) {
	client := NewClient("", "")
	client.Debug = true
	s := client.NewSubAccountUniversalTransferService().
		ToEmail("aaa@binance.com").
		FromEmail("bbb@protonmail.com").
		ToAccountType("SPOT").
		FromAccountType("USDT_FUTURE").
		Asset("USDT").
		Amount("0.6")

	id, err := s.Do(context.Background())
	fmt.Println(id)
	fmt.Println("err:", err)
	//r, err := client.NewGetAccountService().Do(context.Background())
	//fmt.Println(r)
	//fmt.Println(err)
}
