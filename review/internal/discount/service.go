package discount

// Service describes the available actions exposed to a transport
type Service interface {
	// ApplyCoupon applies the discount from coupon with code `code`, in the shopping basket `basket`
	ApplyCoupon(*Basket, string) error
	// CreateCoupon generates a new Coupon entry from the input discount value `discountVal`, coupon code
	// `code` and with a minimum basket value `minBasketValue`
	CreateCoupon(int, string, int) error
	// GetCoupons will fetch the corresponding coupon for each code provided in the input `codes`, if valid
	GetCoupons(...string) ([]Coupon, error)
	// Close implements the io.Closer interface
	Close() error
}
