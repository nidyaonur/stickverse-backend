package logic

import (
	"errors"
	"fmt"
	"time"

	"github.com/nidyaonur/stickverse-backend/entities"
	"github.com/nidyaonur/stickverse-backend/pkg/repository"
)

type StructureService struct {
	repository repository.Repository
}

func NewStructureService(repository repository.Repository) *StructureService {
	return &StructureService{repository}
}

func (s *StructureService) Upgrade(user *entities.User, locStructID uint64, structType string) error {
	locStruct := &entities.StructureBuilt{}
	preloads := []string{"Structure", "Structure.Prerequisites", "Location",
		"Location.Structures", "Location.Resources", "Location.Resources.Resource", "Location.User"}
	s.repository.FindFirstWithCustomPreloads(locStruct, preloads, "id = ?", locStructID)
	if locStruct.UpgradeOngoing {
		fmt.Println("upgrade ongoing", locStruct.ID)
		return errors.New("upgrade already ongoing")
	}
	resources := map[string]float64{}
	for _, res := range locStruct.Location.Resources {
		fmt.Println("res", res.Resource.Name, res.Quantity)
		resources[res.Resource.Name] = float64(res.Quantity)
	}
	fmt.Println("resources", resources)
	costMap := locStruct.UpgradeCost()

	for res, cost := range costMap {
		fmt.Println("cost", res, cost, resources[res])
		if resources[res] < cost {
			return errors.New("not enough resources")
		}
	}
	fmt.Println("upgrade cost", costMap)
	for resName, cost := range costMap {
		for _, res := range locStruct.Location.Resources {
			if res.Resource.Name == resName {
				res.Quantity -= cost
				s.repository.Update(&res)
			}
		}
	}
	upgradeTime := time.Now().Add(time.Duration(locStruct.Level) * time.Minute)
	locStruct.UpgradeOngoing = true
	locStruct.UpgradeEndedAt = &upgradeTime
	s.repository.Update(locStruct)
	return nil
}

func (s *StructureService) RedistributeWorkers(user *entities.User, locStructID uint64, workers map[string]int64) error {
	fmt.Println("redistribute workers", locStructID, workers)
	locStruct := &entities.StructureBuilt{}
	preloads := []string{"Structure", "Structure.Prerequisites", "Location",
		"Location.Structures", "Location.Resources", "Location.Resources.Resource", "Location.User"}
	s.repository.FindFirstWithCustomPreloads(locStruct, preloads, "id = ?", locStructID)
	totalWorkers := int64(0)
	for _, worker := range workers {
		totalWorkers += worker
	}
	fmt.Println("total workers", totalWorkers)
	if totalWorkers > int64(locStruct.Location.Workers) {
		return errors.New("not enough workers")
	}
	fmt.Println("workers", locStruct.ID)
	for _, res := range locStruct.Location.Resources {
		res.AllocatedWorkers = uint64(workers[res.Resource.Name])
		fmt.Println("allocated workers", res.AllocatedWorkers, res.Resource.Name, workers[res.Resource.Name])
		s.repository.Update(&res)
	}
	return nil
}
