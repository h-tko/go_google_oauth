package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func envLoad() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	if os.Getenv("APP_ENV") == "" {
		os.Setenv("APP_ENV", "dev")
	}

	return nil
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CSRF())
	e.Use(middleware.AddTrailingSlash())

	if err := envLoad(); err != nil {
		panic(err)
	}

	route(e)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8888"
	}

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
