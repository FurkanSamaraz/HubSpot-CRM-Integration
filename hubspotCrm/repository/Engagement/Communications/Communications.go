package crm_repository

import (
	"fmt"
	crm_structures "main/structure/Engagement/Communications"

	"gorm.io/gorm"
)

type CommunicationsDB struct {
	CrmCollection *gorm.DB
}
type CommunicationsRepositoryInterface interface {
	GetCommunications(model crm_structures.GetCommunications) (bool, error)
	CreateCommunications(model crm_structures.CreateCommunications) (bool, error)
	UpdateCommunications(model crm_structures.UpdateCommunications) (bool, error)

}

func (t CommunicationsDB) GetCommunications(model crm_structures.GetCommunications) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - Communications |")
	return true, nil
}
func (t CommunicationsDB) CreateCommunications(model crm_structures.CreateCommunications) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - Communications |")
	return true, nil
}
func (t CommunicationsDB) UpdateCommunications(model crm_structures.UpdateCommunications) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - Communications |")
	return true, nil
}
