package server

import (
	"bytes"
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"homework-5/internal/pkg/repository"
	"net/http"
	"net/url"
	"testing"
)

const (
	jsonProduct   = `{"name": "name", "price": 100,"warehouseId": 1}`
	jsonWarehouse = `{"name": "name", "square": 10}`
)

func TestServer_createProduct(t *testing.T) {
	t.Parallel()
	var (
		ctx = context.Background()
		id  = 1
	)
	t.Run("success create product", func(t *testing.T) {
		// arrange
		s := setUp(t)
		defer s.tearDown()
		req, err := http.NewRequest(http.MethodPost, "product", bytes.NewReader([]byte(jsonProduct)))
		require.NoError(t, err)

		s.mProduct.EXPECT().Create(gomock.Any(), &repository.Products{
			Name: "name", Price: 100, WarehouseId: 1}).Return(int64(id), nil)
		// act

		status := s.Server.createProduct(ctx, req)
		// assert
		require.Equal(t, http.StatusOK, status)
	})
	t.Run("fail create", func(t *testing.T) {
		s := setUp(t)
		defer s.tearDown()

		t.Parallel()

		t.Run("erroneous data type", func(t *testing.T) {
			t.Parallel()
			req, err := http.NewRequest(http.MethodPost, "product", bytes.NewReader([]byte(
				`{"name": "name", "price": "100" ,"warehouseId": 1}`)))
			status := s.Server.createProduct(ctx, req)
			require.NoError(t, err)
			require.Equal(t, status, http.StatusBadRequest)
		})

	})
}

//
func TestServer_createWarehouse(t *testing.T) {
	t.Parallel()
	var (
		ctx = context.Background()
		id  = 1
	)
	t.Run("success create product", func(t *testing.T) {
		// arrange
		s := setUp(t)
		defer s.tearDown()
		req, err := http.NewRequest(http.MethodPost, "warehouse", bytes.NewReader([]byte(jsonWarehouse)))
		require.NoError(t, err)

		s.mWarehouse.EXPECT().Add(gomock.Any(), &repository.Warehouses{
			Name: "name", Square: 10}).Return(id, nil)
		// act

		status := s.Server.createWarehouse(ctx, req)
		// assert
		require.Equal(t, http.StatusOK, status)
	})
	t.Run("fail create", func(t *testing.T) {
		s := setUp(t)
		defer s.tearDown()

		t.Parallel()

		t.Run("erroneous data type", func(t *testing.T) {
			t.Parallel()
			req, err := http.NewRequest(http.MethodPost, "warehouse", bytes.NewReader([]byte(
				`{"name": "name", "square": "sdf"}`)))
			status := s.Server.createWarehouse(ctx, req)
			require.NoError(t, err)
			require.Equal(t, status, http.StatusBadRequest)
		})

	})
}

//
func TestServer_deleteProduct(t *testing.T) {
	var (
		ctx = context.Background()
		id  = 1
	)
	t.Parallel()
	t.Run("success delete product", func(t *testing.T) {
		// arrange
		s := setUp(t)
		defer s.tearDown()
		req, err := http.NewRequest(http.MethodDelete, "product?id=1", bytes.NewReader([]byte{}))
		require.NoError(t, err)

		s.mProduct.EXPECT().Delete(gomock.Any(), id).Return(true, nil)
		// act

		status := s.Server.deleteProduct(ctx, req)
		// assert
		require.Equal(t, http.StatusOK, status)
	})
}

func TestServer_getProducts(t *testing.T) {
	var (
		ctx = context.Background()
		id  = 1
	)
	t.Parallel()
	t.Run("success get product", func(t *testing.T) {
		// arrange
		s := setUp(t)
		defer s.tearDown()
		req, err := http.NewRequest(http.MethodGet, "product?id=1", bytes.NewReader([]byte{}))
		require.NoError(t, err)

		s.mProduct.EXPECT().List(gomock.Any(), id).Return([]*repository.Products{
			&repository.Products{Id: 1, Name: "asd", WarehouseId: 1},
			&repository.Products{Id: 2, Name: "asfsf", WarehouseId: 1},
		}, nil)
		// act

		status := s.Server.getProducts(ctx, req)
		// assert
		require.Equal(t, http.StatusOK, status)
	})
	t.Run("fail", func(t *testing.T) {
		s := setUp(t)
		defer s.tearDown()

		t.Parallel()
		tt := []struct {
			name    string
			request *url.URL
			isOk    bool
		}{
			{
				"without id",
				&url.URL{RawQuery: "user?id"},
				false,
			},
			{
				"wrong id",
				&url.URL{RawQuery: "user?id=asdasd"},
				false,
			},
			{
				"empty",
				&url.URL{RawQuery: ""},
				false,
			},
			{
				"ok",
				&url.URL{RawQuery: "user?id=1"},
				true,
			},
		}
		for _, tc := range tt {
			tc := tc
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()
				id, err := getID(tc.request)
				if !tc.isOk {
					assert.EqualError(t, err, "can't get id")
				} else {
					assert.Equal(t, 0, id)
				}
			})
		}

	})
}

func TestServer_getWarehouse(t *testing.T) {
	var (
		ctx = context.Background()
	)
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		// arrange
		s := setUp(t)
		defer s.tearDown()
		req, err := http.NewRequest(http.MethodGet, "warehouse", bytes.NewReader([]byte{}))
		require.NoError(t, err)

		s.mWarehouse.EXPECT().List(gomock.Any()).Return([]*repository.Warehouses{
			&repository.Warehouses{Id: 1, Name: "asd"},
			&repository.Warehouses{Id: 2, Name: "asfsf"},
		}, nil)
		// act

		status := s.Server.getWarehouse(ctx, req)
		// assert
		require.Equal(t, http.StatusOK, status)
	})
}

func TestServer_updateProduct(t *testing.T) {
	t.Parallel()
	var (
		ctx = context.Background()
		//id  = 1
	)
	t.Run("success update product", func(t *testing.T) {
		// arrange
		s := setUp(t)
		defer s.tearDown()
		req, err := http.NewRequest(http.MethodPut, "product", bytes.NewReader([]byte(jsonProduct)))
		require.NoError(t, err)

		s.mProduct.EXPECT().Update(gomock.Any(), &repository.Products{
			Name: "name", Price: 100, WarehouseId: 1}).Return(true, nil)
		// act

		status := s.Server.updateProduct(ctx, req)
		// assert
		require.Equal(t, http.StatusOK, status)
	})
	t.Run("fail update", func(t *testing.T) {
		s := setUp(t)
		defer s.tearDown()

		t.Parallel()

		req, err := http.NewRequest(http.MethodPut, "product", bytes.NewReader([]byte(
			`{"name": "name", "price": "100" ,"warehouseId": 1}`)))
		status := s.Server.createProduct(ctx, req)
		require.NoError(t, err)
		require.Equal(t, status, http.StatusBadRequest)

	})
}
