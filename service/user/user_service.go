package user

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/safatanc/mesa-user-grpc/helper"
	"github.com/safatanc/mesa-user-grpc/model"
	user_pb "github.com/safatanc/mesa-user-grpc/pb/user/proto"
	"gorm.io/gorm"
)

type UserService struct {
	user_pb.UnimplementedUserServiceServer
	DB       *gorm.DB
	Validate *validator.Validate
}

func (u *UserService) CreateUser(ctx context.Context, createUserRequest *user_pb.CreateUserRequest) (*user_pb.UserResponse, error) {
	user := helper.CreateUserRequestToUser(createUserRequest)
	err := u.Validate.Struct(user)
	if err != nil {
		return nil, err
	}

	result := u.DB.Create(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return helper.UserToUserResponse(user), nil
}

func (u *UserService) UpdateUser(ctx context.Context, updateUserRequest *user_pb.UpdateUserRequest) (*user_pb.UserResponse, error) {
	var user *model.User
	result := u.DB.First(&user, "username = ?", updateUserRequest.Username)
	if result.Error != nil {
		return nil, result.Error
	}

	if updateUserRequest.Data.Username != nil {
		user.Username = *updateUserRequest.Data.Username
	}
	if updateUserRequest.Data.FullName != nil {
		user.FullName = *updateUserRequest.Data.FullName
	}
	if updateUserRequest.Data.Email != nil {
		user.Email = *updateUserRequest.Data.Email
	}
	if updateUserRequest.Data.Phone != nil {
		user.Phone = *updateUserRequest.Data.Phone
	}
	if updateUserRequest.Data.Password != nil {
		user.Password = *updateUserRequest.Data.Password
	}

	err := u.Validate.Struct(user)
	if err != nil {
		return nil, err
	}

	result = u.DB.Updates(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return helper.UserToUserResponse(user), nil
}

func (u *UserService) DeleteUser(ctx context.Context, deleteUserRequest *user_pb.DeleteUserRequest) (*user_pb.UserResponse, error) {
	var user *model.User
	result := u.DB.First(&user, "username = ?", deleteUserRequest.Username)
	if result.Error != nil {
		return nil, result.Error
	}

	result = u.DB.Delete(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return helper.UserToUserResponse(user), nil
}

func (u *UserService) FindAllUser(ctx context.Context, findAllUserRequest *user_pb.FindAllUserRequest) (*user_pb.UserResponses, error) {
	var users []*model.User
	u.DB.Find(&users)

	var userResponses []*user_pb.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, helper.UserToUserResponse(user))
	}

	return &user_pb.UserResponses{
		Users: userResponses,
	}, nil
}

func (u *UserService) FindUser(ctx context.Context, findUserRequest *user_pb.FindUserRequest) (*user_pb.UserResponse, error) {
	var user *model.User
	result := u.DB.First(&user, "username = ?", findUserRequest.Username)

	if result.Error != nil {
		return nil, result.Error
	}

	return helper.UserToUserResponse(user), nil
}
