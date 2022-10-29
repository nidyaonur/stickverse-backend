package data

import "github.com/nidyaonur/stickverse-backend/pkg/repository"

type data struct {
	repo repository.Repository
}

func NewDataGenerator(repo repository.Repository) *data {
	return &data{repo: repo}
}
