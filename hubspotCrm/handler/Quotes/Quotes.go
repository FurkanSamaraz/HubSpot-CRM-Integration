package crm_handlers

import (
	crm_services "main/service/Quotes"

	"github.com/gofiber/fiber/v2"
)

type QuotesHandler struct {
	Service crm_services.Service_Quotes
}

func (h QuotesHandler) Quotes(c *fiber.Ctx) error {
	return nil
}

func (h QuotesHandler) QuotesCreate(c *fiber.Ctx) error {
	return nil
}
func (h QuotesHandler) QuotesById(c *fiber.Ctx) error {
	return nil
}
func (h QuotesHandler) QuotesUpdate(c *fiber.Ctx) error {
	return nil
}
func (h QuotesHandler) QuotesDelete(c *fiber.Ctx) error {
	return nil
}
