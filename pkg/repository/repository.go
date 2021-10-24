package repository

import (
	"context"
	"github.com/vhbfernandes/xaveco/pkg/models"
)

type Repository interface {
	FindRandom(context.Context, string) (map[string]interface{}, error)
	FindAll(context.Context) ([]*models.Xaveco, error)
	FindByTag(context.Context, string) ([]*models.Xaveco, error)
	Create(context.Context, *models.Xaveco) error
}