package service

import (
	"github.com/kahuri1/final-proect/pkg/model"
	log "github.com/sirupsen/logrus"
)

func (s *Service) GetTaskById(id int) (*model.Task, error) {

	task, err := s.repo.GetTaskById(id)
	if err != nil {
		log.Errorf("failed to create message: %w", err)
	}
	return task, nil
}
