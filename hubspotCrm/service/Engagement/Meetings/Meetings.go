package crm_services

import (
	"fmt"
	dto "main/dto"
	crm_repository "main/repository/Engagement/Meetings"
	crm_structures "main/structure/Engagement/Meetings"
)

type Service_Meetings interface {
	MeetingsGet(model crm_structures.GetMeetings) (*dto.Crm_DTO, error)
}
type MeetingsCRMService struct {
	Repo crm_repository.MeetingsRepositoryInterface
}

func (t MeetingsCRMService) MeetingsGet(model crm_structures.GetMeetings) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - Meetings |")
	result, err := t.Repo.GetMeetings(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
