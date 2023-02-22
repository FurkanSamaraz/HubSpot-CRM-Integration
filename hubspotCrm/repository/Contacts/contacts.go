package crm_repository

import (
	"fmt"
	crm_structures "main/structure/Contacts"

	"gorm.io/gorm"
)

type ContactsDB struct {
	CrmCollection *gorm.DB
}
type ContactsRepositoryInterface interface {
	GetContacts(model crm_structures.GetContacts) (bool, error)
	CreateContacts(model crm_structures.CreateContacts) (bool, error)
	UpdateContacts(model crm_structures.UpdateContacts) (bool, error)
}

func (t ContactsDB) GetContacts(model crm_structures.GetContacts) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - Contacts |")
	return true, nil
}
func (t ContactsDB) CreateContacts(model crm_structures.CreateContacts) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - Contacts |")
	return true, nil
}
func (t ContactsDB) UpdateContacts(model crm_structures.UpdateContacts) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - Contacts |")
	return true, nil
}
