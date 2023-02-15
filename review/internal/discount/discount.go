package discount

// Coupon describes the entity
//
// A coupon is used to apply a discount
// to a shopping Basket
type Coupon struct {
	ID             string
	Code           string
	Discount       int
	MinBasketValue int
}

// Basket describes the entity
//
// A Basket is used through a shopping transaction
// to serve as reference for the total price, and the
// applied discount value
type Basket struct {
	Value           int
	AppliedDiscount int
}
