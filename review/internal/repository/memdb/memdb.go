package memdb

import (
	"fmt"

	"github.com/zalgonoise/eljoth-go-code-review/coupon_service/internal/discount"
)

// verify that this implementation complies with the repository interface
var _ discount.Repository = &CouponsRepository{}

// CouponsRepository implements discount.Repository
type CouponsRepository struct {
	entries map[string]discount.Coupon
}

func New() *CouponsRepository {
	return &CouponsRepository{
		entries: make(map[string]discount.Coupon),
	}
}

func (r *CouponsRepository) FindByCode(code string) (*discount.Coupon, error) {
	coupon, ok := r.entries[code]
	if !ok {
		return nil, fmt.Errorf("Coupon not found")
	}
	return &coupon, nil
}

func (r *CouponsRepository) Save(coupon discount.Coupon) error {
	r.entries[coupon.Code] = coupon
	return nil
}
