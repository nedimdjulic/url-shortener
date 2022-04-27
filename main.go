package main

import (
	"os"

	"github.com/nedimdjulic/url-shortener/database"
	"github.com/nedimdjulic/url-shortener/handlers"
	repo "github.com/nedimdjulic/url-shortener/repository"
	"github.com/nedimdjulic/url-shortener/service"

	_ "github.com/nedimdjulic/url-shortener/docs"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {

	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	dbConfig := database.Config{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		User:     os.Getenv("POSTGRES_USER"),
		Name:     os.Getenv("POSTGRES_NAME"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
	}

	db, err := database.Open("postgres", &dbConfig)
	if err != nil {

		panic(err)
	}

	if err = db.Init(); err != nil {
		panic(err)
	}

	urlRepo := repo.NewUrlDB(db.DB)
	urlSvc := service.New(urlRepo)
	handler := handlers.NewHTTPSvc(":8080", urlSvc)

	e := echo.New()

	e.POST("/create", func(c echo.Context) error {
		return handler.HandleCreate(c)
	})

	e.GET("/:short", func(c echo.Context) error {
		return handler.HandleRedirect(c)
	})

	e.GET("/count/:short", func(c echo.Context) error {
		return handler.HandleGetCount(c)
	})

	e.DELETE("/delete/:short", func(c echo.Context) error {
		return handler.HandleDelete(c)
	})

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
