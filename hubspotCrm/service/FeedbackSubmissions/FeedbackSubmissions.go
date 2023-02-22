package crm_services

import (
	"fmt"
	dto "main/dto"
	crm_repository "main/repository/FeedbackSubmissions"
	crm_structures "main/structure/FeedbackSubmissions"
)

type Service_FeedbackSubmissions interface {
	FeedbackSubmissionsGet(model crm_structures.GetFeedbackSubmissions) (*dto.Crm_DTO, error)
}
type FeedbackSubmissionsCRMService struct {
	Repo crm_repository.FeedbackSubmissionsRepositoryInterface
}

func (t FeedbackSubmissionsCRMService) FeedbackSubmissionsGet(model crm_structures.GetFeedbackSubmissions) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - FeedbackSubmissions |")
	result, err := t.Repo.GetFeedbackSubmissions(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
