package products

import (
	"context"
	"homework-6/internal/pkg/db"
	"homework-6/internal/pkg/repository"
)

type ProductsRepo struct {
	db db.DBops
}

func NewProducts(db db.DBops) *ProductsRepo {
	return &ProductsRepo{db: db}
}

// Add specific user
func (r *ProductsRepo) Create(ctx context.Context, product *repository.Products) (int64, error) {
	var id int64
	err := r.db.ExecQueryRow(ctx, `INSERT INTO products(name, description, price, warehouse_id) VALUES ($1, $2, $3, $4) RETURNING id`,
		product.Name, product.Description, product.Price, product.WarehouseId).Scan(&id)
	return id, err
}

func (r *ProductsRepo) List(ctx context.Context, warehouseId int) ([]*repository.Products, error) {
	products := make([]*repository.Products, 0)
	err := r.db.Select(ctx, &products,
		"SELECT * FROM products WHERE warehouse_id = $1", warehouseId)
	return products, err
}

func (r *ProductsRepo) Update(ctx context.Context, product *repository.Products) (bool, error) {
	result, err := r.db.Exec(ctx,
		"UPDATE users SET name = $1 WHERE id = $2", product.Name, product.Id)
	return result.RowsAffected() > 0, err
}
func (r *ProductsRepo) Delete(ctx context.Context, id int) (bool, error) {
	result, err := r.db.Exec(ctx,
		"DELETE FROM products WHERE id = $1", id)
	return result.RowsAffected() > 0, err
}
