package notification

import (
	"bytes"
	"errors"
	"html/template"
	"purchase-notifi/pkg/entity"
)

const (
	OperatorEmail = "email"
	OperatorSms   = "sms"
)

var (
	tEmail = template.New("Purchase email template")
	tSms   = template.New("Purchase sms template")
)

func init() {
	tEmail, _ = tEmail.Parse(`<h2>Thank you for your order</h2> 
								    <hr>
									Product:<br>
									{{.Product.Name}} : $ {{.Product.Price}}
									<hr>
									Order Total: ${{.Amount}}
									<hr>`)
	tSms, _ = tSms.Parse("Thank you for your order \nProduct: \n{{.Product.Name}} : $ {{.Product.Price}} \nOrder Total: ${{.Amount}}")
}

type notification struct {
	operatorEmail entity.OperatorEmail
	operatorSms   entity.OperatorSms
}

func New(operatorEmail entity.OperatorEmail, operatorSms entity.OperatorSms) entity.Notification {
	return &notification{
		operatorEmail: operatorEmail,
		operatorSms:   operatorSms,
	}
}

func (n *notification) Purchase(operator string, purchase *entity.Purchase) error {
	switch operator {
	case OperatorEmail:
		return n.SendEmail(purchase.Customer.Email, n.emailMessageForPurchase(purchase))
	case OperatorSms:
		return n.SendSms(purchase.Customer.Phone, n.smsMessageForPurchase(purchase))
	}
	return errors.New("invalid operator")
}

func (n *notification) SendEmail(email string, message []byte) error {
	return n.operatorEmail.Send(email, message)
}

func (n *notification) SendSms(phone string, message []byte) error {
	return n.operatorSms.Send(phone, message)
}

func (n *notification) emailMessageForPurchase(purchase *entity.Purchase) []byte {
	var buf bytes.Buffer
	tEmail.Execute(&buf, purchase)
	return buf.Bytes()
}

func (n *notification) smsMessageForPurchase(purchase *entity.Purchase) []byte {
	var buf bytes.Buffer
	tSms.Execute(&buf, purchase)
	return buf.Bytes()
}
