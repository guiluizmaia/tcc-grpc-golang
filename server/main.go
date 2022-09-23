package main

import (
	"log"
	"net"

	"github.com/guiluizmaia/tcc-grpc-golang/proto/pb"
	"github.com/guiluizmaia/tcc-grpc-golang/services"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterUserServiceServer(grpcServer, services.NewUserService)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Could not serve: %v", err)
	}
	// router := Infra.InitRest()
	// usersController := Controller.UsersController{}
	// usersController.NewUsersController(router)
	// router.Run("server:8080")
}
