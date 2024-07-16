package main

import (
	"go-fiber-newrelic/src/configuration"
	ds "go-fiber-newrelic/src/domain/datasources"
	repo "go-fiber-newrelic/src/domain/repositories"
	"go-fiber-newrelic/src/gateways"
	"go-fiber-newrelic/src/middlewares"
	sv "go-fiber-newrelic/src/services"
	"log"
	"os"

	swagger "github.com/gofiber/contrib/swagger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func main() {

	// // // remove this before deploy ###################
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// /// ############################################

	app := fiber.New(configuration.NewFiberConfiguration())
	app.Use(swagger.New(swagger.Config{
		BasePath: "/api/",
		FilePath: "./src/docs/swagger.yaml",
		Path:     "docs",
	}))
	app.Use(middlewares.NewLogger(app))
	app.Use(recover.New())
	app.Use(middlewares.NewRelic(app))
	app.Use(cors.New())

	mongodb := ds.NewMongoDB(10)

	userMongo := repo.NewUsersRepository(mongodb)

	sv0 := sv.NewUsersService(userMongo)
	sv1 := sv.NewIpService()

	gateways.NewHTTPGateway(app, sv0, sv1)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	app.Listen(":" + PORT)
}
