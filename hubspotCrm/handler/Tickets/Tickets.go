package crm_handlers

import (
	crm_services "main/service/Tickets"

	"github.com/gofiber/fiber/v2"
)

type TicketsHandler struct {
	Service crm_services.Service_Tickets
}

func (h TicketsHandler) Tickets(c *fiber.Ctx) error {
	return nil
}

func (h TicketsHandler) TicketsCreate(c *fiber.Ctx) error {
	return nil
}
func (h TicketsHandler) TicketsById(c *fiber.Ctx) error {
	return nil
}
func (h TicketsHandler) TicketsUpdate(c *fiber.Ctx) error {
	return nil
}
func (h TicketsHandler) TicketsDelete(c *fiber.Ctx) error {
	return nil
}
