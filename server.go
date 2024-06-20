package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/safatanc/mesa-user-grpc/model"
	"github.com/safatanc/mesa-user-grpc/user"
	user_service "github.com/safatanc/mesa-user-grpc/user/proto"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	godotenv.Load()

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "50051"
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", PORT))
	if err != nil {
		log.Fatal(err)
	}

	// Database
	dsn := os.Getenv("DSN")
	log.Println("DSN", dsn)
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&model.User{})

	// gRPC
	server := grpc.NewServer()
	user_service.RegisterUserServiceServer(server, &user.UserService{DB: db})

	log.Printf("Server is running on %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
