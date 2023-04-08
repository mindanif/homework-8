package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"homework-5/internal/pkg/db"
	"homework-5/internal/pkg/repository"
	"homework-5/internal/pkg/repository/postgresql/products"
	"homework-5/internal/pkg/repository/postgresql/warehouses"
	"os"
	"strings"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	database, err := db.NewDB(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	productRepo := products.NewProducts(database)
	warehouseRepo := warehouses.NewWarehouses(database)

	defer database.GetPool(ctx).Close()

	fmt.Println("Type a command (type 'exit' to quit):")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		input = strings.TrimSpace(input)
		switch input {
		case "Create product":
			product, err := scanProduct()
			if err != nil {
				fmt.Println(err)
				continue
			}
			id, err := productRepo.Create(ctx, product)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("New product with id:", id)
		case "Update product":
			fmt.Println("Please enter the product ID:")
			var id int
			_, err := fmt.Scanln(&id)
			if err != nil {
				fmt.Println(err)
				continue
			}
			product, err := scanProduct()
			if err != nil {
				fmt.Println(err)
				continue
			}
			product.Id = id
			updated, err := productRepo.Update(ctx, product)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println(updated)
		case "Delete product":
			fmt.Println("Please enter the product ID:")
			var id int
			_, err := fmt.Scanln(&id)
			if err != nil {
				fmt.Println(err)
				continue
			}
			ok, err := productRepo.Delete(ctx, id)
			if err != nil {
				fmt.Println(err)
				continue
			}
			if ok {
				fmt.Println("Product with id = %1 deleted successfully", id)
				continue
			}
			if !ok {
				fmt.Println("Product with id = %1 not deleted", id)
				continue
			}
		case "Read products in warehouse":
			fmt.Println("Please enter the warehouse ID:")
			var id int
			_, err := fmt.Scanln(&id)
			if err != nil {
				fmt.Println(err)
				continue
			}
			products, err := productRepo.List(ctx, id)
			if err != nil {
				fmt.Println(err)
			}
			jsonProducts, _ := json.Marshal(products)
			fmt.Println(string(jsonProducts))
		case "Create warehouse":
			warehouse, err := scanWarehouse()
			if err != nil {
				fmt.Println(err)
				continue
			}
			id, err := warehouseRepo.Add(ctx, warehouse)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("New warehouse with id:", id)
		case "Update warehouse":
			fmt.Println("Please enter the warehouse ID:")
			var id int
			_, err := fmt.Scanln(&id)
			if err != nil {
				fmt.Println(err)
				continue
			}
			warehouse, err := scanWarehouse()
			if err != nil {
				fmt.Println(err)
				continue
			}
			warehouse.Id = id
			updated, err := warehouseRepo.Update(ctx, warehouse)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println(updated)
		case "Read warehouses":
			warehouses, _ := warehouseRepo.List(ctx)
			jsonProducts, _ := json.Marshal(warehouses)
			fmt.Println(string(jsonProducts))
		case "Delete warehouse":
			fmt.Println("Please enter the warehouse ID:")
			var id int
			_, err := fmt.Scanln(&id)
			if err != nil {
				fmt.Println(err)
				continue
			}
			ok, err := warehouseRepo.Delete(ctx, id)
			if err != nil {
				fmt.Println(err)
				continue
			}
			if ok {
				fmt.Println("Product with id = %1 deleted successfully", id)
				continue
			}
			if !ok {
				fmt.Println("Product with id = %1 not deleted", id)
				continue
			}
		case "exit":
			break
		default:
			fmt.Println("Wrong command:", input)
		}
	}
	fmt.Println("Exiting...")
}

func scanProduct() (*repository.Products, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please enter the product name:")
	name, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	fmt.Println("Please enter the product description:")
	description, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	fmt.Println("Please enter the product price:")
	var price int
	_, err = fmt.Scanf("%d", &price)
	if err != nil {
		return nil, err
	}
	fmt.Println("Please enter the warehouse id:")
	var warehouseId int
	_, err = fmt.Scanf("%d", &warehouseId)
	if err != nil {
		return nil, err
	}
	return &repository.Products{Name: name, Description: description, Price: price, WarehouseId: warehouseId}, nil
}

func scanWarehouse() (*repository.Warehouses, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please enter the warehouse name:")
	name, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	fmt.Println("Please enter the warehouse city:")
	city, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	fmt.Println("Please enter the warehouse square:")
	var square int
	square, err = fmt.Scanf("%d", &square)
	if err != nil {
		return nil, err
	}
	return &repository.Warehouses{Name: name, City: city, Square: square}, nil
}
