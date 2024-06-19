package user

import (
	"context"

	pb "github.com/safatanc/mesa-user-grpc/user/proto"
)

type User struct {
	pb.UnimplementedUserServiceServer
}

func (u *User) CreateUser(ctx context.Context, createUserRequest *pb.CreateUserRequest) (*pb.UserResponse, error) {
	data := &pb.User{
		Username: createUserRequest.Username,
		FullName: createUserRequest.FullName,
		Email:    createUserRequest.Email,
		Phone:    createUserRequest.Phone,
	}
	return &pb.UserResponse{
		Status: "success",
		Error:  nil,
		Data:   data,
	}, nil
}

func (u *User) FindAllUser(ctx context.Context, findAllUserRequest *pb.FindAllUserRequest) (*pb.UserResponses, error) {
	data := &pb.User{
		Username: "Agilistikmal",
		FullName: "Agil Ghani Istikmal",
		Email:    "agilistikmal3@gmail.com",
		Phone:    "6281346173829",
	}

	return &pb.UserResponses{
		Status: "success",
		Error:  nil,
		Data: []*pb.User{
			data,
		},
	}, nil
}

func (u *User) FindUser(ctx context.Context, findUserRequest *pb.FindUserRequest) (*pb.UserResponse, error) {
	data := &pb.User{
		Username: "Agilistikmal",
		FullName: "Agil Ghani Istikmal",
		Email:    "agilistikmal3@gmail.com",
		Phone:    "6281346173829",
	}

	return &pb.UserResponse{
		Status: "success",
		Error:  nil,
		Data:   data,
	}, nil
}
