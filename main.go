package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
)

type Todo struct {
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func main() {
	app := fiber.New()

	todos := []Todo{}

	// Define a route for the GET method on the root path '/'
	app.Get("/api/todos", func(c fiber.Ctx) error {
		// Send a string response to the client
		return c.Status(200).JSON(fiber.Map{"msg": "hello world"})
	})

	app.Post("/api/todos", func(c fiber.Ctx) error {
		todo := Todo{}
		err := c.Bind().JSON(todo)
		if err != nil {
			return err
		}

		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Todo body is required"})
		}

		todo.ID = len(todos) + 1
		todos = append(todos, todo)

		return c.Status(201).JSON(todo)
	})

	app.Patch("/api/todos/:id", func(c fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos[i].Completed = true
				return c.Status(200).JSON(todos[i])
			}
		}

		return c.Status(404).JSON(fiber.Map{"error": "Todod not found"})
	})

	// Start the server on port 3000
	log.Fatal(app.Listen("127.0.0.1:3000"))
}
