package helper

import (
	"github.com/safatanc/mesa-user-grpc/model"
	user_service "github.com/safatanc/mesa-user-grpc/user/proto"
)

func CreateUserRequestToUser(createUserRequest *user_service.CreateUserRequest) *model.User {
	return &model.User{
		Username: createUserRequest.Username,
		FullName: createUserRequest.FullName,
		Email:    createUserRequest.Email,
		Phone:    createUserRequest.Phone,
		Password: createUserRequest.Password,
	}
}

func UserToUserResponse(user *model.User) *user_service.UserResponse {
	return &user_service.UserResponse{
		Id:        user.ID.String(),
		Username:  user.Username,
		FullName:  user.FullName,
		Email:     user.Email,
		Phone:     user.Phone,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
