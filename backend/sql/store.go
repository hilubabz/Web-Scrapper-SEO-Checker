package sql

import (
	"github.com/hilubabz/web-scraper-seo/sql/generated"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Store struct {
	*generated.Queries
}

func NewStore(pool *pgxpool.Pool) *Store {
	return &Store{
		Queries: generated.New(pool),
	}
}