package api

import (
	"errors"
	"net/http"

	. "github.com/zalgonoise/eljoth-go-code-review/coupon_service/internal/api/entity"
	"github.com/zalgonoise/eljoth-go-code-review/coupon_service/internal/discount"
	"github.com/zalgonoise/eljoth-go-code-review/coupon_service/internal/repository/memdb"

	"github.com/gin-gonic/gin"
)

func (a *API) Apply(c *gin.Context) {
	type applicationRequest struct {
		Code   string           `json:"code,omitempty"`
		Basket *discount.Basket `json:"basket,omitempty"`
	}

	apiReq := applicationRequest{}
	if err := c.ShouldBindJSON(&apiReq); err != nil {
		// 400: validation error
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := a.svc.ApplyCoupon(apiReq.Basket, apiReq.Code)
	if err != nil {
		// 500: DB error
		if errors.Is(err, memdb.ErrDBError) {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		// 404: DB coupon not found
		if errors.Is(err, memdb.ErrNotFound) {
			c.AbortWithError(http.StatusNotFound, err)
			return
		}
		// 400: validation error
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// 200: OK
	c.JSON(http.StatusOK, apiReq.Basket)
}

func (a *API) Create(c *gin.Context) {
	apiReq := Coupon{}
	if err := c.ShouldBindJSON(&apiReq); err != nil {
		return
	}
	err := a.svc.CreateCoupon(apiReq.Discount, apiReq.Code, apiReq.MinBasketValue)
	if err != nil {
		return
	}
	c.Status(http.StatusOK)
}

func (a *API) Get(c *gin.Context) {
	apiReq := CouponRequest{}
	if err := c.ShouldBindJSON(&apiReq); err != nil {
		return
	}
	coupons, err := a.svc.GetCoupons(apiReq.Codes...)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, coupons)
}
