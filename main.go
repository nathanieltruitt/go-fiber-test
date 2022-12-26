package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

type Ninja struct {
	Name string
	Weapon string
}

func getNinja(ctx *fiber.Ctx) error {
	wallace := Ninja{Name: "wallace", Weapon: "Ninja Star"}
	return ctx.Status(fiber.StatusOK).JSON(wallace)
}

func createNinja(ctx *fiber.Ctx) error {
	body := new(Ninja)
	err := ctx.BodyParser(body)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
		return err
	} else {
		return ctx.Status(fiber.StatusOK).JSON(&body)
	}
}

func main() {
	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello World!")
	})
	app.Use(logger.New())
	app.Use(requestid.New())
	ninjaApp := app.Group("/ninja")
	ninjaApp.Get("", getNinja)
	ninjaApp.Post("", createNinja)
	app.Listen(":80")
}