package discount

// Repository describes the available actions
type Repository interface {
	FindByCode(string) (*Coupon, error)
	Save(Coupon) error
}
