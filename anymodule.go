package main

import (
	moduleboilerplate "github.com/Twibbonize/go-module-boilerplate-mongodb"
	"github.com/gofiber/fiber/v2"
)

func Create(c *fiber.Ctx, moduleboilerplateCRUD moduleboilerplate.SetterLib) error {
	responseData := "test"
	return c.Status(fiber.StatusOK).JSON(responseData)

}

func Update(c *fiber.Ctx, moduleboilerplateCRUD moduleboilerplate.SetterLib) error {
	responseData := "test"
	return c.Status(fiber.StatusOK).JSON(responseData)

}

func Delete(c *fiber.Ctx, moduleboilerplateCRUD moduleboilerplate.SetterLib) error {
	responseData := "test"
	return c.Status(fiber.StatusOK).JSON(responseData)

}
