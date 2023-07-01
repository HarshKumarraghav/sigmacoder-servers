package routes

import (
	"sigmacoder/pkg/allquestions"

	"github.com/gofiber/fiber/v2"
)

// The `allquestionsHandler` function is a handler function that retrieves all questions from a
// repository and returns them as a JSON response.
func allquestionsHandler(repo *allquestions.Repo) fiber.Handler {
	return func(c *fiber.Ctx) error {
		allquestions, err := repo.ReadAllQuestion()
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(200).JSON(allquestions)
	}
}

// The function `questionByIdHandler` retrieves a question by its ID from a repository and returns it
// as a JSON response.
func questionByIdHandler(repo *allquestions.Repo) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		question, err := repo.ReadByID(id)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(200).JSON(question)
	}
}

// The function creates routes for handling requests related to all questions.
func CreateAllQuestionRoutes(app *fiber.App, allquestionRepo *allquestions.Repo) {
	app.Get("/api/all/allquestions", allquestionsHandler(allquestionRepo))
	app.Get("/api/all/question/:id", questionByIdHandler(allquestionRepo))
}
