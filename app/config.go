package app

import (
	"egoist/internal/database/queries"

	"github.com/jmoiron/sqlx"
)

type Globals struct {
	Queries *queries.Queries
}

func NewGlobal(db *sqlx.DB) *Globals{
	return &Globals{Queries: queries.New(db)}
}