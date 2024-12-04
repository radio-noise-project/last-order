package version

import (
	"context"

	"time"

	"github.com/google/uuid"
	"github.com/radio-noise-project/last-order/internal/client"
	"github.com/radio-noise-project/last-order/internal/database"
	"github.com/radio-noise-project/last-order/pkg/api/runtime/types"
	"github.com/radio-noise-project/last-order/pkg/database/model"
)

func GetSistersVersion(id string) *types.ResponseSistersVersion {
	address, port, name := getSistersInfo(id)
	cli, err := client.Connect(address, port)
	if err != nil {
		panic(err)
	}
	res := client.Version(cli)
	defer cli.Close()

	return &types.ResponseSistersVersion{
		CodeName:           res.CodeName,
		Version:            res.Version,
		GolangVersion:      res.GolangVersion,
		SisterId:           id,
		Name:               name,
		Address:            address,
		BuiltGitCommitHash: res.BuiltGitcommitHash,
		BuiltDate:          parseTime(res.BuiltDate),
		Os:                 res.Os,
		Arch:               res.Arch,
	}
}

func parseTime(dateStr string) time.Time {
	t, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		panic(err)
	}
	return t
}

func getSistersInfo(id string) (string, int, string) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	ctx := context.Background()

	uid := uuid.MustParse(id)
	s, err := model.FindSister(ctx, db, uid)
	if err != nil {
		panic(err)
	}
	return s.Address, s.Port, s.Name
}
