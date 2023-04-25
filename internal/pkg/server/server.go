package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"homework-5/internal/pkg/repository"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

const QueryParamId = "id"

func New(productRepo repository.ProductsRepo, warehouseRepo repository.WarehousesRepo) *Server {
	return &Server{
		productRepo:   productRepo,
		warehouseRepo: warehouseRepo,
	}
}

type Server struct {
	productRepo   repository.ProductsRepo
	warehouseRepo repository.WarehousesRepo
}

func (t *Server) RunServer(ctx context.Context) {
	mux := &http.ServeMux{}
	mux.HandleFunc("/product", func(res http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodPost:
			status := t.createProduct(ctx, req)
			res.WriteHeader(status)
		case http.MethodGet:
			status := t.getProducts(ctx, req)
			res.WriteHeader(status)
		case http.MethodPut:
			status := t.updateProduct(ctx, req)
			res.WriteHeader(status)
		case http.MethodDelete:
			status := t.deleteProduct(ctx, req)
			res.WriteHeader(status)
		default:
			status := t.unsupported(res, req)
			res.WriteHeader(status)
		}
	})
	mux.HandleFunc("/warehouse", func(res http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodGet:
			status := t.getWarehouse(ctx, req)
			res.WriteHeader(status)
		case http.MethodPost:
			status := t.createWarehouse(ctx, req)
			res.WriteHeader(status)
		case http.MethodPut:
			status := t.updateWarehouse(ctx, req)
			res.WriteHeader(status)
		case http.MethodDelete:
			status := t.deleteWarehouse(ctx, req)
			res.WriteHeader(status)
		default:
			status := t.unsupported(res, req)
			res.WriteHeader(status)
		}
	})
}

func (t *Server) createProduct(ctx context.Context, req *http.Request) int {
	product, err := getProductData(req.Body)
	if err != nil {
		log.Println(err)
		return http.StatusBadRequest
	}
	id, err := t.productRepo.Create(ctx, &product)
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError
	}
	log.Println("New product with id:", id)
	return http.StatusOK
}

func (t *Server) getProducts(ctx context.Context, req *http.Request) int {
	warehouseId, err := getIDFromParams(req.URL)
	if err != nil {
		log.Println("invalid WarehouseID")
		return http.StatusBadRequest
	}

	products, err := t.productRepo.List(ctx, warehouseId)
	if err != nil {
		log.Println(err)
		return http.StatusNotFound
	}
	jsonResult, err := json.Marshal(products)
	if err != nil {
		log.Println("error while marshalling products")
		return http.StatusInternalServerError
	}
	log.Println(string(jsonResult))
	return http.StatusOK
}

func (t *Server) updateProduct(ctx context.Context, req *http.Request) int {
	product, err := getProductData(req.Body)
	if err != nil {
		log.Println(err)
		return http.StatusBadRequest
	}
	_, err = t.productRepo.Update(ctx, &product)
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError
	}
	return http.StatusOK
}

func (t *Server) deleteProduct(ctx context.Context, req *http.Request) int {
	value := req.URL.Query().Get(QueryParamId)
	productId, err := strconv.Atoi(value)
	if err != nil {
		log.Println("invalid WarehouseID")
		return http.StatusBadRequest
	}
	ok, err := t.productRepo.Delete(ctx, productId)
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError
	}
	if ok {
		log.Println("Product with id = %1 deleted successfully", productId)
		return http.StatusOK
	}
	log.Println("Product with id = %1 not deleted", productId)
	return http.StatusInternalServerError

}

func (t *Server) createWarehouse(ctx context.Context, req *http.Request) int {
	warehouse, err := getWarehouseData(req.Body)
	if err != nil {
		log.Println(err)
		return http.StatusBadRequest
	}
	id, err := t.warehouseRepo.Add(ctx, &warehouse)
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError
	}
	log.Printf("Warehouse.id: %d", id)
	return http.StatusOK
}

func (t *Server) getWarehouse(ctx context.Context, _ *http.Request) int {
	warehouses, err := t.warehouseRepo.List(ctx)
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError
	}
	jsonResult, err := json.Marshal(warehouses)
	if err != nil {
		log.Println("error while marshalling warehouses")
		return http.StatusInternalServerError
	}
	log.Println(string(jsonResult))
	return http.StatusOK
}

func (t *Server) updateWarehouse(ctx context.Context, req *http.Request) int {
	warehouse, err := getWarehouseData(req.Body)
	if err != nil {
		log.Println(err)
		return http.StatusBadRequest
	}
	updated, err := t.warehouseRepo.Update(ctx, &warehouse)
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError
	}
	log.Println(updated)
	return http.StatusOK
}

func (t *Server) deleteWarehouse(ctx context.Context, req *http.Request) int {
	value := req.URL.Query().Get(QueryParamId)
	productId, err := strconv.Atoi(value)
	if err != nil {
		log.Println("invalid WarehouseID")
		return http.StatusBadRequest
	}
	ok, err := t.productRepo.Delete(ctx, productId)
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError
	}
	if ok {
		log.Println("Warehouse with id = %1 deleted successfully", productId)
		return http.StatusOK
	}
	log.Println("Warehouse with id = %1 not deleted", productId)
	return http.StatusInternalServerError

}

func (t *Server) unsupported(_ http.ResponseWriter, req *http.Request) int {
	log.Printf("Unsupported method для server: %s", req.Method)
	return http.StatusBadRequest
}

func getIDFromParams(reqUrl *url.URL) (int, error) {
	idStr := reqUrl.Query().Get(QueryParamId)
	if len(idStr) == 0 {
		return 0, errors.New("can't get id")
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, fmt.Errorf("can't parse id: %v", err)
	}

	return id, nil
}
func getProductData(reader io.ReadCloser) (repository.Products, error) {
	body, err := io.ReadAll(reader)
	if err != nil {
		return repository.Products{}, err
	}

	data := repository.Products{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}
func getWarehouseData(reader io.ReadCloser) (repository.Warehouses, error) {
	body, err := io.ReadAll(reader)
	if err != nil {
		return repository.Warehouses{}, err
	}

	data := repository.Warehouses{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}
