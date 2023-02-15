package api

import (
	"errors"
	"net/http"

	"github.com/zalgonoise/eljoth-go-code-review/coupon_service/internal/discount"
	"github.com/zalgonoise/eljoth-go-code-review/coupon_service/internal/repository"

	"github.com/gin-gonic/gin"
)

func (a *API) Apply() gin.HandlerFunc {
	return func(c *gin.Context) {
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
			if errors.Is(err, repository.ErrDBError) {
				c.AbortWithError(http.StatusInternalServerError, err)
				return
			}
			// 404: DB coupon not found
			if errors.Is(err, repository.ErrNotFound) {
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
}
func (a *API) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
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
			if errors.Is(err, repository.ErrAlreadyExists) {
				c.AbortWithError(http.StatusConflict, err)
				return
			}
			// 400: validation error
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		coupons, err := a.svc.GetCoupons(apiReq.Code)
		if err != nil {
			// 500: DB error
			if errors.Is(err, repository.ErrDBError) {
				c.AbortWithError(http.StatusInternalServerError, err)
				return
			}
			// 404: DB coupon not found
			if errors.Is(err, repository.ErrNotFound) {
				c.AbortWithError(http.StatusNotFound, err)
				return
			}
			// 400: validation error
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		if len(coupons) != 1 {
			c.AbortWithError(http.StatusInternalServerError, errors.New("unexpected amount of returned coupons"))
			return
		}

		// 200: OK
		c.JSON(http.StatusOK, coupons[0])
	}
}

func (a *API) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
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
			if errors.Is(err, repository.ErrDBError) {
				c.AbortWithError(http.StatusInternalServerError, err)
				return
			}
			// 404: DB coupon not found
			if errors.Is(err, repository.ErrNotFound) {
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
}

func (a *API) List() gin.HandlerFunc {
	return func(c *gin.Context) {
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
			if errors.Is(err, repository.ErrDBError) {
				c.AbortWithError(http.StatusInternalServerError, err)
				return
			}
			// 404: DB coupon not found
			if errors.Is(err, repository.ErrNotFound) {
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
}
