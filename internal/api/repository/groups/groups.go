package groups

import (
	"context"

	"github.com/jackc/pgx/v5"

	"github.com/alex-bodnar/words/internal/api/domain/groups"
	"github.com/alex-bodnar/words/internal/api/repository"
)

var _ repository.Groups = &Repository{}

//go:generate sqlc generate

// Repository implements repository.Description interface.
type Repository struct {
	db *pgx.Conn
}

// NewRepository constructor.
func NewRepository(db *pgx.Conn) *Repository {
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

// Create new group in database.
func (r *Repository) Create(ctx context.Context, name string) (groups.Group, error) {
	group, err := r.getQueries(ctx).create(ctx, name)
	if err != nil {
		return groups.Group{}, err
	}

	return group.toDomain(), nil
}
