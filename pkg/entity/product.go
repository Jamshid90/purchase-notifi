package entity

type Product struct {
	Name  string
	Price float64
}

func NewProduct(name string, amount float64) *Product {
	return &Product{name, amount}
}
