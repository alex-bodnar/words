package words

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/alex-bodnar/words/internal/api/domain/words"
	"github.com/alex-bodnar/words/internal/api/repository"
)

var _ repository.Words = &Repository{}

//go:generate sqlc generate

// Repository implements repository.Description interface.
type Repository struct {
	db *pgxpool.Pool
}

// NewRepository constructor.
func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
	}
}

// getQueries gets queries from context or creates new one.
func (r *Repository) getQueries(ctx context.Context) *Queries {
	tx := repository.ExtractTx(ctx)
	if tx != nil {
		return New(*tx)
	}

	return New(r.db)
}

// Create new word in database.
func (r *Repository) Create(ctx context.Context, val words.Word) (words.Word, error) {
	desc, err := r.getQueries(ctx).create(ctx, toCreateParams(val))
	if err != nil {
		return words.Word{}, err
	}

	return desc.toDomain(), nil
}
