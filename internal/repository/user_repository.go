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
