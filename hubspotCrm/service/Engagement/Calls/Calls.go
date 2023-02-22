package crm_services

import (
	"fmt"
	dto "main/dto"
	crm_repository "main/repository/Engagement/Calls"
	crm_structures "main/structure/Engagement/Calls"
)

type Service_Calls interface {
	CallsGet(model crm_structures.GetCalls) (*dto.Crm_DTO, error)
}
type CallsCRMService struct {
	Repo crm_repository.CallsRepositoryInterface
}

func (t CallsCRMService) CallsGet(model crm_structures.GetCalls) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - Calls |")
	result, err := t.Repo.GetCalls(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
