package discount

// Service describes the available actions exposed to a transport
type Service interface {
	ApplyCoupon(*Basket, string) error
	CreateCoupon(int, string, int) any
	GetCoupons([]string) ([]Coupon, error)
}
