package entity

type Purchase struct {
	ID       uint64
	Amount   float64
	Customer *Customer
	Product  *Product
}

func NewPurchase(id uint64, customer *Customer, product *Product) *Purchase {
	return &Purchase{id, product.Price, customer, product}
}
