package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/safatanc/mesa-user-grpc/user"
	pb "github.com/safatanc/mesa-user-grpc/user/proto"
	"google.golang.org/grpc"
)

func main() {
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "50051"
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", PORT))
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()

	pb.RegisterUserServiceServer(server, &user.User{})

	log.Printf("Server is running on %v", lis.Addr())

	if err := server.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
