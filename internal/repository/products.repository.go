package repository

import (
	"context"

	"github.com/euler-b/maxInventoryProject/internal/entity"
)

const (
	queryInsertProduct = `
		insert into PRODUCTS (name, description, price, created_by) values(?,?,?,?);
	`
	queryGetAllProducts = `
		select 
			id,
			name,
			description,
			price,
			created_by
		from PRODUCTS;
	`

	queryGetProductByID = `
	select
		id,
		name,
		description,
		price,
		created_by
	from PRODUCTS
	where id = ?;
	`
)

func (r *repo) SaveProduct(ctx context.Context, name, description string, price float32, createdBy int64) error {
	_, err := r.db.ExecContext(ctx, queryInsertProduct, name, description, price, createdBy)
	return err
}

func (r *repo) GetProducts(ctx context.Context) ([]entity.Product, error) {
	pp := []entity.Product{}

	err := r.db.SelectContext(ctx, &pp, queryGetAllProducts) // SelectContext es para obtener multiples filas 
	if err != nil {
		return nil, err
	}
	return pp,err
}

func (r *repo) GetProduct(ctx context.Context, id int64) (*entity.Product, error) {
	p := &entity.Product{}
	err := r.db.GetContext(ctx, &p, queryGetProductByID, id) // GetContext me permite obtener solo una fila
	if err != nil {
		return nil, err
	}
	return p, err
}
