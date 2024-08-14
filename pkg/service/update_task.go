package service

import (
	"fmt"
	"github.com/kahuri1/final-proect/pkg/model"
	log "github.com/sirupsen/logrus"
)

func (s *Service) UpdateTask(task *model.Task) (bool, error) {
	_, err := s.repo.UpdateTask(task)
	if err != nil {
		fmt.Println(task)
		log.Errorf("failed to update task: %w", err)
		return false, err
	}
	return true, nil
}
