package model

import validation "github.com/go-ozzo/ozzo-validation"

// Ad ...
type Ad struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	Descript  string   `json:"desc"`
	Price     float64  `json:"price"`
	CreatedAt string   `json:"created_at"`
	MainPic   string   `json:"mainpic"`
	OtherPics []string `json:"pic_links"`
}

// Validate ...
func (a *Ad) Validate() error {
	return validation.ValidateStruct(a,
		validation.Field(&a.Name, validation.Required, validation.Length(1, 200)),
		validation.Field(&a.Descript, validation.Required, validation.Length(1, 1000)),
		validation.Field(&a.Price, validation.Min(0.00)),
		validation.Field(&a.MainPic, validation.Required, validation.Length(1, 2000)),
		validation.Field(&a.OtherPics, validation.Required, validation.Length(1, 2), validation.Each(validation.Length(1, 2000))),
	)
}
