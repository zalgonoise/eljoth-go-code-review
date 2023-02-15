package memdb_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/zalgonoise/eljoth-go-code-review/coupon_service/internal/discount"
	"github.com/zalgonoise/eljoth-go-code-review/coupon_service/internal/repository"
	. "github.com/zalgonoise/eljoth-go-code-review/coupon_service/internal/repository/memdb"
)

func TestRepository(t *testing.T) {
	r := New()
	coupon, err := discount.NewCoupon(10, "abc", 10)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}

	t.Run("Save", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			err := r.Save(coupon)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}
		})

		t.Run("FailNilCoupon", func(t *testing.T) {
			err := r.Save(nil)
			if err == nil || !errors.Is(err, repository.ErrNilCoupon) {
				t.Errorf("unexpected error: %v", err)
				return
			}
		})

		t.Run("FailAlreadyExists", func(t *testing.T) {
			err := r.Save(coupon)
			if err == nil || !errors.Is(err, repository.ErrAlreadyExists) {
				t.Errorf("unexpected error: %v", err)
				return
			}
		})

	})

	t.Run("FindByCode", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			c, err := r.FindByCode(coupon.Code)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if !reflect.DeepEqual(coupon, c) {
				t.Errorf("output mismatch error: wanted %v ; got %v", coupon, c)
			}
		})

		t.Run("FailNotFound", func(t *testing.T) {
			_, err := r.FindByCode("___")
			if err == nil || !errors.Is(err, repository.ErrNotFound) {
				t.Errorf("unexpected error: %v", err)
				return
			}
		})
	})

	t.Run("Close", func(t *testing.T) {
		err := r.Close()
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})
}
