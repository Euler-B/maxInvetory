package service

import (
	"os"
	"testing"

	"github.com/euler-b/maxInventoryProject/encryption"
	"github.com/euler-b/maxInventoryProject/internal/entity"
	"github.com/euler-b/maxInventoryProject/internal/repository"
	mock "github.com/stretchr/testify/mock"
)

var repo *repository.MockRepository
var s Service

func TestMain(m *testing.M) {
	validPassword, _ := encryption.Encrypt([]byte("validPassword"))
	encryptedPassword := encryption.ToBase64(validPassword)
	u := &entity.User{Email: "william.henry.moody@my-own-personal-domain.com", Password: encryptedPassword}
	adminUser := &entity.User{ID: 1, Email: "admin@email.com", Password: encryptedPassword}
	customerUser := &entity.User{ID: 2, Email: "customer@email.com", Password: encryptedPassword}

	// setting de repository.Mock
	repo = &repository.MockRepository{}
	repo.On("GetUserByEmail", mock.Anything, "william.h.taft@my-own-personal-domain.com").Return(nil, nil)
	repo.On("GetUserByEmail", mock.Anything, "william.henry.moody@my-own-personal-domain.com").Return(u, nil)
	repo.On("GetUserByEmail", mock.Anything, "admin@email.com").Return(adminUser, nil)
	repo.On("GetUserByEmail", mock.Anything, "customer@email.com").Return(customerUser, nil)

	repo.On("SaveUser", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	repo.On("GetUserByRoles", mock.Anything, int64(1)).Return([]entity.USER_ROLE{{UserID: 1, RoleID: 1}}, nil) //hay que hacer (si es necesario) casting a los parametros para entregar exactamente lo que se pide.
	repo.On("GetUserByRoles", mock.Anything, int64(2)).Return([]entity.USER_ROLE{{UserID: 2, RoleID: 3}}, nil)
	repo.On("SaveUserRole", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	repo.On("RemoveUserRole", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	repo.On("SaveProduct", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	s = New(repo)

	code := m.Run()
	os.Exit(code)
}
