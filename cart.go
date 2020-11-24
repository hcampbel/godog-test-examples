package main

// Comment
type Cart struct {
	items map[string]float64
	total float64
}

func NewCart() *Cart {
	cart := Cart{
		items: make(map[string]float64),
		total: 0,
	}

	return &cart
}

func (cart *Cart) IsInCart(item string) bool {
	_, found := cart.items[item]
	return found
}

func (cart *Cart) GetCartTotal() float64 {
	return cart.total
}

func (cart *Cart) GetItemPrice(item string) float64 {
	return cart.items[item]
}

func (cart *Cart) AddToCart(item string, price float64) {
	// fmt.Printf("%s\n", item)
	// fmt.Printf("%.2f\n", price)
	// if cart.items == nil {
	//	cart.items = make(map[string]float64)
	// }
	cart.items[item] = price
	cart.updateCart()
}

func (cart *Cart) updateCart() {
	cart.total = cart.calcCartTotal()
}

func (cart *Cart) calcCartTotal() float64 {
	var estimatedTotal float64

	for _, price := range cart.items {
		estimatedTotal += price
	}

	return estimatedTotal
}
