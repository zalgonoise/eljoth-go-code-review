package entity

import "github.com/zalgonoise/eljoth-go-code-review/coupon_service/internal/discount"

type ApplicationRequest struct {
	Code   string
	Basket discount.Basket
}
