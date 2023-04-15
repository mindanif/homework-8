package main

import (
	"bufio"
	"context"
	"fmt"
	"homework-6/internal/console"
	"homework-6/internal/pkg/db"
	"homework-6/internal/pkg/repository/postgresql/products"
	"homework-6/internal/pkg/repository/postgresql/warehouses"
	"os"
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

	console := console.NewServer(productRepo, warehouseRepo)

	defer database.GetPool(ctx).Close()

	fmt.Println("Type a command (type 'exit' to quit):")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			continue
		}
		err = console.Action(ctx, input)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}
