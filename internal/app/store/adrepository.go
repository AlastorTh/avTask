package store

import (
	"github.com/AlastorTh/avTask/model"
	"github.com/lib/pq"
)

// AdRepository ...
type AdRepository struct {
	store *Store
}

// Create ...
func (r *AdRepository) Create(a *model.Ad) (*model.Ad, error) {
	if err := r.store.db.QueryRow(
		"INSERT INTO ads (name, descript, price, piclinks) VALUES ($1, $2, $3, $4) RETURNING id",
		a.Name, a.Descript, a.Price, pq.Array(a.PicLinks)).Scan(&a.ID); err != nil {
		return nil, err
	}
	return a, nil
}
