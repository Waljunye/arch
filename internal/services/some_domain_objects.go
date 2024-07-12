package services

import (
	"arch/internal/stores/models"
	"context"
)

func NewSomeDomainObjects(
	storeSomeObject someObjectStore,
) *someDomainObject {
	return &someDomainObject{
		storeSomeObject: storeSomeObject,
	}
}

type someDomainObject struct {
	storeSomeObject someObjectStore
}

func (sdo *someDomainObject) DoBuisnessLogic(ctx *context.Context, uid string) (result *models.SomeObject, err error) {
	// Do business logic using sdo.storeSomeObject and other stores, that we want to use
	return
}
