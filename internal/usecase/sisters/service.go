package sisters

import (
	"context"

	"github.com/google/uuid"
)

//nolint:revive
type SistersReopsitory interface {
	GetByID(ctx context.Context, id uuid.UUID)
}

type Service struct {
	sistersRepo SistersReopsitory
}

func NewService(s SistersReopsitory) *Service {
	return &Service{
		sistersRepo: s,
	}
}

func (*Service) GetByID(_ context.Context, _ uuid.UUID) {
	panic("unimplemented")
}
