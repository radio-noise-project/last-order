package node

import (
	"context"
	"strconv"

	"github.com/radio-noise-project/last-order/internal/database"
	"github.com/radio-noise-project/last-order/pkg/database/model"
)

func GetNodeList() map[string]string {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	ctx := context.Background()

	sisters, err := model.Sisters().All(ctx, db)
	if err != nil {
		panic(err)
	}

	nodeList := make(map[string]string)
	for _, sister := range sisters {
		nodeList["name"] = sister.Name
		nodeList["address"] = sister.Address
		nodeList["port"] = strconv.Itoa(sister.Port)
	}

	return nodeList
}
