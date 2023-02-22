package crm_services

import (
	"fmt"
	dto "main/dto"
	crm_repository "main/repository/Products"
	crm_structures "main/structure/Products"
)

type Service_Products interface {
	ProductsCreate(model crm_structures.CreateProducts) (*dto.Crm_DTO, error)
	ProductsGet(model crm_structures.GetProducts) (*dto.Crm_DTO, error)
	ProductsUpdate(model crm_structures.UpdateProducts) (*dto.Crm_DTO, error)
}
type ProductsCRMService struct {
	Repo crm_repository.ProductsRepositoryInterface
}

func (t ProductsCRMService) ProductsGet(model crm_structures.GetProducts) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - Products |")
	result, err := t.Repo.GetProducts(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
func (t ProductsCRMService) ProductsCreate(model crm_structures.CreateProducts) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - Products |")
	result, err := t.Repo.CreateProducts(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
func (t ProductsCRMService) ProductsUpdate(model crm_structures.UpdateProducts) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - Products |")
	result, err := t.Repo.UpdateProducts(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
