package repository

import "errors"

var (
	ErrNotFound      = errors.New("not found")
	ErrDBError       = errors.New("database error")
	ErrAlreadyExists = errors.New("already exists")
	ErrNilCoupon     = errors.New("coupon cannot be nil")
)
