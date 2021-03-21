package entity

type Customer struct {
	ID    uint64
	Email string
	Phone string
}

func NewCustomer(id uint64, email, phone string) *Customer {
	return &Customer{id, email, phone}
}
