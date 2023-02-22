package crm_handlers

import (
	crm_services "main/service/Engagement/Emails"

	"github.com/gofiber/fiber/v2"
)

type EmailsHandler struct {
	Service crm_services.Service_Emails
}

func (h EmailsHandler) Emails(c *fiber.Ctx) error {
	return nil
}

func (h EmailsHandler) EmailsCreate(c *fiber.Ctx) error {
	return nil
}
func (h EmailsHandler) EmailsById(c *fiber.Ctx) error {
	return nil
}
func (h EmailsHandler) EmailsUpdate(c *fiber.Ctx) error {
	return nil
}
func (h EmailsHandler) EmailsDelete(c *fiber.Ctx) error {
	return nil
}
