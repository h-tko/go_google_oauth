package handler

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/h-tko/go_oauth/api"
)

func Callback(c echo.Context) error {
	form := &api.CallbackRequest{}

	if err := c.Bind(form); err != nil {
		return err
	}

	authAPI := api.NewAuthAPI()
	tokenInfo, err := authAPI.Callback(form)
	if err != nil {
		return err
	}

	c.Logger().Warnf("%#v", tokenInfo)

	return c.String(http.StatusOK, "finish!")
}
