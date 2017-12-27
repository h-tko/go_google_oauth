package handler

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/mm-lvgs/mss_auth/api"
)

func Callback(c echo.Context) error {
	form := &api.CallbackRequest{}

	if err := c.Bind(form); err != nil {
		return err
	}

	return c.String(http.StatusOK, "finish!")
}
