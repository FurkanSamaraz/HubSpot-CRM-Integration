package crm_services

import (
	"fmt"
	dto "main/dto"
	crm_repository "main/repository/Companies"
	crm_structures "main/structure/Companies"
)

type CompaniesCRMService struct {
	Repo crm_repository.CompaniesRepositoryInterface
}

type Service_Companies interface {
	CompaniesCreate(model crm_structures.CreateCompanies) (*dto.Crm_DTO, error)
	CompaniesGet(model crm_structures.GetCompanies) (*dto.Crm_DTO, error)
	CompaniesUpdate(model crm_structures.UpdateCompanies) (*dto.Crm_DTO, error)
}

func (t CompaniesCRMService) CompaniesGet(model crm_structures.GetCompanies) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - Companies |")
	result, err := t.Repo.GetCompanies(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
func (t CompaniesCRMService) CompaniesCreate(model crm_structures.CreateCompanies) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - Companies |")
	result, err := t.Repo.CreateCompanies(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
func (t CompaniesCRMService) CompaniesUpdate(model crm_structures.UpdateCompanies) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - Companies |")
	result, err := t.Repo.UpdateCompanies(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
