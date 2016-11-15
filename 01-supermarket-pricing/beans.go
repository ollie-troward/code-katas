package supermarket

// Beans model
type Beans struct{}

// GetPrice for returning price.
func (b Beans) GetPrice() float32 {
	return 0.65
}
