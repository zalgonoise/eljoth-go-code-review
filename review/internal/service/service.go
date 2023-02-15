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
	ErrZeroCodes     = errors.New("no coupon codes were provided")
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

func (s CouponService) GetCoupons(codes ...string) ([]discount.Coupon, error) {
	if len(codes) == 0 {
		return nil, ErrZeroCodes
	}
	coupons := make([]discount.Coupon, 0, len(codes))

	for idx, code := range codes {
		if code == "" {
			return nil, fmt.Errorf("%w: couldn't fetch code on index %d", discount.ErrInvalidCode, idx)
		}
		coupon, err := s.repo.FindByCode(code)
		if err != nil {
			return nil, fmt.Errorf("%w: couldn't fetch code on index %d", err, idx)
		}
		coupons = append(coupons, *coupon)
	}
	return coupons, nil
}

func (s CouponService) Close() error {
	return s.repo.Close()
}
