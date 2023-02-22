package crm_repository

import (
	"fmt"
	crm_structures "main/structure/Engagement/PostalMail"

	"gorm.io/gorm"
)

type PostalMailDB struct {
	CrmCollection *gorm.DB
}
type PostalMailRepositoryInterface interface {
	GetPostalMail(model crm_structures.GetPostalMail) (bool, error)
	CreatePostalMail(model crm_structures.CreatePostalMail) (bool, error)
	UpdatePostalMail(model crm_structures.UpdatePostalMail) (bool, error)
}

func (t PostalMailDB) GetPostalMail(model crm_structures.GetPostalMail) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - PostalMail |")
	return true, nil
}
func (t PostalMailDB) CreatePostalMail(model crm_structures.CreatePostalMail) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - PostalMail |")
	return true, nil
}
func (t PostalMailDB) UpdatePostalMail(model crm_structures.UpdatePostalMail) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - PostalMail |")
	return true, nil
}
