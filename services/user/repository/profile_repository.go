package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/you/connect-market/services/user/models"
)

type ProfileRepository struct {
	db *pgxpool.Pool
}

func NewProfileRepository(db *pgxpool.Pool) *ProfileRepository {
	return &ProfileRepository{db: db}
}

func (r *ProfileRepository) Get(ctx context.Context, id int64) (*models.Profile, error) {
	var p models.Profile
	err := r.db.QueryRow(ctx, "SELECT id, user_id, display_name, bio, tags, created_at FROM profiles WHERE id=$1", id).
		Scan(&p.ID, &p.UserID, &p.DisplayName, &p.Bio, &p.Tags, &p.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *ProfileRepository) Update(ctx context.Context, p *models.Profile) error {
	_, err := r.db.Exec(ctx,
		"UPDATE profiles SET display_name=$1, bio=$2, tags=$3 WHERE id=$4",
		p.DisplayName, p.Bio, p.Tags, p.ID)
	return err
}

func (r *ProfileRepository) SearchByTag(ctx context.Context, tag string) ([]models.Profile, error) {
	rows, err := r.db.Query(ctx, "SELECT id, user_id, display_name, bio, tags, created_at FROM profiles WHERE $1 = ANY(tags)", tag)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var res []models.Profile
	for rows.Next() {
		var p models.Profile
		rows.Scan(&p.ID, &p.UserID, &p.DisplayName, &p.Bio, &p.Tags, &p.CreatedAt)
		res = append(res, p)
	}
	return res, nil
}
