package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/seccret404/PaymenMini-Payment-Gateway-Fraud-Detection-GoFiber/internal/handler"
	"github.com/seccret404/PaymenMini-Payment-Gateway-Fraud-Detection-GoFiber/internal/middleware"
)

func RoutesApp(app *fiber.App, authHandler *handler.AuthHandler, producthandler *handler.ProductHandler, cartHandler *handler.CartHandler) {
	api := app.Group("/api")

	api.Post("/register", authHandler.Register)
	api.Post("/login", authHandler.Login)

	api.Post("/create-product",middleware.JWTProtected(), producthandler.CreateProducts)
	api.Get("/get-product", producthandler.GetAllProducts)
	api.Get("/get-product/:id", producthandler.GetByIdProduct)
	api.Put("/update-product/:id", producthandler.UpdateProduct)
	api.Delete("/delete-product/:id", producthandler.DeleteProduct)

	api.Post("/create-cart", cartHandler.CreateCarts)
	//for role setup
	// protected := api.Group("/admin")
	// protected.Use(middleware.JWTProtected(), middleware.RoleRequired("role(buyer/admin)"))
	// protected.Get("/dashboard", func(c *fiber.Ctx) error {
	// 	return c.SendString("Admin Dashboard!")
	// })
}
