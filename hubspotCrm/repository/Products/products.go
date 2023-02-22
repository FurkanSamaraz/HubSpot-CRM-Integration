package crm_repository

import (
	"fmt"
	crm_structures "main/structure/Products"

	"gorm.io/gorm"
)

type ProductsDB struct {
	CrmCollection *gorm.DB
}
type ProductsRepositoryInterface interface {
	GetProducts(model crm_structures.GetProducts) (bool, error)
	CreateProducts(model crm_structures.CreateProducts) (bool, error)
	UpdateProducts(model crm_structures.UpdateProducts) (bool, error)
}

func (t ProductsDB) GetProducts(model crm_structures.GetProducts) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - Products |")
	return true, nil
}
func (t ProductsDB) CreateProducts(model crm_structures.CreateProducts) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - Products |")
	return true, nil
}
func (t ProductsDB) UpdateProducts(model crm_structures.UpdateProducts) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - Products |")
	return true, nil
}
