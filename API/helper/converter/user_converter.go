package converter

import (
	"github.com/guancang10/BookStore/API/helper"
	"github.com/guancang10/BookStore/API/models/domain"
	"github.com/guancang10/BookStore/API/models/web/request"
	"github.com/guancang10/BookStore/API/models/web/response"
	"golang.org/x/crypto/bcrypt"
)

func FromUserRegisterToUser(user request.UserRegisterRequest) domain.User {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	helper.CheckError(err)

	//convert time
	dob := ConvertDateFromString(user.DOB)
	helper.CheckError(err)
	return domain.User{Username: user.Username,
		Password:      hashPassword,
		DOB:           dob,
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		RoleId:        user.RoleId,
		AuditUsername: user.AuditUsername,
	}
}

func FromUserChangePassToUser(username string, password string) domain.User {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	helper.CheckError(err)
	return domain.User{
		Username: username,
		Password: hashPassword,
	}
}

func FromUpdateUserToUser(user request.UserUpdateRequest) domain.User {
	dob := ConvertDateFromString(user.DOB)
	return domain.User{
		Username:      user.Username,
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		RoleId:        user.RoleId,
		DOB:           dob,
		AuditUsername: user.AuditUsername,
	}
}

func FromUserToUserResponse(user domain.User) response.UserGetResponse {
	return response.UserGetResponse{
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		RoleId:    user.RoleId,
		DOB:       ConvertDateFromTime(user.DOB),
	}
}

func FromUserArrToUserResponse(user []domain.User) []response.UserGetResponse {
	var result []response.UserGetResponse
	for _, v := range user {
		data := response.UserGetResponse{
			Username:  v.Username,
			FirstName: v.FirstName,
			LastName:  v.LastName,
			RoleId:    v.RoleId,
			DOB:       ConvertDateFromTime(v.DOB),
		}
		result = append(result, data)
	}
	return result
}
