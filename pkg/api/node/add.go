package node

import (
	"context"
	"strconv"

	"github.com/google/uuid"
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

	intPort, err := strconv.Atoi(port)
	if err != nil {
		panic(err)
	}

	uuidv4, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}

	sister := &model.Sister{
		SisterID: uuidv4,
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
