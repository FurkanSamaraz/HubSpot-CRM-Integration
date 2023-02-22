package crm_repository

import (
	"fmt"
	crm_structures "main/structure/LineItems"

	"gorm.io/gorm"
)

type LineItemsDB struct {
	CrmCollection *gorm.DB
}
type LineItemsRepositoryInterface interface {
	GetLineItems(model crm_structures.GetLineItems) (bool, error)
	CreateLineItems(model crm_structures.CreateLineItems) (bool, error)
	UpdateLineItems(model crm_structures.UpdateLineItems) (bool, error)
}

func (t LineItemsDB) GetLineItems(model crm_structures.GetLineItems) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - LineItems |")
	return true, nil
}
func (t LineItemsDB) CreateLineItems(model crm_structures.CreateLineItems) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - LineItems |")
	return true, nil
}
func (t LineItemsDB) UpdateLineItems(model crm_structures.UpdateLineItems) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - LineItems |")
	return true, nil
}
