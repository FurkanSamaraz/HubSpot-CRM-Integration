package crm_repository

import (
	"fmt"
	crm_structures "main/structure/Engagement/Emails"

	"gorm.io/gorm"
)

type EmailsDB struct {
	CrmCollection *gorm.DB
}
type EmailsRepositoryInterface interface {
	GetEmails(model crm_structures.GetEmails) (bool, error)
	CreateEmails(model crm_structures.CreateEmails) (bool, error)
	UpdateEmails(model crm_structures.UpdateEmails) (bool, error)
}

func (t EmailsDB) GetEmails(model crm_structures.GetEmails) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - Emails |")
	return true, nil
}
func (t EmailsDB) CreateEmails(model crm_structures.CreateEmails) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - Emails |")
	return true, nil
}
func (t EmailsDB) UpdateEmails(model crm_structures.UpdateEmails) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - Emails |")
	return true, nil
}
