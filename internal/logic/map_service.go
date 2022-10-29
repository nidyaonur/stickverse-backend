package logic

import (
	"fmt"

	"github.com/nidyaonur/stickverse-backend/entities"
	"github.com/nidyaonur/stickverse-backend/pkg/repository"
	"github.com/nidyaonur/stickverse-backend/pkg/utils"
)

type MapService struct {
	repository repository.Repository
}

func NewMapService(repository repository.Repository) *MapService {
	return &MapService{repository}
}

func (s *MapService) PickUserLocation() (*entities.Location, error) {
	capacities := s.repository.GetCapacities(10)
	preferedCapacity := utils.WeightedRandomCapacity(capacities, 10)
	fmt.Println(preferedCapacity, "is the prefered capacity")
	ids := s.repository.GetMinGridWithCapacity(int(preferedCapacity), 10)
	fmt.Println("ids", ids)
	randId := utils.RandomElement(ids)
	fmt.Println("randId", randId)
	locations := []*entities.Location{}
	_, err := s.repository.FindAllWithConditionValues(
		&locations, nil, nil,
		"grid_id = ? AND user_id IS NULL", "", randId)
	if err != nil {
		return nil, err
	}
	randLocation := utils.RandomElement(locations)
	return randLocation, nil

}

func (s *MapService) GetGridsByCoord(xMin, yMin, xMax, yMax int) ([]*entities.Grid, error) {
	grids := []*entities.Grid{}
	_, err := s.repository.FindAllWithConditionValues(&grids, nil, nil, "x > ? AND y > ? AND x < ? AND y < ?", "", xMin, yMin, xMax, yMax)
	return grids, err

}
