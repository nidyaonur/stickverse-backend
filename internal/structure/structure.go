package structure

import (
	"context"

	"github.com/nidyaonur/stickverse-backend/entities"
	"github.com/nidyaonur/stickverse-backend/internal/auth"
	"github.com/nidyaonur/stickverse-backend/internal/logic"
	"github.com/nidyaonur/stickverse-backend/pkg/pb"
	"github.com/nidyaonur/stickverse-backend/pkg/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// AuthServer is the server for authentication
type StructureServer struct {
	pb.UnimplementedStructureServiceServer
	repo       repository.Repository
	Logic      *logic.Service
	jwtManager *auth.JWTManager
}

// NewAuthServer returns a new auth server
func NewStructureServer(repo repository.Repository, logicService *logic.Service, jwtManager *auth.JWTManager) *StructureServer {
	return &StructureServer{repo: repo, Logic: logicService, jwtManager: jwtManager}
}

// Login is a unary RPC to login structure
func (server *StructureServer) StructureOperation(ctx context.Context, req *pb.StructureRequest) (*pb.StructureResponse, error) {
	claims, err := server.jwtManager.VerifyFromCtx(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	user := &entities.User{}
	user.ID = claims.UserID
	server.repo.LoadUser(user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot find user : %v", err)
	}
	switch req.Operation {
	case pb.StructOperation_upgrade:
		err = server.Logic.Structure.Upgrade(user, *req.Id, req.Type.String())
	case pb.StructOperation_redistributeWorkers:
		err = server.Logic.Structure.RedistributeWorkers(user, *req.Id, req.IntParams)
	}
	if err != nil {
		return &pb.StructureResponse{Success: false}, nil
	}
	return &pb.StructureResponse{Success: true}, err

}
