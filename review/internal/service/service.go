package service

import (
	"fmt"

	"github.com/zalgonoise/eljoth-go-code-review/coupon_service/internal/discount"

	"github.com/google/uuid"
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

func (s CouponService) ApplyCoupon(basket discount.Basket, code string) (b *discount.Basket, e error) {
	b = &basket
	coupon, err := s.repo.FindByCode(code)
	if err != nil {
		return nil, err
	}

	if b.Value > 0 {
		b.AppliedDiscount = coupon.Discount
	}
	if b.Value == 0 {
		return
	}

	return nil, fmt.Errorf("Tried to apply discount to negative value")
}

func (s CouponService) CreateCoupon(discountVal int, code string, minBasketValue int) any {
	coupon := discount.Coupon{
		Discount:       discountVal,
		Code:           code,
		MinBasketValue: minBasketValue,
		ID:             uuid.NewString(),
	}

	if err := s.repo.Save(coupon); err != nil {
		return err
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
