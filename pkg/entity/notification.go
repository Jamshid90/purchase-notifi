package entity

type Notification interface {
	Purchase(operator string, purchase *Purchase) error
	SendEmail(email string, message []byte) error
	SendSms(phone string, message []byte) error
}
