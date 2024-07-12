package v1

import "log"

func New(l *log.Logger, someDomainObjects someDomainObjectsService) *Contract {
	return &Contract{
		l:                 l,
		someDomainObjects: someDomainObjects,
	}
}

type Contract struct {
	l                 *log.Logger
	someDomainObjects someDomainObjectsService
}
