package node

import "github.com/labstack/echo"

func Handler(api *echo.Group) {
	api.POST("/node/add", postAddNode)
}
