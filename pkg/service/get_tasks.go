package service

import (
	"fmt"
	"github.com/kahuri1/final-proect/pkg/model"
)

func (s *Service) GetTasks(search string) (model.TasksResp, error) {
	tasks, err := s.repo.GetTasks(search)
	if err != nil {
		fmt.Println(search)
		return model.TasksResp{}, err
	}

	return tasks, nil
}
