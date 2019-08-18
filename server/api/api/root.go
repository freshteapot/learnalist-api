package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (m *Manager) V1GetRoot(c echo.Context) error {
	message := "1, 2, 3. Lets go!"
	response := HttpResponseMessage{
		Message: message,
	}

	return c.JSON(http.StatusOK, response)
}