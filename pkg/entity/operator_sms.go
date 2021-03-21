package entity

type OperatorSms interface {
	Send(phone string, message []byte) error
}
