package repository

import (
	"database/sql"
	models "mygo/internal/model"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetBannersHandler() ([]models.Banner, error) {
	query := `
		SELECT * 
		FROM banners 
		WHERE is_active = TRUE
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var banners []models.Banner
	for rows.Next() {
		var b models.Banner
		if err := rows.Scan(&b.ID, &b.JSONData, &b.FeatureID, &b.IsActive); err != nil {
			return nil, err
		}

		banners = append(banners, b)
	}

	return banners, nil
}

func (r *Repository) CreateBannerHandler(banner models.Banner) error {
	query := `
		INSERT INTO banners (json_data, feature_id, is_active) 
		VALUES ($1, $2, $3)
	`
	_, err := r.db.Exec(query, banner.JSONData, banner.FeatureID, banner.IsActive)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateBannerHandler(banner models.Banner) error {
	query := `
		UPDATE banners 
		SET json_data = $1, feature_id = $2, is_active = $3 
		WHERE id = $4
	`
	_, err := r.db.Exec(query, banner.JSONData, banner.FeatureID, banner.IsActive, banner.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) DeleteBannerHandler(id int) error {
	query := `
		DELETE FROM banners 
		WHERE id = $1
	`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
