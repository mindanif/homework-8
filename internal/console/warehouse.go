package console

import (
	"context"
	"encoding/json"
	"fmt"
	"homework-5/internal/pkg/repository"
	"strconv"
)

type Warehouse struct {
	warehouseRepo repository.WarehousesRepo
	ctx           context.Context
}

func newWarehouseCommand(ctx context.Context, WarehouseRepo repository.WarehousesRepo) *Warehouse {
	return &Warehouse{
		ctx:           ctx,
		warehouseRepo: WarehouseRepo}
}
func (w *Warehouse) Process(params []string) error {
	switch params[0] {
	case "create":
	case "update":
	case "read":
		err := w.read()
		if err != nil {
			return err
		}
	case "delete":
		err := w.delete(params[1:])
		if err != nil {
			return err
		}
	default:
		return InvalidInput
	}
	return nil
}

func scanWarehouse(params []string) (*repository.Warehouses, error) {
	if len(params) != 3 {
		return nil, InvalidInput
	}
	name := params[0]
	city := params[1]
	square, err := strconv.Atoi(params[2])
	if err != nil {
		return nil, err
	}
	return &repository.Warehouses{Name: name, City: city, Square: square}, nil
}

func (w *Warehouse) create(params []string) error {
	warehouse, err := scanWarehouse(params)
	if err != nil {
		return err
	}
	id, err := w.warehouseRepo.Add(w.ctx, warehouse)
	if err != nil {
		return err
	}
	fmt.Println("New warehouse with id:", id)
	return nil
}

func (w *Warehouse) read() error {
	warehouses, _ := w.warehouseRepo.List(w.ctx)
	jsonProducts, _ := json.Marshal(warehouses)
	fmt.Println(string(jsonProducts))
	return nil
}

func (w *Warehouse) delete(params []string) error {
	id, err := strconv.Atoi(params[0])
	if err != nil {
		return err
	}
	ok, err := w.warehouseRepo.Delete(w.ctx, id)
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
func (w *Warehouse) update(params []string) error {
	id, err := strconv.Atoi(params[0])
	if err != nil {
		return err
	}
	warehouse, err := scanWarehouse(params[1:])
	if err != nil {
		return err
	}
	warehouse.Id = id
	updated, err := w.warehouseRepo.Update(w.ctx, warehouse)
	if err != nil {
		return err
	}
	fmt.Println(updated)
	return nil
}
