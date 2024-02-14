package service

import (
	"context"
	"os"
	"testing"

	"github.com/euler-b/maxInventoryProject/encryption"
	"github.com/euler-b/maxInventoryProject/internal/entity"
	"github.com/euler-b/maxInventoryProject/internal/repository"
	"github.com/stretchr/testify/mock"
)

var repo *repository.MockRepository
var s Service

func TestMain(m *testing.M) {
	validPassword, _ := encryption.Encrypt([]byte("validPassword"))
	encryptedPassword := encryption.ToBase64(validPassword)
	u := &entity.User{Email: "william.henry.moody@my-own-personal-domain.com", Password: encryptedPassword}

	// setting de repository.Mock
	repo = &repository.MockRepository{}
	repo.On("GetUserByEmail", mock.Anything, "william.h.taft@my-own-personal-domain.com").Return(nil, nil)
	repo.On("GetUserByEmail", mock.Anything, "william.henry.moody@my-own-personal-domain.com").Return(u, nil)
	repo.On("SaveUser", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	repo.On("GetUserByRoles", mock.Anything, int64(1)).Return([]entity.USER_ROLE{{UserID: 1, RoleID: 1}}, nil) //hay que hacer (si es necesario) casting a los parametros para entregar exactamente lo que se pide.
	repo.On("SaveUserRole", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	repo.On("RemoveUserRole", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	s = New(repo)

	code := m.Run()
	os.Exit(code)
}

func TestRegisterUser(t *testing.T) {
	testCases := []struct {
		Name          string
		Email         string
		UserName      string
		Password      string
		ExpectedError error
	}{
		{
			Name:          "RegisterUser_Success",
			Email:         "william.h.taft@my-own-personal-domain.com",
			UserName:      "test",
			Password:      "validPassword",
			ExpectedError: nil,
		},
		{
			Name:          "RegisterUser_UserAlreadyExists",
			Email:         "william.henry.moody@my-own-personal-domain.com",
			UserName:      "test",
			Password:      "validPassword",
			ExpectedError: ErrUserAlreadyExists,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i] //<~ aca creamos una nueva referencia del indice de los testCases, ya que esto nos permite hacer unitests paralelos
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			repo.Mock.Test(t)
			err := s.RegisterUser(ctx, tc.Email, tc.UserName, tc.Password)
			if err != tc.ExpectedError {
				t.Errorf("Expected error %v, got %v", tc.ExpectedError, err)
			}
		})
	}
}

func TestLoginUser(t *testing.T) {
	testCases := []struct {
		Name          string
		Email         string
		Password      string
		ExpectedError error
	}{
		{
			Name:          "LoginUser_Success",
			Email:         "william.henry.moody@my-own-personal-domain.com",
			Password:      "validPassword",
			ExpectedError: nil,
		},
		{
			Name:          "LoginUser_InvalidPassword",
			Email:         "william.henry.moody@my-own-personal-domain.com",
			Password:      "Invalid_Password",
			ExpectedError: ErrInvalidCredential,
		},
	}

	var ctx context.Context
	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			repo.Mock.Test(t)

			_, err := s.LoginUser(ctx, tc.Email, tc.Password)
			if err != tc.ExpectedError {
				t.Errorf("Expected error %v, got %v", tc.ExpectedError, err)
			}
		})
	}
}

func TestAddUserRole(t *testing.T) {
	testCases := []struct {
		Name          string
		UserID        int64
		RoleID        int64
		ExpectedError error
	}{
		{
			Name:          "AddUserRole_Succes",
			UserID:        1,
			RoleID:        2,
			ExpectedError: nil,
		},
		{
			Name:          "AddUserRole_UserAlreadyHasRole",
			UserID:        1,
			RoleID:        1,
			ExpectedError: ErrRoleAlreadyAdded,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			repo.Mock.Test(t)

			err := s.AddUserRole(ctx, tc.UserID, tc.RoleID)
			if err != tc.ExpectedError {
				t.Errorf("ExpectedError %v, got %v", tc.ExpectedError, err)
			}

		})
	}
}

func TestRemoveUserRole(t *testing.T) {
	testCases := []struct {
		Name          string
		UserID        int64
		RoleID        int64
		ExpectedError error
	}{
		{
			Name:          "RemoveUserRole_Success",
			UserID:        1,
			RoleID:        1,
			ExpectedError: nil,
		},
		{
			Name:          "RemoveUserRole_UserDoesNotHaveRole",
			UserID:        1,
			RoleID:        3,
			ExpectedError: ErrRoleNotFound,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			repo.Mock.Test(t) // <~ Que hace esta linea de codigo?
			err := s.RemoveUserRole(ctx, tc.UserID, tc.RoleID)

			if err != tc.ExpectedError {
				t.Errorf("ExpectedError %v, got %v", tc.ExpectedError, err)
			}
		})

	}
}
