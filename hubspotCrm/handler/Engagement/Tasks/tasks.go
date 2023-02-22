package crm_handlers

import (
	crm_services "main/service/Engagement/Tasks"

	"github.com/gofiber/fiber/v2"
)

type TasksHandler struct {
	Service crm_services.Service_Tasks
}

func (h TasksHandler) Tasks(c *fiber.Ctx) error {
	return nil
}

func (h TasksHandler) TasksCreate(c *fiber.Ctx) error {
	return nil
}
func (h TasksHandler) TasksById(c *fiber.Ctx) error {
	return nil
}
func (h TasksHandler) TasksUpdate(c *fiber.Ctx) error {
	return nil
}
func (h TasksHandler) TasksDelete(c *fiber.Ctx) error {
	return nil
}
