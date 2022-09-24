package services

import (
	"context"

	"github.com/guiluizmaia/tcc-grpc-golang/server/pb"
	Repository "github.com/guiluizmaia/tcc-grpc-golang/server/repository"
	Structs "github.com/guiluizmaia/tcc-grpc-golang/server/structs"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (userService *UserService) CreateUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	// fmt.Printf("Recebido %s \n", req.GetId())
	repository := Repository.NewUsersRepository()
	repository.Create(Structs.User{
		Id:          req.GetId(),
		Name:        req.GetName(),
		LastName:    req.GetLastName(),
		Age:         req.GetAge(),
		Document:    req.GetDocument(),
		Address:     req.GetAddress(),
		Nationality: req.GetNationality(),
		MotherName:  req.GetMotherName(),
		FatherName:  req.GetFatherName(),
		Gender:      req.GetGender(),
		Birthday:    req.GetBirthday(),
		Email:       req.GetEmail(),
	})

	return &pb.User{
		Id:          req.GetId(),
		Name:        req.GetName(),
		LastName:    req.GetLastName(),
		Age:         req.GetAge(),
		Document:    req.GetDocument(),
		Address:     req.GetAddress(),
		Nationality: req.GetNationality(),
		MotherName:  req.GetMotherName(),
		FatherName:  req.GetFatherName(),
		Gender:      req.GetGender(),
		Birthday:    req.GetBirthday(),
		Email:       req.GetEmail(),
	}, nil
}
