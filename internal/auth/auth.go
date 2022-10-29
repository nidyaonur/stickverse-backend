package auth

import (
	"context"

	"github.com/nidyaonur/stickverse-backend/entities"
	"github.com/nidyaonur/stickverse-backend/internal/logic"
	"github.com/nidyaonur/stickverse-backend/pkg/pb"
	"github.com/nidyaonur/stickverse-backend/pkg/repository"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// AuthServer is the server for authentication
type AuthServer struct {
	pb.UnimplementedAuthServiceServer
	repo       repository.Repository
	jwtManager *JWTManager
	logic      *logic.Service
}

// NewAuthServer returns a new auth server
func NewAuthServer(repo repository.Repository, logicService *logic.Service, jwtManager *JWTManager) pb.AuthServiceServer {
	return &AuthServer{repo: repo, jwtManager: jwtManager, logic: logicService}
}

// Login is a unary RPC to login user
func (server *AuthServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.AuthResponse, error) {
	user, err := server.repo.LoadWithUsername(req.GetUsername())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot find user: %v", err)
	}
	if user == nil || !user.IsCorrectPassword(req.GetPassword()) {
		return nil, status.Errorf(codes.NotFound, "incorrect username/password")
	}

	token, err := server.jwtManager.Generate(user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}

	res := &pb.AuthResponse{AccessToken: token}
	return res, nil
}

func (server *AuthServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.AuthResponse, error) {
	hashedPasword, err := bcrypt.GenerateFromPassword([]byte(req.GetPassword()), bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate password hash")
	}
	location, err := server.logic.Map.PickUserLocation()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot pick user location")
	}
	user := &entities.User{
		Username:       req.GetUsername(),
		Email:          req.GetEmail(),
		HashedPassword: string(hashedPasword),
		Role:           "user",
	}
	condition := "username = ? OR email = ?"
	users := []*entities.User{}
	userInterface, err := server.repo.FindAllWithConditionValues(
		&users, nil, nil, condition, "", req.GetUsername(), req.GetEmail())
	users = *userInterface.(*[]*entities.User)
	if len(users) > 0 {
		return nil, status.Errorf(codes.AlreadyExists, "A user with this email or username already exists")
	}

	userInterface, err = server.repo.Insert(user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot insert user: %v", err)
	}
	user, err = server.logic.User.SetUserLocation(user.ID, location)
	token, err := server.jwtManager.Generate(user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}

	res := &pb.AuthResponse{AccessToken: token}
	return res, nil
}
