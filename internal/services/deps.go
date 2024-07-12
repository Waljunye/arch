package services

import (
	"arch/internal/stores/models"
	"context"
)

type someObjectStore interface {
	ByUid(ctx *context.Context, uid string) (result *models.SomeObject, err error)
}
