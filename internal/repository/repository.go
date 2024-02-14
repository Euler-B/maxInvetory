package repository

import (
	"context"

	"github.com/euler-b/maxInventoryProject/internal/entity"
	"github.com/jmoiron/sqlx"
)

// Repository es el conjunto de interfaces que engloban las propiedades basicas CRUD

//go:generate mockery --name=Repository --output=repository --inpackage
type Repository interface {
	SaveUser(ctx context.Context, email, name, password string) error
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)

	SaveUserRole(ctx context.Context, userID, roleID int64) error
	RemoveUserRole(ctx context.Context, userID, roleID int64) error
	GetUserByRoles(ctx context.Context, userID int64) ([]entity.USER_ROLE, error)
}

// Este struct implementa la interfaz de repository
type repo struct {
	db *sqlx.DB // Esta es la referencia a la conexion de la base de datos.
}

func New(db *sqlx.DB) Repository {
	return &repo{
		db: db, // el idiomatic de Go establece que no se deben retornar interfaces, sino structs que las implemente
	} // entonces deberiamos recibir interfaces y retornar structs ( en teoria )
} // Sin embargo aqui se rompe con ese principio, pora no hacer no doble trabajo
