package crm_services

import (
	"fmt"
	dto "main/dto"
	crm_repository "main/repository/Engagement/Emails"
	crm_structures "main/structure/Engagement/Emails"
)

type Service_Emails interface {
	EmailsCreate(model crm_structures.CreateEmails) (*dto.Crm_DTO, error)
	EmailsGet(model crm_structures.GetEmails) (*dto.Crm_DTO, error)
	EmailsUpdate(model crm_structures.UpdateEmails) (*dto.Crm_DTO, error)
}
type EmailsCRMService struct {
	Repo crm_repository.EmailsRepositoryInterface
}

func (t EmailsCRMService) EmailsGet(model crm_structures.GetEmails) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - Emails |")
	result, err := t.Repo.GetEmails(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
func (t EmailsCRMService) EmailsCreate(model crm_structures.CreateEmails) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - Emails |")
	result, err := t.Repo.CreateEmails(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
func (t EmailsCRMService) EmailsUpdate(model crm_structures.UpdateEmails) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - Emails |")
	result, err := t.Repo.UpdateEmails(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
