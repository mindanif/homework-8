package console

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"homework-5/internal/pkg/repository"
	"os"
	"strings"
)

var (
	InvalidInput = errors.New("Invalid Input")
)

type server struct {
	ProductRepo    repository.ProductsRepo
	WarehousesRepo repository.WarehousesRepo
}

func NewServer(ProductRepo repository.ProductsRepo, WarehousesRepo repository.WarehousesRepo) *server {
	return &server{
		ProductRepo:    ProductRepo,
		WarehousesRepo: WarehousesRepo,
	}
}

type consoleCommand interface {
	Process(params []string) error
}

func (s *server) Action(ctx context.Context, input string) error {
	fmt.Println("Type a command (type 'exit' to quit):")
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("> ")
	input, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	input = strings.TrimSpace(input)
	input = strings.ToLower(input)
	params := strings.Split(input, " ")
	var consoleCommand consoleCommand
	switch params[0] {
	case "help":
		consoleCommand = newHelpCommand()
	case "product":
		consoleCommand = newProductCommand(ctx, s.ProductRepo)
	case "warehouse":
		consoleCommand = newWarehouseCommand(ctx, s.WarehousesRepo)
	case "spell":
		consoleCommand = newSpellCommand()
	default:
		return InvalidInput
	}
	return consoleCommand.Process(params[1:])

}
