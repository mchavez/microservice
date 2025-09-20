package main

import (
	"log"
	"net"
	"os"

	"microservice/internal/middleware"
	grpcdelivery "microservice/internal/user/delivery/grpc"
	userHttp "microservice/internal/user/delivery/http"
	"microservice/internal/user/repository"
	"microservice/internal/user/usecase"
	pb "microservice/proto"

	_ "microservice/docs"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// @title User Microservice API
// @version 1.0
// @description REST API for managing users
// @host localhost:8080
// @BasePath /
func main() {
	logger := setupLogger()

	repo, err := initializeRepository()
	if err != nil {
		logger.Fatalf("failed to connect to DB: %v", err)
	}

	uc := usecase.NewUserUseCase(repo, logger)

	startServers(uc, logger)
}

func setupLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.InfoLevel)
	return logger
}

func initializeRepository() (repository.UserRepository, error) {
	if os.Getenv("USE_DB") == "true" {
		return repository.NewPostgresUserRepo(
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASS"),
			os.Getenv("DB_NAME"),
		)
	}
	log.Println("Using In-Memory repository")
	return repository.NewInMemoryUserRepo(), nil
}

func startServers(uc *usecase.UserUseCase, logger *logrus.Logger) {
	go startGRPCServer(uc, logger)
	startRESTServer(uc, logger)
}

func startGRPCServer(uc *usecase.UserUseCase, logger *logrus.Logger) {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.UnaryLoggingInterceptor(logger)),
	)
	pb.RegisterUserServiceServer(grpcServer, grpcdelivery.NewUserGRPCServer(uc, logger))
	reflection.Register(grpcServer)
	log.Println("gRPC server running on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		logger.Fatalf("failed to serve gRPC: %v", err)
	}
}

func startRESTServer(uc *usecase.UserUseCase, logger *logrus.Logger) {
	router := gin.Default()
	userHttp.NewUserHandler(router, uc, logger)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	log.Println("REST server running on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("failed to run REST server: %v", err)
	}
}
