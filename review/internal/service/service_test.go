package service

import (
	"errors"
	"reflect"
	"testing"

	"github.com/zalgonoise/eljoth-go-code-review/coupon_service/internal/discount"
	"github.com/zalgonoise/eljoth-go-code-review/coupon_service/internal/repository"
	"github.com/zalgonoise/eljoth-go-code-review/coupon_service/internal/repository/memdb"
)

func TestService(t *testing.T) {
	s := New(memdb.New())
	inputCode := "abc"
	inputDiscount := 10
	inputMinBasketValue := 10

	t.Run("CreateCoupon", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			err := s.CreateCoupon(inputDiscount, inputCode, inputMinBasketValue)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}
		})
		t.Run("FailInvalidCode", func(t *testing.T) {
			err := s.CreateCoupon(inputDiscount, "", inputMinBasketValue)
			if err == nil || !errors.Is(err, discount.ErrInvalidCode) {
				t.Errorf("unexpected error: %v", err)
				return
			}
		})

		t.Run("FailAlreadyExists", func(t *testing.T) {
			err := s.CreateCoupon(inputDiscount, inputCode, inputMinBasketValue)
			if err == nil || !errors.Is(err, repository.ErrAlreadyExists) {
				t.Errorf("unexpected error: %v", err)
				return
			}
		})
	})
	t.Run("GetCoupons", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			wants := &discount.Coupon{
				Discount:       inputDiscount,
				Code:           inputCode,
				MinBasketValue: inputMinBasketValue,
			}
			coupons, err := s.GetCoupons(inputCode)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}
			if len(coupons) != 1 {
				t.Errorf("unexpected coupons slice length: wanted %d ; got %d", 1, len(coupons))
				return
			}
			wants.ID = coupons[0].ID
			if !reflect.DeepEqual(*wants, coupons[0]) {
				t.Errorf("output mismatch error: wanted %v ; got %v", wants, coupons[0])
			}
		})

		t.Run("FailZeroLen", func(t *testing.T) {
			_, err := s.GetCoupons()
			if err == nil || !errors.Is(err, ErrZeroCodes) {
				t.Errorf("unexpected error: %v", err)
				return
			}
		})

		t.Run("FailOneIsInvalid", func(t *testing.T) {
			_, err := s.GetCoupons(inputCode, "")
			if err == nil || !errors.Is(err, discount.ErrInvalidCode) {
				t.Errorf("unexpected error: %v", err)
				return
			}
		})
		t.Run("FailOneIsNotFound", func(t *testing.T) {
			_, err := s.GetCoupons(inputCode, "___")
			if err == nil || !errors.Is(err, repository.ErrNotFound) {
				t.Errorf("unexpected error: %v", err)
				return
			}
		})
	})
	t.Run("ApplyCoupon", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			input := &discount.Basket{
				Value:           100,
				AppliedDiscount: 0,
			}
			err := s.ApplyCoupon(input, inputCode)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}
			if input.AppliedDiscount != inputDiscount {
				t.Errorf("unexpected discount value in basket: wanted %d ; got %d", inputDiscount, input.AppliedDiscount)
			}
		})
		t.Run("SuccessZeroValue", func(t *testing.T) {
			input := &discount.Basket{
				Value:           0,
				AppliedDiscount: 0,
			}
			err := s.ApplyCoupon(input, inputCode)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}
		})
		t.Run("FailNilBasket", func(t *testing.T) {
			err := s.ApplyCoupon(nil, inputCode)
			if err == nil || !errors.Is(err, ErrNilBasket) {
				t.Errorf("unexpected error: %v", err)
				return
			}
		})

		t.Run("FailEmptyCode", func(t *testing.T) {
			input := &discount.Basket{
				Value:           100,
				AppliedDiscount: 0,
			}
			err := s.ApplyCoupon(input, "")
			if err == nil || !errors.Is(err, ErrEmptyCode) {
				t.Errorf("unexpected error: %v", err)
				return
			}
		})

		t.Run("FailNegativeValue", func(t *testing.T) {
			input := &discount.Basket{
				Value:           -100,
				AppliedDiscount: 0,
			}
			err := s.ApplyCoupon(input, inputCode)
			if err == nil || !errors.Is(err, ErrNegativeValue) {
				t.Errorf("unexpected error: %v", err)
				return
			}
		})

		t.Run("FailNotFoundCode", func(t *testing.T) {
			input := &discount.Basket{
				Value:           100,
				AppliedDiscount: 0,
			}
			err := s.ApplyCoupon(input, "___")
			if err == nil || !errors.Is(err, repository.ErrNotFound) {
				t.Errorf("unexpected error: %v", err)
				return
			}
		})
	})
	t.Run("Close", func(t *testing.T) {
		err := s.Close()
		if err != nil {
			t.Errorf("unexpected error: %v", err)
			return
		}
	})
}
