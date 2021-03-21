package operators

import "purchase-notifi/pkg/entity"

func NewEmail() entity.OperatorEmail {
	return &email{}
}

type email struct{}

func (e *email) Send(email string, message []byte) error {
	return nil
}
