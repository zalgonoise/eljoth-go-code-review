package memdb

import (
	"fmt"
	"sync"

	"github.com/zalgonoise/eljoth-go-code-review/coupon_service/internal/discount"
	"github.com/zalgonoise/eljoth-go-code-review/coupon_service/internal/repository"
)

// verify that this implementation complies with the repository interface
var _ discount.Repository = &CouponsRepository{}

// CouponsRepository implements discount.Repository
type CouponsRepository struct {
	entries *sync.Map // map[string]*discount.Coupon
}

// New creates a CouponsRepository based off of an in-memory database
func New() *CouponsRepository {
	return &CouponsRepository{
		entries: new(sync.Map),
	}
}

// FindByCode loads the coupon with the code `code` from the repository, if it exists
func (r *CouponsRepository) FindByCode(code string) (*discount.Coupon, error) {
	c, ok := r.entries.Load(code)
	if !ok {
		return nil, fmt.Errorf("%w: no coupon with code %s", repository.ErrNotFound, code)
	}
	if coupon, ok := (c).(*discount.Coupon); ok {
		return coupon, nil
	}
	return nil, fmt.Errorf("%w: invalid coupon type: %T", repository.ErrDBError, c)
}

// Save stores the input coupon `coupon` in the repository, if not yet present
func (r *CouponsRepository) Save(coupon *discount.Coupon) error {
	if coupon == nil {
		return repository.ErrNilCoupon
	}
	if _, ok := r.entries.Load(coupon.Code); ok {
		return fmt.Errorf("%w: coupon with code %s", repository.ErrAlreadyExists, coupon.Code)
	}
	r.entries.Store(coupon.Code, coupon)
	return nil
}

// Close implements the io.Closer interface
func (r *CouponsRepository) Close() error {
	// in here would go the shutdown routine for a DB that persists the data
	return nil
}
