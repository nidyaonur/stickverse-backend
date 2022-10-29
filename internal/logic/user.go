package logic

import (
	"fmt"
	"time"

	"github.com/nidyaonur/stickverse-backend/entities"
	"github.com/nidyaonur/stickverse-backend/pkg/repository"
)

type UserService struct {
	repository repository.Repository
}

func NewUserService(repository repository.Repository) *UserService {
	return &UserService{repository}
}

func (s *UserService) SetUserLocation(userId uint, location *entities.Location) (*entities.User, error) {
	grid := location.Grid
	grid.Occupied += 1
	_, err := s.repository.Update(grid)
	if err != nil {
		return nil, err
	}
	locationType := entities.LocationType{Name: "town"}
	_, err = s.repository.FindFirstWithCondition(&locationType, "name = ?", "town")
	if err != nil {
		return nil, err
	}
	fmt.Println("locationType", locationType)
	location.LocationTypeID = locationType.ID
	location.LocationType = locationType
	location.UserID = &userId
	location.Level = 1
	_, err = s.repository.Update(location)
	/*
		structure := entities.Structure{}
		_, err = s.repository.FindFirstWithCondition(&structure, "name = ?", "townhall")
		structureBuilt := entities.StructureBuilt{}
		_, err = s.repository.FindFirstWithCondition(&structureBuilt, "location_id = ? AND structure_id = ?", location.ID, structure.ID)
		if err != nil {
			return nil, err
		}
		structureBuilt.Level = 1
		_, err = s.repository.Update(structureBuilt)
		if err != nil {
			return nil, err
		}
	*/
	user := &entities.User{}
	_, err = s.repository.FindFirstWithCondition(&user, "id = ?", userId)
	return user, err
}

func (s *UserService) UpdateUserResources(user *entities.User, requestTime time.Time) error {
	s.repository.LoadUser(user)
	locationResources := []*entities.LocationResource{}
	builts := []*entities.StructureBuilt{}
	for _, city := range user.Locations {
		for r := range city.Resources {
			resource := city.Resources[r]
			timeInterval := requestTime.Sub(resource.UpdatedAt)
			timeMinutes := timeInterval.Minutes()
			resource.Quantity += timeMinutes * resource.Multiplier * float64(resource.AllocatedWorkers)
			locationResources = append(locationResources, &resource)
		}
		//check built structures if upgrade finished
		for i := range city.Structures {
			built := city.Structures[i]
			fmt.Println("built", built.ID, built.UpgradeOngoing)
			if !built.UpgradeOngoing {
				continue

			}
			if built.UpgradeEndedAt.Before(requestTime) {

				built.Level += 1
				built.UpgradeOngoing = false
				builts = append(builts, &built)
			}
		}
	}
	_, err := s.repository.Update(locationResources)
	if err != nil {
		return err
	}
	_, err = s.repository.Update(&builts)

	err = s.repository.LoadUser(user)
	return err
}
