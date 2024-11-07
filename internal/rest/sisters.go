package rest

import (
	"context"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type SistersService interface {
	GetByID(ctx context.Context, id uuid.UUID)
}

type SistersHandler struct {
	Service SistersService
}

func NewSistersHandler(e *echo.Echo, sv SistersService) {
	handler := &SistersHandler{
		Service: sv,
	}

	e.GET("/:id", handler.GetByID)
}

func (s *SistersHandler) GetByID(_ echo.Context) error {
	panic("TODO")
}
