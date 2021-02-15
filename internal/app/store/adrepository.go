package store

import (
	"github.com/AlastorTh/avTask/model"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/lib/pq"
)

// AdRepository ...
type AdRepository struct {
	store *Store
}

// Create ...
func (r *AdRepository) Create(a *model.Ad) (*model.Ad, error) {
	if err := a.Validate(); err != nil {
		return nil, err
	}
	if err := r.store.db.QueryRow(
		"INSERT INTO ads (name, descript, price, mainpic, otherpics) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		a.Name, a.Descript, a.Price, a.MainPic, pq.Array(a.OtherPics)).Scan(&a.ID); err != nil {
		return nil, err
	}
	return a, nil
}

// GetAd ...
func (r *AdRepository) GetAd(id string, fields []string) (*model.Ad, error) {
	optList := []string{"descript", "otherpics", ""}
	err := validation.Validate(fields, validation.Length(0, 2), validation.Each(validation.In(optList)))
	if err != nil {
		return nil, err
	}

	ad := &model.Ad{}

	if err := r.store.db.QueryRow("SELECT * FROM ads WHERE id = $1", id).Scan(&ad.ID, &ad.Name, &ad.Price, &ad.MainPic, &ad.Descript, &ad.OtherPics); err != nil {
		return nil, err
	}

	return nil, nil

}

func (r *AdRepository) GetAdList() error {

	return nil
}
