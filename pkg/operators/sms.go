package operators

import "purchase-notifi/pkg/entity"

func NewSms() entity.OperatorSms {
	return &sms{}
}

type sms struct{}

func (s *sms) Send(phone string, message []byte) error {
	return nil
}
