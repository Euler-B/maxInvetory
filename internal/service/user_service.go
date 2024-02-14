package service

import (
	"context"
	"errors"

	"github.com/euler-b/maxInventoryProject/encryption"
	"github.com/euler-b/maxInventoryProject/internal/models"
)

var (
	ErrUserAlreadyExists = errors.New("el usuario ya existe")
	ErrInvalidCredential = errors.New("la contraseña es incorrecta")
	ErrRoleAlreadyAdded  = errors.New("el rol ya fue agregado")
	ErrRoleNotFound      = errors.New("el rol ya fue eliminado")
)

func (s *serv) RegisterUser(ctx context.Context, email, name, passsword string) error {
	// primero debemos revisar si el usuario ya existe
	u, _ := s.repo.GetUserByEmail(ctx, email)
	if u != nil {
		return ErrUserAlreadyExists
	}

	// Si el usuario no existe, este se crea y su contraseña se encripta
	bb, err := encryption.Encrypt([]byte(passsword))
	if err != nil {
		return err
	}
	pass := encryption.ToBase64(bb)

	return s.repo.SaveUser(ctx, email, name, pass)
}

func (s *serv) LoginUser(ctx context.Context, email, password string) (*models.User, error) {
	u, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	bb, err := encryption.FromBase64(u.Password) //<~ Base64 se guarda la contraseña encriptada, aqui la transforma a arreglo de bytes
	if err != nil {
		return nil, err
	}
	decryptedPassword, err := encryption.Decrypt(bb) //<~ Decrypt tambien me arroja un arreglo de bytes, pero este arreglo es el que me representa la clave guardada
	if err != nil {
		return nil, err
	}

	if string(decryptedPassword) != password { //<~ transformamos el arreglo de bytes a string
		return nil, ErrInvalidCredential
	}

	return &models.User{
		ID:    u.ID,
		Email: u.Email,
		Name:  u.Name,
	}, nil
}

func (s *serv) AddUserRole(ctx context.Context, userID, roleID int64) error {
	roles, err := s.repo.GetUserByRoles(ctx, userID)
	if err != nil {
		return err
	}

	for _, r := range roles {
		if r.RoleID == roleID {
			return ErrRoleAlreadyAdded
		}
	}

	return s.repo.SaveUserRole(ctx, userID, roleID)
}

func (s *serv) RemoveUserRole(ctx context.Context, userID, roleID int64) error {
	roles, err := s.repo.GetUserByRoles(ctx, userID)
	if err != nil {
		return err
	}

	roleFound := false
	for _, r := range roles {
		if r.RoleID == roleID {
			roleFound = true
			break
		}
	}

	if !roleFound {
		return ErrRoleNotFound
	}

	return s.repo.RemoveUserRole(ctx, userID, roleID)
}
