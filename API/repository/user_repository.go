package repository

import (
	"context"
	"database/sql"
	"github.com/guancang10/BookStore/API/models/domain"
)

type UserRepository interface {
	Register(ctx context.Context, tx *sql.Tx, user domain.User)
	UpdatePassword(ctx context.Context, tx *sql.Tx, user domain.User)
	UpdateProfile(ctx context.Context, tx *sql.Tx, user domain.User)
	GetUser(ctx context.Context, tx *sql.Tx, username string) (domain.User, error)
	GetAllUser(ctx context.Context, tx *sql.Tx) []domain.User
}
