package console

import (
	"encoding/json"
	"fmt"
)

type helpCommand struct {
	help map[string]*map[string]string
}

func newHelpCommand() *helpCommand {
	help := map[string]*map[string]string{
		"Консольные команды": {
			"Help":  "Получить информацию обо всех имеющихся командах ",
			"Spell": "Вывести в консоль все буквы переданного на вход слова через пробел",
		},
		"Warehouse": {
			"Create [name] [city] [square]":      "Создать новый склад",
			"Update [id] [name] [city] [square]": "Обновить информацию о склады по id",
			"Read":                               "Получить все склады",
			"Delete [id]":                        "Удалить склад по id",
		},
		"Product": {
			"Create [name] [description] [price] [warehouse Id]":      "Создать новый товар",
			"Update [id] [name] [description] [price] [warehouse Id]": "Обновить информацию о товаре по id",
			"Read [warehouse Id]": "Получить все товары в определенном складе",
			"Delete [id]":         "Удалить товар по id",
		},
	}
	return &helpCommand{
		help: help,
	}
}

func (s *helpCommand) Process(params []string) error {

	jsonData, err := json.MarshalIndent(s.help, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(jsonData))
	return nil
}
