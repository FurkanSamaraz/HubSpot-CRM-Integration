package crm_repository

import (
	"fmt"
	crm_structures "main/structure/Quotes"

	"gorm.io/gorm"
)

type QuotesDB struct {
	CrmCollection *gorm.DB
}
type QuotesRepositoryInterface interface {
	GetQuotes(model crm_structures.GetQuotes) (bool, error)
	CreateQuotes(model crm_structures.CreateQuotes) (bool, error)
	UpdateQuotes(model crm_structures.UpdateQuotes) (bool, error)
}

func (t QuotesDB) GetQuotes(model crm_structures.GetQuotes) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - Quotes |")
	return true, nil
}
func (t QuotesDB) CreateQuotes(model crm_structures.CreateQuotes) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - Quotes |")
	return true, nil
}
func (t QuotesDB) UpdateQuotes(model crm_structures.UpdateQuotes) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - Quotes |")
	return true, nil
}
