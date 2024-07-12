package v1

import (
	"arch/internal/stores/models"
	"context"
)

type someDomainObjectsService interface {
	DoBuisnessLogic(ctx *context.Context, uid string) (result *models.SomeObject, err error)
}
