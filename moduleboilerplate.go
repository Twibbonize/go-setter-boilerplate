package main

import (
	moduleboilerplate "github.com/Twibbonize/go-module-boilerplate-mongodb"
	"github.com/gofiber/fiber/v2"
)

func CreateModuleBoilerplate(c *fiber.Ctx, moduleboilerplateCRUD moduleboilerplate.SetterLib) error {
	responseData := "test"
	return c.Status(fiber.StatusOK).JSON(responseData)

}

func UpdateModuleBoilerplate(c *fiber.Ctx, moduleboilerplateCRUD moduleboilerplate.SetterLib) error {
	responseData := "test"
	return c.Status(fiber.StatusOK).JSON(responseData)
	
}

func DeleteModuleBoilerplate(c *fiber.Ctx, moduleboilerplateCRUD moduleboilerplate.SetterLib) error {
	responseData := "test"
	return c.Status(fiber.StatusOK).JSON(responseData)
	
}
