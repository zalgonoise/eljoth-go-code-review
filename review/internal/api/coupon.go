package api

import (
	"errors"
	"net/http"

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
	type coupon struct {
		Discount       int    `json:"discount,omitempty"`
		Code           string `json:"code,omitempty"`
		MinBasketValue int    `json:"min_basket_value,omitempty"`
	}

	apiReq := coupon{}
	if err := c.ShouldBindJSON(&apiReq); err != nil {
		// 400: validation error
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	err := a.svc.CreateCoupon(apiReq.Discount, apiReq.Code, apiReq.MinBasketValue)
	if err != nil {
		// 409: conflict with existing coupon code
		if errors.Is(err, memdb.ErrAlreadyExists) {
			c.AbortWithError(http.StatusConflict, err)
			return
		}
		// 400: validation error
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// 200: OK
	c.Status(http.StatusOK)
}

func (a *API) Get(c *gin.Context) {
	type couponRequest struct {
		Code string `json:"code,omitempty"`
	}

	apiReq := couponRequest{}
	if err := c.ShouldBindJSON(&apiReq); err != nil {
		// 400: validation error
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	coupon, err := a.svc.GetCoupons(apiReq.Code)
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
	c.JSON(http.StatusOK, coupon)
}

func (a *API) List(c *gin.Context) {
	type couponRequest struct {
		Codes []string `json:"codes,omitempty"`
	}

	apiReq := couponRequest{}
	if err := c.ShouldBindJSON(&apiReq); err != nil {
		// 400: validation error
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	coupons, err := a.svc.GetCoupons(apiReq.Codes...)
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
	c.JSON(http.StatusOK, coupons)
}
