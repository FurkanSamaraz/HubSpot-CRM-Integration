package crm_repository

import (
	"fmt"
	crm_structures "main/structure/Engagement/Calls"

	"gorm.io/gorm"
)

type CallsDB struct {
	CrmCollection *gorm.DB
}
type CallsRepositoryInterface interface {
	GetCalls(model crm_structures.GetCalls) (bool, error)
}

func (t CallsDB) GetCalls(model crm_structures.GetCalls) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - Calls |")
	return true, nil
}
