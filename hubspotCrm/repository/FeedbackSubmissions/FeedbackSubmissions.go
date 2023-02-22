package crm_repository

import (
	"fmt"
	crm_structures "main/structure/FeedbackSubmissions"

	"gorm.io/gorm"
)

type FeedbackSubmissionsDB struct {
	CrmCollection *gorm.DB
}
type FeedbackSubmissionsRepositoryInterface interface {
	GetFeedbackSubmissions(model crm_structures.GetFeedbackSubmissions) (bool, error)
}

func (t FeedbackSubmissionsDB) GetFeedbackSubmissions(model crm_structures.GetFeedbackSubmissions) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - FeedbackSubmissions |")
	return true, nil
}
