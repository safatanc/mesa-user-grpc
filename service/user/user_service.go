package user

import (
	"context"

	"github.com/safatanc/mesa-user-grpc/helper"
	"github.com/safatanc/mesa-user-grpc/model"
	user_pb "github.com/safatanc/mesa-user-grpc/pb/user/proto"
	"gorm.io/gorm"
)

type UserService struct {
	user_pb.UnimplementedUserServiceServer
	DB *gorm.DB
}

func (u *UserService) CreateUser(ctx context.Context, createUserRequest *user_pb.CreateUserRequest) (*user_pb.UserResponse, error) {
	user := helper.CreateUserRequestToUser(createUserRequest)
	result := u.DB.Create(&user)

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
