package maplist

import (
	"context"
	"fmt"
	"time"

	"github.com/nidyaonur/stickverse-backend/entities"
	"github.com/nidyaonur/stickverse-backend/internal/auth"
	"github.com/nidyaonur/stickverse-backend/internal/logic"
	pb "github.com/nidyaonur/stickverse-backend/pkg/pb"
	"github.com/nidyaonur/stickverse-backend/pkg/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// MapListServer is the server for authentication
type MaplistServer struct {
	pb.UnimplementedMaplistServiceServer
	repo         repository.Repository
	LogicService *logic.Service
	jwtManager   *auth.JWTManager
}

// NewMapListServer returns a new auth server
func NewMaplistServer(repo repository.Repository, logicService *logic.Service, jwtManager *auth.JWTManager) *MaplistServer {
	return &MaplistServer{repo: repo, LogicService: logicService, jwtManager: jwtManager}
}

func (server *MaplistServer) GetTable(ctx context.Context, req *pb.TableRequest) (*pb.TableResponse, error) {
	requestTime := time.Now()
	claims, err := server.jwtManager.VerifyFromCtx(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}
	grid := &entities.Grid{}
	server.repo.FindFirstWithCondition(grid, "id = ?", req.GetGridId())

	locations := []*entities.Location{}
	_, err = server.repo.FindAllWithConditionValues(&locations, nil, nil, "grid_id = ?", "", req.GetGridId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get locations: %v", err)
	}
	usersOnGrid := []*entities.User{}
	for _, location := range locations {
		if location.UserID == nil {
			continue
		}
		user := &entities.User{}
		user.ID = *location.UserID
		usersOnGrid = append(usersOnGrid, user)
	}
	server.repo.LoadUsers(usersOnGrid)
	for _, user := range usersOnGrid {
		server.LogicService.User.UpdateUserResources(user, requestTime)
	}
	response := TableAdapter(grid, locations, claims.UserID)
	return response, nil
}

func (server *MaplistServer) GetLibrary(ctx context.Context, req *pb.LibraryRequest) (*pb.LibraryResponse, error) {
	_, err := server.jwtManager.VerifyFromCtx(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}
	fmt.Println("GetLibrary")

	grids := []*entities.Grid{}
	_, err = server.repo.FindAllWithConditionValues(&grids,
		nil, nil,
		"x >= ? AND y >= ? AND x <= ? AND y <= ?", "",
		req.GetMinX(), req.GetMinY(), req.GetMaxX(), req.GetMaxY())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get locations: %v", err)
	}
	fmt.Println("grids", grids)

	response := LibraryAdapter(grids)
	return response, nil
}
