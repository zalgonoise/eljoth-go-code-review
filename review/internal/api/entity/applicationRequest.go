package entity

import "github.com/zalgonoise/eljoth-go-code-review/coupon_service/internal/service/entity"

type ApplicationRequest struct {
	Code   string
	Basket entity.Basket
}
