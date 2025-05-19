package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"student-service/dbx"
	"student-service/handler"
	"student-service/middleware"
)

func main() {
	_ = godotenv.Load()
	if err := dbx.InitDB(); err != nil {
		log.Fatal("DB Error: ", err)
	}

	app := fiber.New()

	app.Post("/signup", handler.Signup)
	app.Post("/login", handler.Login)

	api := app.Group("/student", middleware.JWTMiddleware)
	api.Get("/list", handler.ListStudents)
	api.Put("/update/:id", handler.UpdateStudent)
	api.Delete("/delete/:id", handler.DeleteStudent)

	log.Fatal(app.Listen(":8080"))
}
