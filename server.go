package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/safatanc/mesa-user-grpc/model"
	user_pb "github.com/safatanc/mesa-user-grpc/pb/user/proto"
	"github.com/safatanc/mesa-user-grpc/service/user"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	godotenv.Load()

	// Database
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&model.User{})

	// Validator
	val := validator.New(validator.WithRequiredStructEnabled())

	// gRPC Server
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "50051"
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", PORT))
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	user_pb.RegisterUserServiceServer(server, &user.UserService{DB: db, Validate: val})

	log.Printf("Server is running on %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
