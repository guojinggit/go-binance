package binance

import (
	"context"
)

//用于实现万能划转相关api
type TransferService struct {
	c           	*Client
	fromEmail		string
	toEmail			string
	fromAccountType	string	//	"SPOT","USDT_FUTURE","COIN_FUTURE"
	toAccountType	string	//	"SPOT","USDT_FUTURE","COIN_FUTURE"
	asset			string
	amount			string
}

func(s *TransferService) FromEmail(email string) *TransferService{
	s.fromEmail = email
	return s
}

func(s *TransferService) ToEmail(email string) *TransferService{
	s.toEmail = email
	return s
}

func(s *TransferService) FromAccountType(fromType string)  *TransferService{
	s.fromAccountType = fromType
	return s
}

func(s *TransferService) ToAccountType(toType string) *TransferService{
	s.toAccountType = toType
	return s
}

func(s *TransferService) Asset(asset string)  *TransferService{
	s.asset = asset
	return s
}

func(s *TransferService) Amount(amount string)  *TransferService{
	s.amount = amount
	return s
}

func (s *TransferService) Do(ctx context.Context) (tranId string , err error) {
	r := &request{
		method:   "POST",
		endpoint: "/sapi/v1/sub-account/universalTransfer",
		secType:  secTypeSigned,
	}

	r.setParam("fromEmail", s.fromEmail)
	r.setParam("toEmail", s.toEmail)
	r.setParam("fromAccountType", s.fromAccountType)
	r.setParam("toAccountType", s.toAccountType)
	r.setParam("asset", s.asset)
	r.setParam("amount", s.amount)


	data, err := s.c.callAPI(ctx, r)
	tranId = ""
	if err != nil {
		return
	}

	j, err := newJSON(data)
	if err != nil {
		return "", err
	}
	tranId = j.Get("tranId").MustString()
	return
}






