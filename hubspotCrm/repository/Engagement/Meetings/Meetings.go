package crm_repository

import (
	"fmt"
	crm_structures "main/structure/Engagement/Meetings"

	"gorm.io/gorm"
)

type MeetingsDB struct {
	CrmCollection *gorm.DB
}
type MeetingsRepositoryInterface interface {
	GetMeetings(model crm_structures.GetMeetings) (bool, error)
}

func (t MeetingsDB) GetMeetings(model crm_structures.GetMeetings) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - Meetings |")
	return true, nil
}
