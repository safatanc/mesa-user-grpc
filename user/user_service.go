package user

import (
	"context"

	"github.com/safatanc/mesa-user-grpc/helper"
	"github.com/safatanc/mesa-user-grpc/model"
	user_service "github.com/safatanc/mesa-user-grpc/user/proto"
	"gorm.io/gorm"
)

type UserService struct {
	user_service.UnimplementedUserServiceServer
	DB *gorm.DB
}

func (u *UserService) CreateUser(ctx context.Context, createUserRequest *user_service.CreateUserRequest) (*user_service.Response, error) {
	user := helper.CreateUserRequestToUser(createUserRequest)
	result := u.DB.Create(&user)

	if result.Error != nil {
		return &user_service.Response{
			Status: "error",
			Error: &user_service.Error{
				Message: result.Error.Error(),
			},
			Data: nil,
		}, nil
	}

	return &user_service.Response{
		Status: "success",
		Error:  nil,
		Data:   helper.UserToUserResponse(user),
	}, nil
}

func (u *UserService) FindAllUser(ctx context.Context, findAllUserRequest *user_service.FindAllUserRequest) (*user_service.Responses, error) {
	var users []*model.User
	u.DB.Find(&users)

	var userResponses []*user_service.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, helper.UserToUserResponse(user))
	}

	return &user_service.Responses{
		Status: "success",
		Error:  nil,
		Data:   userResponses,
	}, nil
}

func (u *UserService) FindUser(ctx context.Context, findUserRequest *user_service.FindUserRequest) (*user_service.Response, error) {
	var user *model.User
	result := u.DB.First(&user, "username = ?", findUserRequest.Username)

	if result.Error != nil {
		return &user_service.Response{
			Status: "error",
			Error: &user_service.Error{
				Message: result.Error.Error(),
			},
			Data: nil,
		}, nil
	}

	return &user_service.Response{
		Status: "success",
		Error:  nil,
		Data:   helper.UserToUserResponse(user),
	}, nil
}
