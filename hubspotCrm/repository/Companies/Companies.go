package crm_repository

import (
	"fmt"
	crm_structures "main/structure/Companies"

	"gorm.io/gorm"
)

type CompaniesDB struct {
	CrmCollection *gorm.DB
}
type CompaniesRepositoryInterface interface {
	GetCompanies(model crm_structures.GetCompanies) (bool, error)
	CreateCompanies(model crm_structures.CreateCompanies) (bool, error)
	UpdateCompanies(model crm_structures.UpdateCompanies) (bool, error)
}

func (t CompaniesDB) GetCompanies(model crm_structures.GetCompanies) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - Companies |")
	return true, nil
}
func (t CompaniesDB) CreateCompanies(model crm_structures.CreateCompanies) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - Companies |")
	return true, nil
}
func (t CompaniesDB) UpdateCompanies(model crm_structures.UpdateCompanies) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - Companies |")
	return true, nil
}
