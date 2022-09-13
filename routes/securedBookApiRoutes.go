package routes

import (
	"go-fiber-restful-microservice/controllers"
	"go-fiber-restful-microservice/middleware"

	"github.com/gofiber/fiber/v2"
)

// PrivateRoutes func for describe group of private routes.
func PrivateRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/secured")

	// Routes for POST method:
	route.Post("/book", middleware.JWTProtected(), controllers.CreateBook) // create a new book

	// Routes for PUT method:
	route.Put("/book", middleware.JWTProtected(), controllers.UpdateBook) // update one book by ID

	// Routes for DELETE method:
	route.Delete("/book", middleware.JWTProtected(), controllers.DeleteBook) // delete one book by ID
}
