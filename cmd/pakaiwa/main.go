/*
 * Copyright (c) 2025 KAnggara75
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * See <https://www.gnu.org/licenses/gpl-3.0.html>.
 *
 * @author KAnggara75 on Fri 08/08/25 08.29
 * @project PakaiWA pakaiwa
 * https://github.com/PakaiWA/PakaiWA/tree/main/cmd/pakaiwa
 */

package main

import (
	"github.com/KAnggara75/scc2go"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
)

func init() {
	scc2go.GetEnv(os.Getenv("SCC_URL"), os.Getenv("AUTH"))
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/api/posts", func(c *fiber.Ctx) error {
		return c.Status(404).JSON(&fiber.Map{
			"success": false,
			"error":   "There are no posts!",
		})
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("GET request")
	})

	app.Get("/:param", func(c *fiber.Ctx) error {
		return c.SendString("param: " + c.Params("param"))
	})

	app.Post("/", func(c *fiber.Ctx) error {
		return c.SendString("POST request")
	})

	log.Fatal(app.Listen(":3000"))
}
