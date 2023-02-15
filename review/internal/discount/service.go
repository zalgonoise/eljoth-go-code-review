package discount

// Service describes the available actions exposed to a transport
type Service interface {
	ApplyCoupon(*Basket, string) error
	CreateCoupon(int, string, int) error
	GetCoupons(...string) ([]Coupon, error)
	Close() error
}
