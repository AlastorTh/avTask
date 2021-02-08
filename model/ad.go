package model

// Ad ...
type Ad struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Descript string `json:"desc"`
	Price    float64
	PicLinks []string `json:"pic_links"`
}
