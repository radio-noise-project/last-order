package node

import "github.com/labstack/echo"

func Handler(api *echo.Group) {
	api.GET("/node/add", putAddNode)
}
