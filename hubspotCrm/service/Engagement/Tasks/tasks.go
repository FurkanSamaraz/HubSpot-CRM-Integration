package crm_services

import (
	"fmt"
	dto "main/dto"
	crm_repository "main/repository/Engagement/Tasks"
	crm_structures "main/structure/Engagement/Tasks"
)

type Service_Tasks interface {
	TasksCreate(model crm_structures.CreateTasks) (*dto.Crm_DTO, error)
	TasksGet(model crm_structures.GetTasks) (*dto.Crm_DTO, error)
	TasksUpdate(model crm_structures.UpdateTasks) (*dto.Crm_DTO, error)
}
type TasksCRMService struct {
	Repo crm_repository.TasksRepositoryInterface
}

func (t TasksCRMService) TasksGet(model crm_structures.GetTasks) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - Tasks |")
	result, err := t.Repo.GetTasks(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
func (t TasksCRMService) TasksCreate(model crm_structures.CreateTasks) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - Tasks |")
	result, err := t.Repo.CreateTasks(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
func (t TasksCRMService) TasksUpdate(model crm_structures.UpdateTasks) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - Tasks |")
	result, err := t.Repo.UpdateTasks(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
