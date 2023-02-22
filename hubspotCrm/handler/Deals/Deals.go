package crm_handlers

import (
	crm_integration "main/package"
	crm_services "main/service/Deals"

	"github.com/gofiber/fiber/v2"
)

type DealsHandler struct {
	Service crm_services.Service_Deals
}
type CRM_CONTROLLER struct {
	Base *crm_integration.CRM_BASE
}

func (h DealsHandler) Deals(c *fiber.Ctx) error {
	return nil
}

func (h DealsHandler) DealsCreate(c *fiber.Ctx) error {
	return nil
}
func (h DealsHandler) DealsById(c *fiber.Ctx) error {
	return nil
}
func (h DealsHandler) DealsUpdate(c *fiber.Ctx) error {
	return nil
}
func (h DealsHandler) DealsDelete(c *fiber.Ctx) error {
	return nil
}
