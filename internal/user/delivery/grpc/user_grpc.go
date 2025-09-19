package grpc

// package grpcdelivery

import (
	"context"
	"microservice/internal/user/entity"
	"microservice/internal/user/usecase"
	pb "microservice/proto"
)

type UserGRPCServer struct {
	pb.UnimplementedUserServiceServer
	uc *usecase.UserUseCase
}

func NewUserGRPCServer(uc *usecase.UserUseCase) *UserGRPCServer {
	return &UserGRPCServer{uc: uc}
}

func (s *UserGRPCServer) GetUsers(ctx context.Context, _ *pb.Empty) (*pb.UserList, error) {
	users, _ := s.uc.GetUsers()
	res := make([]*pb.User, 0, len(users))
	for _, u := range users {
		res = append(res, &pb.User{Id: u.ID, Name: u.Name})
	}
	return &pb.UserList{Users: res}, nil
}

func (s *UserGRPCServer) AddUser(ctx context.Context, user *pb.User) (*pb.User, error) {
	u, _ := s.uc.CreateUser(&entity.User{Name: user.Name})
	return &pb.User{Id: u.ID, Name: u.Name}, nil
}
