package crm_handlers

import (
	crm_services "main/service/FeedbackSubmissions"

	"github.com/gofiber/fiber/v2"
)

type FeedbackSubmissionsHandler struct {
	Service crm_services.Service_FeedbackSubmissions
}

func (h FeedbackSubmissionsHandler) FeedbackSubmissions(c *fiber.Ctx) error {
	return nil
}

func (h FeedbackSubmissionsHandler) FeedbackSubmissionsCreate(c *fiber.Ctx) error {
	return nil
}
func (h FeedbackSubmissionsHandler) FeedbackSubmissionsById(c *fiber.Ctx) error {
	return nil
}
func (h FeedbackSubmissionsHandler) FeedbackSubmissionsUpdate(c *fiber.Ctx) error {
	return nil
}
func (h FeedbackSubmissionsHandler) FeedbackSubmissionsDelete(c *fiber.Ctx) error {
	return nil
}
