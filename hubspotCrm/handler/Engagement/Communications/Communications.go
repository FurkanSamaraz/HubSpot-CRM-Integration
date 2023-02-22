package crm_repository


import (
	crm_services "main/service/Engagement/Communications"

	"github.com/gofiber/fiber/v2"
)

type CommunicationsHandler struct {
	Service crm_services.Service_Communications
}

func (h CommunicationsHandler) Communications(c *fiber.Ctx) error {
	return nil
}

func (h CommunicationsHandler) CommunicationsCreate(c *fiber.Ctx) error {
	return nil
}
func (h CommunicationsHandler) CommunicationsById(c *fiber.Ctx) error {
	return nil
}
func (h CommunicationsHandler) CommunicationsUpdate(c *fiber.Ctx) error {
	return nil
}
func (h CommunicationsHandler) CommunicationsDelete(c *fiber.Ctx) error {
	return nil
}
