package user

import (
	"fmt"

	"github.com/nidyaonur/stickverse-backend/entities"
	"github.com/nidyaonur/stickverse-backend/pkg/pb"
)

func UserAdapter(user *entities.User) *pb.UserDataResponse {
	response := &pb.UserDataResponse{
		Name:     user.Username,
		Email:    user.Email,
		Username: user.Username,
		Diamonds: uint64(user.Diamonds),
	}
	for _, city := range user.Locations {
		cityResp := &pb.City{
			Id:   uint64(city.ID),
			Name: city.Name,
			Grid: &pb.Grid{
				X:        int64(city.Grid.X),
				Y:        int64(city.Grid.Y),
				Capacity: uint64(city.Grid.Capacity),
				Occupied: uint64(city.Grid.Occupied),
				Type:     city.Grid.GridType.Name,
				Id:       uint64(city.Grid.ID),
			},
		}
		resources := &pb.Resources{}
		for _, resource := range city.Resources {
			switch resource.Resource.Name {
			case "coal":
				resources.Coal = uint64(resource.Quantity)
			case "paper":
				resources.Paper = uint64(resource.Quantity)
			case "eraser":
				resources.Eraser = uint64(resource.Quantity)
			case "ink":
				resources.Ink = uint64(resource.Quantity)
			}
		}
		buildings := []*pb.Building{}
		for _, building := range city.Structures {
			if building.Structure.Name == "townhall" {
				fmt.Println("---", building.Structure)
			}
			costsFloat64 := building.UpgradeCost()
			costs := map[string]float32{}
			for k, v := range costsFloat64 {
				costs[k] = float32(v)
			}
			fmt.Println("costs", costs)

			buildingResp := &pb.Building{
				Id:          uint64(building.ID),
				Name:        pb.StructType(pb.StructType_value[building.Structure.Name]),
				Description: building.Structure.Description["en"].(string),
				Type:        building.Structure.Name,
				Level:       uint64(building.Level),
				FloatValues: costs,
			}
			if building.Structure.Name == "townhall" {
				workerMap := make(map[string]uint64)
				for _, resource := range city.Resources {
					workerMap[resource.Resource.Name] = uint64(resource.AllocatedWorkers)
				}
				workerMap["total"] = uint64(city.Workers)
				fmt.Println("workerMap", workerMap["total"])
				buildingResp.NumberValues = workerMap
			}
			buildings = append(buildings, buildingResp)

		}

		cityResp.Resources = resources
		cityResp.Buildings = buildings
		response.Cities = append(response.Cities, cityResp)
	}
	return response

}
