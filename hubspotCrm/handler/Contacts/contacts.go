package crm_handlers

import (
	crm_services "main/service/Contacts"

	"github.com/gofiber/fiber/v2"
)

type ContactsHandler struct {
	Service crm_services.Service_Contacts
}

func (h ContactsHandler) Contacts(c *fiber.Ctx) error {
	return nil
}

func (h ContactsHandler) ContactsCreate(c *fiber.Ctx) error {
	return nil
}
func (h ContactsHandler) ContactsById(c *fiber.Ctx) error {
	return nil
}
func (h ContactsHandler) ContactsUpdate(c *fiber.Ctx) error {
	return nil
}
func (h ContactsHandler) ContactsDelete(c *fiber.Ctx) error {
	return nil
}
