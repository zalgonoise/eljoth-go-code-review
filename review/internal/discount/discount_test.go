package discount_test

import (
	"errors"
	"reflect"
	"testing"

	. "github.com/zalgonoise/eljoth-go-code-review/coupon_service/internal/discount"
)

func TestNewCoupon(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var (
			inputCode           = "abc"
			inputDiscount       = 10
			inputMinBasketValue = 10
		)
		wants := &Coupon{
			Code:           inputCode,
			Discount:       inputDiscount,
			MinBasketValue: inputMinBasketValue,
		}

		coupon, err := NewCoupon(inputDiscount, inputCode, inputMinBasketValue)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		wants.ID = coupon.ID
		if !reflect.DeepEqual(wants, coupon) {
			t.Errorf("output mismatch error: wanted %v ; got %v", wants, coupon)
		}
	})

	t.Run("NegativeDiscount", func(t *testing.T) {
		var (
			inputCode           = "abc"
			inputDiscount       = -3
			inputMinBasketValue = 10
		)

		_, err := NewCoupon(inputDiscount, inputCode, inputMinBasketValue)
		if err == nil || !errors.Is(err, ErrInvalidDiscount) {
			t.Errorf("invalid error value: %v", err)
		}
	})

	t.Run("NegativeMinimumBasketValue", func(t *testing.T) {
		var (
			inputCode           = "abc"
			inputDiscount       = 10
			inputMinBasketValue = -3
		)

		_, err := NewCoupon(inputDiscount, inputCode, inputMinBasketValue)
		if err == nil || !errors.Is(err, ErrInvalidMinimumValue) {
			t.Errorf("invalid error value: %v", err)
		}
	})

	t.Run("EmptyCodeString", func(t *testing.T) {
		var (
			inputCode           = ""
			inputDiscount       = 10
			inputMinBasketValue = 10
		)

		_, err := NewCoupon(inputDiscount, inputCode, inputMinBasketValue)
		if err == nil || !errors.Is(err, ErrInvalidCode) {
			t.Errorf("invalid error value: %v", err)
		}
	})
}
