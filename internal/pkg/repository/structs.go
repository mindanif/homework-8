package repository

import (
	"time"
)

type Warehouses struct {
	Id        int       `db:"id"`
	Name      string    `db:"name"`
	City      string    `db:"city"`
	Square    int       `db:"square"`
	CreatedAt time.Time `db:"created_at"`
}

type Products struct {
	Id          int       `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	Price       int       `db:"price"`
	WarehouseId int       `db:"warehouse_id"`
	CreatedAt   time.Time `db:"created_at"`
}
