package crm_repository

import (
	"fmt"
	dto "main/dto"
	crm_repository "main/repository/Engagement/Notes"
	crm_structures "main/structure/Engagement/Notes"
)

type Service_Notes interface {
	NotesCreate(model crm_structures.CreateNotes) (*dto.Crm_DTO, error)
	NotesGet(model crm_structures.GetNotes) (*dto.Crm_DTO, error)
	NotesUpdate(model crm_structures.UpdateNotes) (*dto.Crm_DTO, error)
}
type NotesCRMService struct {
	Repo crm_repository.NotesRepositoryInterface
}

func (t NotesCRMService) NotesGet(model crm_structures.GetNotes) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - Notes |")
	result, err := t.Repo.GetNotes(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
func (t NotesCRMService) NotesCreate(model crm_structures.CreateNotes) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - Notes |")
	result, err := t.Repo.CreateNotes(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
func (t NotesCRMService) NotesUpdate(model crm_structures.UpdateNotes) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - Notes |")
	result, err := t.Repo.UpdateNotes(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
