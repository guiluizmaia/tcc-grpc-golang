package services

import (
	"context"

	"github.com/guiluizmaia/tcc-grpc-golang/proto/pb"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (*UserService) CreateUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	return &pb.User{
		id:          req.GetId(),
		name:        req.GetName(),
		lastName:    req.GetLastName(),
		age:         req.GetAge(),
		document:    req.GetDocument(),
		address:     req.GetAddress(),
		nationality: req.GetNationality(),
		motherName:  req.GetMotherName(),
		fatherName:  req.GetFatherName(),
		gender:      req.GetGender(),
		birthday:    req.GetBirthday(),
		email:       req.GetEmail(),
	}, nil
}
