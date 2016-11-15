package supermarket

// Product model interface
type Product interface {
	GetPrice() float32
}

// Supermarket model for figuring out pricing
type Supermarket struct {
	Products []Product
}

// GetTotal for returning price of product.
func (s Supermarket) GetTotal() float32 {
	var total float32
	total = 0

	for i := range s.Products {
		total += s.Products[i].GetPrice()
	}

	return total
}
