package crm_repository

import (
	"fmt"
	crm_structures "main/structure/Engagement/Tasks"

	"gorm.io/gorm"
)

type TasksDB struct {
	CrmCollection *gorm.DB
}
type TasksRepositoryInterface interface {
	GetTasks(model crm_structures.GetTasks) (bool, error)
	CreateTasks(model crm_structures.CreateTasks) (bool, error)
	UpdateTasks(model crm_structures.UpdateTasks) (bool, error)
}

func (t TasksDB) GetTasks(model crm_structures.GetTasks) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - Tasks |")
	return true, nil
}
func (t TasksDB) CreateTasks(model crm_structures.CreateTasks) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - Tasks |")
	return true, nil
}
func (t TasksDB) UpdateTasks(model crm_structures.UpdateTasks) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - Tasks |")
	return true, nil
}
