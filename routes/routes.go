package routes

import (
	"keep-api/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App) {
	app.Get("/notes", controllers.GetNotes)
	app.Post("/note", controllers.CreateNote)
	app.Put("/note", controllers.UpdateNote)
	app.Delete("/note/:id", controllers.DeleteNote)
}
