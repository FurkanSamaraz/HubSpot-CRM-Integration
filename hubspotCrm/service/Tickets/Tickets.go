package crm_services

import (
	"fmt"
	dto "main/dto"
	crm_repository "main/repository/Tickets"
	crm_structures "main/structure/Tickets"
)

type Service_Tickets interface {
	TicketsCreate(model crm_structures.Create_Tickets) (*dto.Crm_DTO, error)
	TicketsGet(model crm_structures.Get_Tickets) (*dto.Crm_DTO, error)
	TicketsUpdate(model crm_structures.Update_Tickets) (*dto.Crm_DTO, error)
}
type TicketsCRMService struct {
	Repo crm_repository.KnowledgerArticlesRepositoryInterface
}

func (t TicketsCRMService) TicketsGet(model crm_structures.Get_Tickets) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - Tickets |")
	result, err := t.Repo.GetKnowledgerArticles(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
func (t TicketsCRMService) TicketsCreate(model crm_structures.Create_Tickets) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - Tickets |")
	result, err := t.Repo.CreateKnowledgerArticles(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
func (t TicketsCRMService) TicketsUpdate(model crm_structures.Update_Tickets) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - Tickets |")
	result, err := t.Repo.UpdateKnowledgerArticles(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
