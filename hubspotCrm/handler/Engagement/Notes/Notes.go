package crm_repository

import (
	crm_services "main/service/Engagement/Notes"

	"github.com/gofiber/fiber/v2"
)

type NotesHandler struct {
	Service crm_services.Service_Notes
}

func (h NotesHandler) Notes(c *fiber.Ctx) error {
	return nil
}

func (h NotesHandler) NotesCreate(c *fiber.Ctx) error {
	return nil
}
func (h NotesHandler) NotesById(c *fiber.Ctx) error {
	return nil
}
func (h NotesHandler) NotesUpdate(c *fiber.Ctx) error {
	return nil
}
func (h NotesHandler) NotesDelete(c *fiber.Ctx) error {
	return nil
}
