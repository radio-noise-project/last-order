package node

import (
	"context"
	"strconv"

	"github.com/gofrs/uuid"
	"github.com/radio-noise-project/last-order/internal/database"
	"github.com/radio-noise-project/last-order/pkg/database/model"
	"github.com/volatiletech/sqlboiler/boil"
)

func AddNode(name string, host string, port string) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	ctx := context.Background()

	uuid, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}

	intPort, err := strconv.Atoi(port)
	if err != nil {
		panic(err)
	}
	sister := &model.Sister{
		SisterID: uuid.String(),
		Name:     name,
		Role:     0,
		Address:  host,
		Port:     intPort,
	}

	err = sister.Insert(ctx, db, boil.Blacklist())
	if err != nil {
		panic(err)
	}
}
