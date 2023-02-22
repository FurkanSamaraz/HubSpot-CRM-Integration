package crm_repository

import (
	"fmt"
	crm_structures "main/structure/Engagement/Notes"

	"gorm.io/gorm"
)

type NotesDB struct {
	CrmCollection *gorm.DB
}
type NotesRepositoryInterface interface {
	GetNotes(model crm_structures.GetNotes) (bool, error)
	CreateNotes(model crm_structures.CreateNotes) (bool, error)
	UpdateNotes(model crm_structures.UpdateNotes) (bool, error)
}

func (t NotesDB) GetNotes(model crm_structures.GetNotes) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - Notes |")
	return true, nil
}
func (t NotesDB) CreateNotes(model crm_structures.CreateNotes) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - Notes |")
	return true, nil
}
func (t NotesDB) UpdateNotes(model crm_structures.UpdateNotes) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - Notes |")
	return true, nil
}
