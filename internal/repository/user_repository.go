package repository

import (
	"context"

	"github.com/euler-b/maxInventoryProject/internal/entity"
)

const (
	queryInsertUser = `
		insert into USERS (email, name, password)
		values (?, ?, ?);
	`
	queryGetUserByEmail = `
		select
			id,
			email,
			name,
			password
		from USERS
		where email = ?;
	`
	// Aca se muestra otra manera en la que la librreria lee los campos de la db
	queryInsertUserRole = `
		insert into USER_ROLES (user_id, role_id) values (:user_id, :role_id); 
	`
	queryDeleteUserRole = `
		delete from USER_ROLES where user_id = :user_id and role_id = :role_id;
	`
)

func (r *repo) SaveUser(ctx context.Context, email, name, password string) error {
	_, err := r.db.ExecContext(ctx, queryInsertUser, email, name, password)
	return err
}

func (r *repo) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	u := &entity.User{}
	err := r.db.GetContext(ctx, u, queryGetUserByEmail, email)
	if err != nil {
		return nil, err
	}
	return u, nil // clean better than clever
}

func (r *repo) SaveUserRole(ctx context.Context, userID, roleID int64) error {
	// Esta es la interfaz que me requiere el parametro de la interfaz
	data := entity.USER_ROLE{
		UserID: userID,
		RoleID: roleID,
	}
	// En este caso uso Name execContex porque estoy usando los nombre de las sentencias sql en la query
	_, err := r.db.NamedExecContext(ctx, queryInsertUserRole, data)
	return err
}

func (r *repo) RemoveUserRole(ctx context.Context, userID, roleID int64) error {
	data := entity.USER_ROLE{
		UserID: userID,
		RoleID: roleID,
	}

	_, err := r.db.NamedExecContext(ctx, queryDeleteUserRole, data)
	return err
}

func (r *repo) GetUserByRoles(ctx context.Context, userID int64) ([]entity.USER_ROLE, error) {
	roles := []entity.USER_ROLE{}
	err := r.db.SelectContext(ctx, &roles, "select user_id, role_id from USER_ROLES where user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	return roles, nil
}
