package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/seccret404/PaymenMini-Payment-Gateway-Fraud-Detection-GoFiber/internal/handler"
	"github.com/seccret404/PaymenMini-Payment-Gateway-Fraud-Detection-GoFiber/internal/repository"
	"github.com/seccret404/PaymenMini-Payment-Gateway-Fraud-Detection-GoFiber/internal/routes"
	"github.com/seccret404/PaymenMini-Payment-Gateway-Fraud-Detection-GoFiber/internal/usecase"
	"github.com/seccret404/PaymenMini-Payment-Gateway-Fraud-Detection-GoFiber/pkg/database"
)

func main() {
	fmt.Println("hello world")
	log.Println("🔥 Very Nice!")

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Gagal load env")
	}

	db, err := database.ConnectDatabase()
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate()

	//setup fiber

	app := fiber.New()

	//cors set
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET, POST, DELETE, PUT",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	//auth
	authRepo := repository.NewAuthRepository()
	authUsecase := usecase.NewAuthUsecase(authRepo)
	authHandler := handler.NewAuthHandler(authUsecase)

	//product
	productRepo := repository.NewProductRepository()
	productUsecase := usecase.NewProductUsecase(productRepo)
	productHandler := handler.NewProductHandler(productUsecase)

	//cart
	cartRepo := repository.NewCartRepository()
	cartUsecase := usecase.NewCartUsecase(cartRepo)
	cartHandler := handler.NewCartHandler(cartUsecase)
	
	routes.RoutesApp(app, authHandler, productHandler, cartHandler)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Edward Tua Panjaitan S.Tr.Kom")
	})
	 
	log.Fatal(app.Listen(":8080"))

	

}
