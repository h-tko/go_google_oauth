package main

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

const (
	defaultPort = "8888"
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

func setLogLevel(e *echo.Echo) {
	if os.Getenv("APP_ENV") == "dev" {
		e.Logger.SetLevel(log.DEBUG)
	} else {
		e.Logger.SetLevel(log.INFO)
	}
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

	// ログの設定（必ずenvLoad後に実行）
	day := fmt.Sprintf(time.Now().Format("2006-01-02"))
	f, err := os.OpenFile(fmt.Sprintf("./tmp/log/go_oauth_%s.log", day), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	e.Logger.SetOutput(f)
	e.Logger.Info("initialized log settings.")

	setLogLevel(e)

	route(e)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = defaultPort
	}

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
