package crm_handlers

import (
	crm_services "main/service/Engagement/Calls"

	"github.com/gofiber/fiber/v2"
)

type CallsHandler struct {
	Service crm_services.Service_Calls
}

func (h CallsHandler) Calls(c *fiber.Ctx) error {
	return nil
}
