package service

import (
	"context"

	"github.com/euler-b/maxInventoryProject/internal/models"
	"github.com/euler-b/maxInventoryProject/internal/repository"
)

// service es la logica de negocio de la aplicacion
//
//go:generate mockery --name=Service --output=service --inpackage
type Service interface {
	RegisterUser(ctx context.Context, email, name, password string) error
	LoginUser(ctx context.Context, email, password string) (*models.User, error) // aca necesito retornar el usuario sin tomar en cuenta el retorno de la clave
	AddUserRole(ctx context.Context, userID, roleID int64) error
	RemoveUserRole(ctx context.Context, userID, roleID int64) error
	GetProducts(ctx context.Context) ([]models.Product, error)
	GetProduct(ctx context.Context, id int64) (*models.Product, error)
	AddProduct(ctx context.Context, product models.Product, userEmail string) error
}

type serv struct {
	repo repository.Repository // esta es una referencia a el repositorio
}

func New(repo repository.Repository) Service {
	return &serv{
		repo: repo,
	}
}
