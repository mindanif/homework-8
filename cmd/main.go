package main

import (
	"context"
	"fmt"
	"homework-5/internal/pkg/db"
	"homework-5/internal/pkg/repository/postgresql/products"
	"homework-5/internal/pkg/repository/postgresql/warehouses"
	"homework-5/internal/pkg/server"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "test"
	password = "test"
	dbname   = "test"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	database, err := db.NewDB(ctx, dsn)
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
