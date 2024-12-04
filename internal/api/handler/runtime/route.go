package runtime

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/radio-noise-project/last-order/pkg/api/runtime"
	"github.com/radio-noise-project/last-order/pkg/api/runtime/version"
)

// ping is a handler for checking server status
func getPing(c echo.Context) error {
	u := runtime.GetPing()
	if err := c.Bind(u); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, u)
}

func getVersion(c echo.Context) error {
	arch := c.Param("archtecture")
	if arch == "last-order" {
		u := version.GetLastOrderVersion()
		return c.JSON(http.StatusOK, u)
	} else if arch == "sisters" {
		id := c.QueryParam("sisterId")
		u := version.GetSistersVersion(id)
		return c.JSON(http.StatusOK, u)
	} else {
		return c.JSON(http.StatusBadRequest, "Invalid archtecture")
	}
}
