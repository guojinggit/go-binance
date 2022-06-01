package binance

import (
	"context"
)

//SubAccountUniversalTransferService subaccount universal transfer
type SubAccountUniversalTransferService struct {
	c               *Client
	fromEmail       string
	toEmail         string
	fromAccountType string //	"SPOT","USDT_FUTURE","COIN_FUTURE"
	toAccountType   string //	"SPOT","USDT_FUTURE","COIN_FUTURE"
	asset           string
	amount          string
}

func (s *SubAccountUniversalTransferService) FromEmail(email string) *SubAccountUniversalTransferService {
	s.fromEmail = email
	return s
}

func (s *SubAccountUniversalTransferService) ToEmail(email string) *SubAccountUniversalTransferService {
	s.toEmail = email
	return s
}

func (s *SubAccountUniversalTransferService) FromAccountType(fromType string) *SubAccountUniversalTransferService {
	s.fromAccountType = fromType
	return s
}

func (s *SubAccountUniversalTransferService) ToAccountType(toType string) *SubAccountUniversalTransferService {
	s.toAccountType = toType
	return s
}

func (s *SubAccountUniversalTransferService) Asset(asset string) *SubAccountUniversalTransferService {
	s.asset = asset
	return s
}

func (s *SubAccountUniversalTransferService) Amount(amount string) *SubAccountUniversalTransferService {
	s.amount = amount
	return s
}

func (s *SubAccountUniversalTransferService) Do(ctx context.Context) (tranId string, err error) {
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
