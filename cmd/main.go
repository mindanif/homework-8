package main

import (
	"context"
	"fmt"
	"homework-5/internal/pkg/db"
	"homework-5/internal/pkg/repository/postgresql/products"
	"homework-5/internal/pkg/repository/postgresql/warehouses"
	"homework-5/internal/pkg/server"
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

	consoleServer := server.New(productRepo, warehouseRepo)

	consoleServer.RunServer(ctx)
}
