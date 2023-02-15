package memdb

import (
	"fmt"
	"sync"

	"github.com/zalgonoise/eljoth-go-code-review/coupon_service/internal/discount"
)

// verify that this implementation complies with the repository interface
var _ discount.Repository = &CouponsRepository{}

// CouponsRepository implements discount.Repository
type CouponsRepository struct {
	entries *sync.Map // map[string]discount.Coupon
}

func New() *CouponsRepository {
	return &CouponsRepository{
		entries: new(sync.Map),
	}
}

func (r *CouponsRepository) FindByCode(code string) (*discount.Coupon, error) {
	c, ok := r.entries.Load(code)
	if !ok {
		return nil, fmt.Errorf("Coupon not found")
	}
	if coupon, ok := (c).(discount.Coupon); ok {
		return &coupon, nil
	}
	return nil, fmt.Errorf("internal: invalid coupon type")
}

func (r *CouponsRepository) Save(coupon discount.Coupon) error {
	r.entries.Store(coupon.Code, coupon)
	return nil
}
