package crm_services

import (
	"fmt"
	dto "main/dto"
	crm_repository "main/repository/Deals"
	crm_structures "main/structure/Deals"
)

type Service_Deals interface {
	DealsGet(model crm_structures.GetDeals) (*dto.Crm_DTO, error)
}
type DealsCRMService struct {
	Repo crm_repository.DealsRepositoryInterface
}

func (t DealsCRMService) DealsGet(model crm_structures.GetDeals) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - Deals |")
	result, err := t.Repo.GetDeals(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
