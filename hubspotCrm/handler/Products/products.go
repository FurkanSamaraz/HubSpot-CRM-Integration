package crm_handlers

import (
	crm_services "main/service/Products"

	"github.com/gofiber/fiber/v2"
)

type ProductsHandler struct {
	Service crm_services.Service_Products
}

func (h ProductsHandler) Products(c *fiber.Ctx) error {
	return nil
}

func (h ProductsHandler) ProductsCreate(c *fiber.Ctx) error {
	return nil
}
func (h ProductsHandler) ProductsById(c *fiber.Ctx) error {
	return nil
}
func (h ProductsHandler) ProductsUpdate(c *fiber.Ctx) error {
	return nil
}
func (h ProductsHandler) ProductsDelete(c *fiber.Ctx) error {
	return nil
}
