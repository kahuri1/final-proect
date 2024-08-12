package service

import (
	"github.com/kahuri1/final-proect/pkg/model"
	log "github.com/sirupsen/logrus"
)

func (s *Service) GetTasks(search string) (model.TasksResp, error) {
	tasks, err := s.repo.GetTasks(search)
	if err != nil {
		log.Errorf("failed Get tasks: %w", err)
	}

	return tasks, nil
}
