//go:build integration
// +build integration

package tests

import (
	"context"
	"github.com/stretchr/testify/assert"
	"homework-5/internal/pkg/repository"
	"homework-5/internal/pkg/repository/postgresql/products"
	"testing"
)

func TestCreatePoduct(t *testing.T) {
	t.Run("succes", func(t *testing.T) {
		t.Parallel()
		Db.SetUp(t)
		defer Db.TearDown()
		//arrange
		productRepo := products.NewProducts(Db.DB)
		//act
		res, err = productRepo.Create(context.Background(), &repository.Products{
			Name:        "qwer",
			Price:       10,
			WarehouseId: 1,
		})

		//assert
		assert.NoError(t, err)
		assert.Equal(t, res, 1)
	})
	t.Run("fail", func(t *testing.T) {
		t.Parallel()
		Db.SetUp(t)
		defer Db.TearDown()
		//arrange
		productRepo := products.NewProducts(Db.DB)
		//act
		res, err = productRepo.Create(context.Background(), &repository.Products{
			WarehouseId: 1,
		})

		//assert
		assert.Error(t, err)
	})
}
