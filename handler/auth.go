package handler

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/h-tko/go_oauth/api"
)

func Auth(c echo.Context) error {
	authAPI := api.NewAuthAPI()
	url, err := authAPI.Auth()
	if err != nil {
		return err
	}

	return c.Redirect(http.StatusMovedPermanently, url)
}
