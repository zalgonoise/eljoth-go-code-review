package memdb

import (
	"errors"
	"fmt"
	"sync"

	"github.com/zalgonoise/eljoth-go-code-review/coupon_service/internal/discount"
)

var (
	ErrNotFound      = errors.New("not found")
	ErrDBError       = errors.New("database error")
	ErrAlreadyExists = errors.New("already exists")
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
		return nil, fmt.Errorf("%w: no coupon with code %s", ErrNotFound, code)
	}
	if coupon, ok := (c).(discount.Coupon); ok {
		return &coupon, nil
	}
	return nil, fmt.Errorf("%w: invalid coupon type: %T", ErrDBError, c)
}

func (r *CouponsRepository) Save(coupon discount.Coupon) error {
	if _, ok := r.entries.Load(coupon.Code); ok {
		return fmt.Errorf("%w: coupon with code %s", ErrAlreadyExists, coupon.Code)
	}
	r.entries.Store(coupon.Code, coupon)
	return nil
}
