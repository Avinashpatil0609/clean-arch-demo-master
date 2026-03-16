package ent

type Address struct {
	street string
	city string
}

type LineItem struct {
	item string
	price int
	amount int
}

type Order struct {
	deliveryAddress Address
	billingAddress Address
	lineItems []LineItem
}

func (o Order) total() int {
	var total int
	for _,li := range o.lineItems {
		total += li.amount * li.price
	}
	return total
}