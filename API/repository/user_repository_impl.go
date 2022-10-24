package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/guancang10/BookStore/API/helper"
	"github.com/guancang10/BookStore/API/models/domain"
)

type UserRepositoryImpl struct{}

func NewUserRepositoryImpl() UserRepository {
	return &UserRepositoryImpl{}
}

func (u UserRepositoryImpl) Register(ctx context.Context, tx *sql.Tx, user domain.User) {
	script := "INSERT INTO User(UserName,Password,FirstName,LastName,DOB,RoleId,AuditUsername) VALUES(?,?,?,?,?,?,?)"
	_, err := tx.ExecContext(ctx, script, user.Username, user.Password, user.FirstName, user.LastName, user.DOB, user.RoleId, user.AuditUsername)
	helper.CheckError(err)
}

func (u UserRepositoryImpl) UpdatePassword(ctx context.Context, tx *sql.Tx, user domain.User) {
	script := "UPDATE User SET password = ? WHERE username = ?"
	_, err := tx.ExecContext(ctx, script, user.Password, user.Username)
	helper.CheckError(err)
}

func (u UserRepositoryImpl) UpdateProfile(ctx context.Context, tx *sql.Tx, user domain.User) {
	script := "UPDATE User SET FirstName = ?, LastName = ?,DOB = ?, RoleId = ?, AuditUsername = ? WHERE username = ?"
	_, err := tx.ExecContext(ctx, script, user.FirstName, user.LastName, user.DOB, user.RoleId, user.AuditUsername, user.Username)
	helper.CheckError(err)
}

func (u UserRepositoryImpl) GetUser(ctx context.Context, tx *sql.Tx, username string) (domain.User, error) {
	script := "SELECT Username,Password,FirstName,LastName,DOB,RoleId,AuditUsername FROM User WHERE username = ?"
	rows, err := tx.QueryContext(ctx, script, username)
	helper.CheckError(err)
	var result domain.User

	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&result.Username, &result.Password, &result.FirstName, &result.LastName, &result.DOB, &result.RoleId, &result.AuditUsername)
		helper.CheckError(err)
		return result, nil
	} else {
		return result, errors.New("user not exists")
	}
}

func (u UserRepositoryImpl) GetAllUser(ctx context.Context, tx *sql.Tx) []domain.User {
	script := "SELECT Username,Password,FirstName,LastName,DOB,RoleId,AuditUsername FROM User"
	rows, err := tx.QueryContext(ctx, script)
	helper.CheckError(err)
	var result []domain.User
	for rows.Next() {
		var data domain.User
		err := rows.Scan(&data.Username, &data.Password, &data.FirstName, &data.LastName, &data.DOB, &data.RoleId, &data.AuditUsername)
		helper.CheckError(err)
		result = append(result, data)
	}
	return result
}
