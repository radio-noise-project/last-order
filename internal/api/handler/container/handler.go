package container

import "github.com/labstack/echo"

func Handler(api *echo.Group) {
	api.POST("/run", runContainer)
}
