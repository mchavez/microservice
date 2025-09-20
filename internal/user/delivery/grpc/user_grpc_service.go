package grpc

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
	users, err := s.uc.GetUsers()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to retrieve users: %v", err)
	}
	return convertToListUsersResponse(users), nil
}

func (s *UserGRPCServer) AddUser(ctx context.Context, user *pb.User) (*pb.User, error) {
	u, err := s.uc.CreateUser(&entity.User{Name: user.Name})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create user: %v", err)
	}
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
	return convertToGetUsersByNameResponse(users), nil
}

// Helper function to convert users to ListUsersResponse
func convertToListUsersResponse(users []*entity.User) *pb.ListUsersResponse {
	res := make([]*pb.User, len(users))
	for i, u := range users {
		res[i] = &pb.User{Id: u.ID, Name: u.Name}
	}
	return &pb.ListUsersResponse{Users: res}
}

// Helper function to convert users to GetUsersByNameResponse
func convertToGetUsersByNameResponse(users []*entity.User) *pb.GetUsersByNameResponse {
	res := make([]*pb.User, len(users))
	for i, u := range users {
		res[i] = &pb.User{Id: u.ID, Name: u.Name}
	}
	return &pb.GetUsersByNameResponse{Users: res}
}
