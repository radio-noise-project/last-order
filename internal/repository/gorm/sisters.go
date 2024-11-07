package gorm

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SistersRepository struct {
	db *gorm.DB
}

func NewSistersRepository(db *gorm.DB) *SistersRepository {
	return &SistersRepository{db}
}

func (r *SistersRepository) GetByID(_ context.Context, _ uuid.UUID) {
	panic("TODO")
}
