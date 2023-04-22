package console

import (
	"context"
	"encoding/json"
	"fmt"
	"homework-5/internal/pkg/repository"
	"strconv"
)

type Product struct {
	productRepo repository.ProductsRepo
	ctx         context.Context
}

func newProductCommand(ctx context.Context, ProductRepo repository.ProductsRepo) *Product {
	return &Product{
		ctx:         ctx,
		productRepo: ProductRepo}
}
func (p *Product) Process(params []string) error {
	switch params[0] {
	case "update":
		err := p.update(params[1:])
		if err != nil {
			return err
		}
	case "create":
		err := p.create(params[1:])
		if err != nil {
			return err
		}
	case "delete":
		err := p.delete(params[1:])
		if err != nil {
			return err
		}
	case "read":
		err := p.read(params[1:])
		if err != nil {
			return err
		}
	default:
		return InvalidInput
	}

	return nil
}
func scanProduct(params []string) (*repository.Products, error) {
	if len(params) != 4 {
		return nil, InvalidInput
	}
	name := params[0]
	description := params[1]
	price, err := strconv.Atoi(params[2])
	if err != nil {
		return nil, err
	}
	warehouseId, err := strconv.Atoi(params[3])
	if err != nil {
		return nil, err
	}
	return &repository.Products{Name: name, Description: description, Price: price, WarehouseId: warehouseId}, nil
}
func (p *Product) create(params []string) error {
	product, err := scanProduct(params)
	if err != nil {
		return err
	}
	id, err := p.productRepo.Create(p.ctx, product)
	if err != nil {
		return err
	}
	fmt.Println("New product with id:", id)
	return nil
}
func (p *Product) delete(params []string) error {
	id, err := strconv.Atoi(params[0])
	if err != nil {
		return err
	}
	ok, err := p.productRepo.Delete(p.ctx, id)
	if err != nil {
		return err
	}
	if ok {
		fmt.Println("Product with id = %1 deleted successfully", id)
		return nil
	}
	fmt.Println("Product with id = %1 not deleted", id)
	return InvalidInput
}
func (p *Product) update(params []string) error {
	id, err := strconv.Atoi(params[0])
	if err != nil {
		return err
	}
	product, err := scanProduct(params[1:])
	if err != nil {
		return err
	}
	product.Id = id
	updated, err := p.productRepo.Update(p.ctx, product)
	if err != nil {
		return err
	}
	fmt.Println(updated)
	return nil
}
func (p *Product) read(params []string) error {
	id, err := strconv.Atoi(params[0])
	if err != nil {
		return err
	}
	products, err := p.productRepo.List(p.ctx, id)
	if err != nil {
		fmt.Println(err)
	}
	jsonProducts, _ := json.Marshal(products)
	fmt.Println(string(jsonProducts))
	return nil
}
