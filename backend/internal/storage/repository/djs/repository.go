package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ksusonic/niti/backend/internal/storage/repository/base"
)

type Repository struct {
	*base.Repository
}

func New(pool *pgxpool.Pool) *Repository {
	return &Repository{Repository: base.NewBaseRepository(pool)}
}

// You can add methods for CRUD operations below, for example:

// Example: Get DJ by ID
/*
func (r *Repository) GetDJByID(ctx context.Context, id int64) (*DJ, error) {
	row := r.Pool.QueryRow(ctx, "SELECT id, user_id, stage_name, avatar_url, socials FROM djs WHERE id=$1", id)
	var dj DJ
	err := row.Scan(&dj.ID, &dj.UserID, &dj.StageName, &dj.AvatarURL, &dj.Socials)
	if err != nil {
		return nil, err
	}
	return &dj, nil
}
*/

// Define DJ struct in your models package as needed.
