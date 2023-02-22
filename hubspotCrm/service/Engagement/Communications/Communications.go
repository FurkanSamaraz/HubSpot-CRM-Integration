package crm_repository

import (
	"fmt"
	dto "main/dto"
	crm_repository "main/repository/Engagement/Communications"
	crm_structures "main/structure/Engagement/Communications"
)

type Service_Communications interface {
	CommunicationsCreate(model crm_structures.CreateCommunications) (*dto.Crm_DTO, error)
	CommunicationsGet(model crm_structures.GetCommunications) (*dto.Crm_DTO, error)
	CommunicationsUpdate(model crm_structures.UpdateCommunications) (*dto.Crm_DTO, error)
}
type CommunicationsCRMService struct {
	Repo crm_repository.CommunicationsRepositoryInterface
}

func (t CommunicationsCRMService) CommunicationsGet(model crm_structures.GetCommunications) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - Communications |")
	result, err := t.Repo.GetCommunications(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
func (t CommunicationsCRMService) CommunicationsCreate(model crm_structures.CreateCommunications) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - Communications |")
	result, err := t.Repo.CreateCommunications(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
func (t CommunicationsCRMService) CommunicationsUpdate(model crm_structures.UpdateCommunications) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - Communications |")
	result, err := t.Repo.UpdateCommunications(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
