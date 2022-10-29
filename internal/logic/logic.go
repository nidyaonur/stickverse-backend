package logic

import "github.com/nidyaonur/stickverse-backend/pkg/repository"

type Service struct {
	Map       *MapService
	Structure *StructureService
	User      *UserService
}

func NewService(repository repository.Repository) *Service {
	return &Service{
		Map:       NewMapService(repository),
		Structure: NewStructureService(repository),
		User:      NewUserService(repository),
	}
}
