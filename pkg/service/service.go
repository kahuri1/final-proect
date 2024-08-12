package service

import (
	"github.com/kahuri1/final-proect/pkg/model"
	log "github.com/sirupsen/logrus"
)

type repo interface {
	CreateDbTask(task model.Task) (int64, error)
}

type Service struct {
	repo repo
}

func NewService(repo repo) *Service {
	log.Info("service init")

	return &Service{
		repo: repo,
	}
}
