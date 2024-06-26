package repository

import (
	"github.com/MaksKazantsev/Chatter/user/internal/db"
	"github.com/jmoiron/sqlx"
)

type Postgres struct {
	*sqlx.DB
}

var _ db.Repository = &Postgres{}

func NewRepository(db *sqlx.DB) *Postgres {
	return &Postgres{
		db,
	}
}
