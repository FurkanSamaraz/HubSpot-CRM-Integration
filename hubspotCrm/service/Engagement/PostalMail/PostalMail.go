package crm_repository

import (
	"fmt"
	dto "main/dto"
	crm_repository "main/repository/Engagement/PostalMail"
	crm_structures "main/structure/Engagement/PostalMail"
)

type Service_PostalMail interface {
	PostalMailCreate(model crm_structures.CreatePostalMail) (*dto.Crm_DTO, error)
	PostalMailGet(model crm_structures.GetPostalMail) (*dto.Crm_DTO, error)
	PostalMailUpdate(model crm_structures.UpdatePostalMail) (*dto.Crm_DTO, error)
}
type PostalMailCRMService struct {
	Repo crm_repository.PostalMailRepositoryInterface
}

func (t PostalMailCRMService) PostalMailGet(model crm_structures.GetPostalMail) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - PostalMail |")
	result, err := t.Repo.GetPostalMail(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
func (t PostalMailCRMService) PostalMailCreate(model crm_structures.CreatePostalMail) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - PostalMail |")
	result, err := t.Repo.CreatePostalMail(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
func (t PostalMailCRMService) PostalMailUpdate(model crm_structures.UpdatePostalMail) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - PostalMail |")
	result, err := t.Repo.UpdatePostalMail(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
