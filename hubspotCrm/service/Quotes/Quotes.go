package crm_services

import (
	"fmt"
	dto "main/dto"
	crm_repository "main/repository/Quotes"
	crm_structures "main/structure/Quotes"
)

type Service_Quotes interface {
	QuotesCreate(model crm_structures.CreateQuotes) (*dto.Crm_DTO, error)
	QuotesGet(model crm_structures.GetQuotes) (*dto.Crm_DTO, error)
	QuotesUpdate(model crm_structures.UpdateQuotes) (*dto.Crm_DTO, error)
}
type QuotesCRMService struct {
	Repo crm_repository.QuotesRepositoryInterface
}

func (t QuotesCRMService) QuotesGet(model crm_structures.GetQuotes) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - Quotes |")
	result, err := t.Repo.GetQuotes(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
func (t QuotesCRMService) QuotesCreate(model crm_structures.CreateQuotes) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - Quotes |")
	result, err := t.Repo.CreateQuotes(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
func (t QuotesCRMService) QuotesUpdate(model crm_structures.UpdateQuotes) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - Quotes |")
	result, err := t.Repo.UpdateQuotes(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
