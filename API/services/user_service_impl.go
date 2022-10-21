package services

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/guancang10/BookStore/API/helper"
	"github.com/guancang10/BookStore/API/helper/converter"
	"github.com/guancang10/BookStore/API/models/web/request"
	"github.com/guancang10/BookStore/API/models/web/response"
	"github.com/guancang10/BookStore/API/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	Db             *sql.DB
	Validator      *validator.Validate
	UserRepository repository.UserRepository
}

func NewUserServiceImpl(db *sql.DB, validator *validator.Validate, userRepository repository.UserRepository) UserServices {
	return &UserServiceImpl{Db: db, Validator: validator, UserRepository: userRepository}
}

func (u UserServiceImpl) Register(ctx context.Context, userRequest request.UserRegisterRequest) {
	err := u.Validator.Struct(userRequest)
	helper.CheckError(err)

	tx, err := u.Db.Begin()
	helper.CheckError(err)
	defer helper.CheckErrorTx(tx)

	user := converter.FromUserRegisterToUser(userRequest)
	exists, err := u.UserRepository.GetUser(ctx, tx, user.Username)

	if exists.Username != "" {
		panic(errors.New("username already exists"))
	} else {
		u.UserRepository.Register(ctx, tx, user)
	}

}

func (u UserServiceImpl) CheckPassword(ctx context.Context, user request.UserLoginRequest) response.UserGetResponse {
	err := u.Validator.Struct(user)
	helper.CheckError(err)

	tx, err := u.Db.Begin()
	helper.CheckError(err)

	exist, err := u.UserRepository.GetUser(ctx, tx, user.Username)
	if exist.Username == "" {
		panic(errors.New("username not exists"))
	} else {
		err := bcrypt.CompareHashAndPassword(exist.Password, []byte(user.Password))
		if err != nil {
			panic(errors.New("Username or password not match"))
		} else {
			return converter.FromUserToUserResponse(exist)
		}
	}
}

func (u UserServiceImpl) UpdateProfile(ctx context.Context, userRequest request.UserUpdateRequest) {
	err := u.Validator.Struct(userRequest)
	helper.CheckError(err)

	tx, err := u.Db.Begin()
	helper.CheckError(err)
	defer helper.CheckErrorTx(tx)
	user := converter.FromUpdateUserToUser(userRequest)
	u.UserRepository.UpdateProfile(ctx, tx, user)
}

func (u UserServiceImpl) GetUser(ctx context.Context, username string) response.UserGetResponse {
	tx, err := u.Db.Begin()
	helper.CheckError(err)

	user, err := u.UserRepository.GetUser(ctx, tx, username)
	helper.CheckError(err)
	return converter.FromUserToUserResponse(user)

}

func (u UserServiceImpl) GetAllUser(ctx context.Context) []response.UserGetResponse {
	tx, err := u.Db.Begin()
	helper.CheckError(err)

	result := u.UserRepository.GetAllUser(ctx, tx)

	return converter.FromUserArrToUserResponse(result)
}

func (u UserServiceImpl) UpdatePassword(ctx context.Context, userRequest request.UserChangePassword) {
	err := u.Validator.Struct(userRequest)
	helper.CheckError(err)
	tx, err := u.Db.Begin()
	helper.CheckError(err)
	defer helper.CheckErrorTx(tx)

	//Check is the username and password match

	u.CheckPassword(ctx, request.UserLoginRequest{Username: userRequest.Username, Password: userRequest.Password})

	user := converter.FromUserChangePassToUser(userRequest.Username, userRequest.NewPassword)
	u.UserRepository.UpdatePassword(ctx, tx, user)
}
