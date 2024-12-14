package node

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/radio-noise-project/last-order/pkg/api/node"
)

type body struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Port    string `json:"port"`
}

func postAddNode(c echo.Context) error {
	data := new(body)
	if err := c.Bind(&data); err != nil {
		panic(err)
	}

	name := data.Name
	address := data.Address
	port := data.Port

	node.AddNode(name, address, port)
	return c.JSON(http.StatusOK, map[string]string{"message": "Node added successfully"})
}

func getNodeList(c echo.Context) error {
	nodes := node.GetNodeList()
	return c.JSON(http.StatusOK, nodes)
}
