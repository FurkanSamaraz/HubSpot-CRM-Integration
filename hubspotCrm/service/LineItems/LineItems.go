package crm_services

import (
	"fmt"
	dto "main/dto"
	crm_repository "main/repository/LineItems"
	crm_structures "main/structure/LineItems"
)

type Service_LineItems interface {
	LineItemsCreate(model crm_structures.CreateLineItems) (*dto.Crm_DTO, error)
	LineItemsGet(model crm_structures.GetLineItems) (*dto.Crm_DTO, error)
	LineItemsUpdate(model crm_structures.UpdateLineItems) (*dto.Crm_DTO, error)
}
type LineItemsCRMService struct {
	Repo crm_repository.LineItemsRepositoryInterface
}

func (t LineItemsCRMService) LineItemsGet(model crm_structures.GetLineItems) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - LineItems |")
	result, err := t.Repo.GetLineItems(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
func (t LineItemsCRMService) LineItemsCreate(model crm_structures.CreateLineItems) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - LineItems |")
	result, err := t.Repo.CreateLineItems(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
func (t LineItemsCRMService) LineItemsUpdate(model crm_structures.UpdateLineItems) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - LineItems |")
	result, err := t.Repo.UpdateLineItems(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
