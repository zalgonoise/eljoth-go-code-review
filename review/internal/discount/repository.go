package discount

// Repository describes the available actions
type Repository interface {
	// FindByCode loads the coupon with the code `code` from the repository, if it exists
	FindByCode(string) (*Coupon, error)
	// Save stores the input coupon `coupon` in the repository, if not yet present
	Save(*Coupon) error
	// Close implements the io.Closer interface
	Close() error
}
