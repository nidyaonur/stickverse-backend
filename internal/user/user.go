package user

import (
	"context"
	"time"

	"github.com/nidyaonur/stickverse-backend/entities"
	"github.com/nidyaonur/stickverse-backend/internal/auth"
	"github.com/nidyaonur/stickverse-backend/internal/logic"
	"github.com/nidyaonur/stickverse-backend/pkg/pb"
	"github.com/nidyaonur/stickverse-backend/pkg/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// AuthServer is the server for authentication
type UserServer struct {
	pb.UnimplementedUserServiceServer
	repo         repository.Repository
	LogicService *logic.Service
	jwtManager   *auth.JWTManager
}

// NewAuthServer returns a new auth server
func NewUserServer(repo repository.Repository, logicService *logic.Service, jwtManager *auth.JWTManager) *UserServer {
	return &UserServer{repo: repo, LogicService: logicService, jwtManager: jwtManager}
}

// Login is a unary RPC to login user
func (server *UserServer) GetUserData(ctx context.Context, req *pb.UserDataRequest) (*pb.UserDataResponse, error) {
	requestTime := time.Now()
	claims, err := server.jwtManager.VerifyFromCtx(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	user := &entities.User{}
	user.ID = claims.UserID
	server.repo.LoadUser(user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot find user: %v", err)
	}
	server.LogicService.User.UpdateUserResources(user, requestTime)

	response := UserAdapter(user)
	return response, nil
}
