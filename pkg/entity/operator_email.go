package entity

type OperatorEmail interface {
	Send(email string, message []byte) error
}
