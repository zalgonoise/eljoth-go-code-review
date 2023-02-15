package service

import (
	"errors"
	"fmt"

	"github.com/zalgonoise/eljoth-go-code-review/coupon_service/internal/discount"
)

var (
	ErrEmptyCode     = errors.New("coupon code cannot be empty")
	ErrNilBasket     = errors.New("input basket cannot be nil")
	ErrNegativeValue = errors.New("basket value cannot be a negative amount")
)

// verify that this implementation complies with the service interface
var _ discount.Service = CouponService{}

type CouponService struct {
	repo discount.Repository
}

func New(repo discount.Repository) CouponService {
	return CouponService{
		repo: repo,
	}
}

func (s CouponService) ApplyCoupon(basket *discount.Basket, code string) error {
	if code == "" {
		return ErrEmptyCode
	}
	if basket == nil {
		return ErrNilBasket
	}
	if basket.Value < 0 {
		return ErrNegativeValue
	}
	if basket.Value == 0 {
		// short-circuit out; empty basket
		return nil
	}

	coupon, err := s.repo.FindByCode(code)
	if err != nil {
		return fmt.Errorf("%w: failed to apply discount", err)
	}

	basket.AppliedDiscount = coupon.Discount
	return nil
}

func (s CouponService) CreateCoupon(discountVal int, code string, minBasketValue int) error {
	coupon, err := discount.NewCoupon(discountVal, code, minBasketValue)
	if err != nil {
		return fmt.Errorf("%w: failed to create coupon", err)
	}
	if err := s.repo.Save(coupon); err != nil {
		return fmt.Errorf("%w: failed to create coupon", err)
	}
	return nil
}

func (s CouponService) GetCoupons(codes []string) ([]discount.Coupon, error) {
	coupons := make([]discount.Coupon, 0, len(codes))
	var e error = nil

	for idx, code := range codes {
		coupon, err := s.repo.FindByCode(code)
		if err != nil {
			if e == nil {
				e = fmt.Errorf("code: %s, index: %d", code, idx)
			} else {
				e = fmt.Errorf("%w; code: %s, index: %d", e, code, idx)
			}
		}
		coupons = append(coupons, *coupon)
	}

	return coupons, e
}
