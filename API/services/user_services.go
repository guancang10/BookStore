package services

import (
	"context"
	"github.com/guancang10/BookStore/API/models/web/request"
	"github.com/guancang10/BookStore/API/models/web/response"
)

type UserServices interface {
	Register(ctx context.Context, userRequest request.UserRegisterRequest)
	CheckPassword(ctx context.Context, user request.UserLoginRequest) response.UserGetResponse
	UpdateProfile(ctx context.Context, userRequest request.UserUpdateRequest)
	GetUser(ctx context.Context, username string) response.UserGetResponse
	GetAllUser(ctx context.Context) []response.UserGetResponse
	UpdatePassword(ctx context.Context, userRequest request.UserChangePassword)
}
