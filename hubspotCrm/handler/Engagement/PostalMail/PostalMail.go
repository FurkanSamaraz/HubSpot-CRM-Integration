package crm_repository

import (
	crm_services "main/service/Engagement/PostalMail"

	"github.com/gofiber/fiber/v2"
)

type PostalMailHandler struct {
	Service crm_services.Service_PostalMail
}

func (h PostalMailHandler) PostalMail(c *fiber.Ctx) error {
	return nil
}

func (h PostalMailHandler) PostalMailCreate(c *fiber.Ctx) error {
	return nil
}
func (h PostalMailHandler) PostalMailById(c *fiber.Ctx) error {
	return nil
}
func (h PostalMailHandler) PostalMailUpdate(c *fiber.Ctx) error {
	return nil
}
func (h PostalMailHandler) PostalMailDelete(c *fiber.Ctx) error {
	return nil
}
