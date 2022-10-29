package maplist

import (
	"github.com/nidyaonur/stickverse-backend/entities"
	pb "github.com/nidyaonur/stickverse-backend/pkg/pb"
)

func TableAdapter(grid *entities.Grid, locations []*entities.Location, userID uint) *pb.TableResponse {
	response := &pb.TableResponse{
		Grid: &pb.Grid{
			X:        int64(grid.X),
			Y:        int64(grid.X),
			Capacity: uint64(grid.Capacity),
			Occupied: uint64(grid.Occupied),
			Type:     grid.GridType.Name,
		},
		Locations: []*pb.Location{},
	}
	for _, location := range locations {
		respLocation := &pb.Location{
			Name:        location.Name,
			Level:       uint64(location.Level),
			Description: location.LocationType.Description["en"].(string),
		}
		if location.User != nil {
			respLocation.Username = &location.User.Username
			if location.User.ID == userID {
				respLocation.Owned = true
			}
		}
		response.Locations = append(response.Locations, respLocation)
	}
	return response
}

func LibraryAdapter(grid []*entities.Grid) *pb.LibraryResponse {
	response := &pb.LibraryResponse{
		Grids: []*pb.Grid{},
	}
	for _, g := range grid {
		response.Grids = append(response.Grids, &pb.Grid{
			X:        int64(g.X),
			Y:        int64(g.Y),
			Capacity: uint64(g.Capacity),
			Occupied: uint64(g.Occupied),
			Type:     g.GridType.Name,
			Id:       uint64(g.ID),
		})
	}
	return response
}
