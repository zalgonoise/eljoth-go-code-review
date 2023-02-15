package discount

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

var (
	ErrInvalidDiscount     = errors.New("invalid coupon discount value")
	ErrInvalidCode         = errors.New("invalid coupon code")
	ErrInvalidMinimumValue = errors.New("invalid coupon minimum basket value")
)

// Coupon describes the entity
//
// A coupon is used to apply a discount
// to a shopping Basket
type Coupon struct {
	ID             string
	Code           string
	Discount       int
	MinBasketValue int
}

// NewCoupon creates a Coupon with validation
func NewCoupon(discountVal int, code string, minBasketValue int) (*Coupon, error) {
	if discountVal <= 0 {
		return nil, fmt.Errorf("%w: cannot be %d", ErrInvalidDiscount, discountVal)
	}
	if code == "" {
		return nil, fmt.Errorf("%w: empty string", ErrInvalidCode)
	}
	if minBasketValue < 0 {
		return nil, fmt.Errorf("%w: cannot be below zero: %d", ErrInvalidMinimumValue, minBasketValue)
	}

	return &Coupon{
		Discount:       discountVal,
		Code:           code,
		MinBasketValue: minBasketValue,
		ID:             uuid.NewString(),
	}, nil

}

// Basket describes the entity
//
// A Basket is used through a shopping transaction
// to serve as reference for the total price, and the
// applied discount value
type Basket struct {
	Value           int
	AppliedDiscount int
}
