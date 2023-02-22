package crm_handlers

import (
	crm_services "main/service/Companies"

	"github.com/gofiber/fiber/v2"
)

type CompaniesHandler struct {
	Service crm_services.Service_Companies
}

func (h CompaniesHandler) Companies(c *fiber.Ctx) error {
	return nil
}

func (h CompaniesHandler) CompaniesCreate(c *fiber.Ctx) error {
	return nil
}
func (h CompaniesHandler) CompaniesById(c *fiber.Ctx) error {
	return nil
}
func (h CompaniesHandler) CompaniesUpdate(c *fiber.Ctx) error {
	return nil
}
func (h CompaniesHandler) CompaniesDelete(c *fiber.Ctx) error {
	return nil
}
