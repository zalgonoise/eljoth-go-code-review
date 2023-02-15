package memdb

import (
	"fmt"

	"github.com/zalgonoise/eljoth-go-code-review/coupon_service/internal/discount"
)

type Config struct{}

type repository interface {
	FindByCode(string) (*discount.Coupon, error)
	Save(discount.Coupon) error
}

type Repository struct {
	entries map[string]discount.Coupon
}

func New() *Repository {
	return &Repository{}
}

func (r *Repository) FindByCode(code string) (*discount.Coupon, error) {
	coupon, ok := r.entries[code]
	if !ok {
		return nil, fmt.Errorf("Coupon not found")
	}
	return &coupon, nil
}

func (r *Repository) Save(coupon discount.Coupon) error {
	r.entries[coupon.Code] = coupon
	return nil
}
