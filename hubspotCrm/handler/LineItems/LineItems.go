package crm_handlers

import (
	crm_services "main/service/LineItems"

	"github.com/gofiber/fiber/v2"
)

type LineItemsHandler struct {
	Service crm_services.Service_LineItems
}

func (h LineItemsHandler) LineItems(c *fiber.Ctx) error {
	return nil
}

func (h LineItemsHandler) LineItemsCreate(c *fiber.Ctx) error {
	return nil
}
func (h LineItemsHandler) LineItemsById(c *fiber.Ctx) error {
	return nil
}
func (h LineItemsHandler) LineItemsUpdate(c *fiber.Ctx) error {
	return nil
}
func (h LineItemsHandler) LineItemsDelete(c *fiber.Ctx) error {
	return nil
}
