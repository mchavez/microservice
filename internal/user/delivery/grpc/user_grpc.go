package grpc

// package grpcdelivery

import (
	"context"
	"microservice/internal/user/entity"
	"microservice/internal/user/usecase"
	pb "microservice/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserGRPCServer struct {
	pb.UnimplementedUserServiceServer
	uc *usecase.UserUseCase
}

func NewUserGRPCServer(uc *usecase.UserUseCase) *UserGRPCServer {
	return &UserGRPCServer{uc: uc}
}

func (s *UserGRPCServer) GetUsers(ctx context.Context, _ *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	users, _ := s.uc.GetUsers()
	res := make([]*pb.User, 0, len(users))
	for _, u := range users {
		res = append(res, &pb.User{Id: u.ID, Name: u.Name})
	}
	return &pb.ListUsersResponse{Users: res}, nil
}

func (s *UserGRPCServer) AddUser(ctx context.Context, user *pb.User) (*pb.User, error) {
	u, _ := s.uc.CreateUser(&entity.User{Name: user.Name})
	return &pb.User{Id: u.ID, Name: u.Name}, nil
}

// GetUserByID retrieves a user by ID
func (s *UserGRPCServer) GetUserByID(ctx context.Context, req *pb.GetUserByIDRequest) (*pb.GetUserByIDResponse, error) {
	user, err := s.uc.GetUserByID(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}
	return &pb.GetUserByIDResponse{
		User: &pb.User{Id: user.ID, Name: user.Name},
	}, nil
}

// GetUsersByName retrieves users filtered by name
func (s *UserGRPCServer) GetUsersByName(ctx context.Context, req *pb.GetUsersByNameRequest) (*pb.GetUsersByNameResponse, error) {
	users, err := s.uc.GetUsersByName(req.Name)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "no users found")
	}

	var pbUsers []*pb.User
	for _, u := range users {
		pbUsers = append(pbUsers, &pb.User{Id: u.ID, Name: u.Name})
	}

	return &pb.GetUsersByNameResponse{Users: pbUsers}, nil
}
