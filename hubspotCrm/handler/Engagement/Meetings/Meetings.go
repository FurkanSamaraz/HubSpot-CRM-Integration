package crm_handlers

import (
	crm_services "main/service/Engagement/Meetings"

	"github.com/gofiber/fiber/v2"
)

type MeetingsHandler struct {
	Service crm_services.Service_Meetings
}

func (h MeetingsHandler) Meetings(c *fiber.Ctx) error {
	return nil
}
