package stores

import (
	"arch/internal/stores/models"
	"context"
	"database/sql"
)

func NewSomeObjects(db *sql.DB) *someObjects {
	return &someObjects{
		db: db,
	}
}

type someObjects struct {
	db *sql.DB
}

func (so *someObjects) ByUid(ctx *context.Context, uid string) (result *models.SomeObject, err error) {
	// Get some_object from database using so.db

	return
}

// other functions, that we want to use
