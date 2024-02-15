package entity

type Product struct {
	ID          int64 `db:"id"`
	Name        string `db:"name"`
	Despription string `db:"description"`
	Price       float32 `db:"price"`
	CreatedBy   int64   `db:"created_by"`
}
