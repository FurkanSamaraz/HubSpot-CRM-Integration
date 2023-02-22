package crm_services

import (
	"fmt"
	dto "main/dto"
	crm_repository "main/repository/Contacts"
	crm_structures "main/structure/Contacts"
)

type Service_Contacts interface {
	ContactsCreate(model crm_structures.CreateContacts) (*dto.Crm_DTO, error)
	ContactsGet(model crm_structures.GetContacts) (*dto.Crm_DTO, error)
	ContactsUpdate(model crm_structures.UpdateContacts) (*dto.Crm_DTO, error)
}
type ContactsCRMService struct {
	Repo crm_repository.ContactsRepositoryInterface
}

func (t ContactsCRMService) ContactsGet(model crm_structures.GetContacts) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - Contacts |")
	result, err := t.Repo.GetContacts(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
func (t ContactsCRMService) ContactsCreate(model crm_structures.CreateContacts) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - Contacts |")
	result, err := t.Repo.CreateContacts(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
func (t ContactsCRMService) ContactsUpdate(model crm_structures.UpdateContacts) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - Contacts |")
	result, err := t.Repo.UpdateContacts(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
