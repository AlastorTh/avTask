package model

import "testing"

func TestAd(t *testing.T) *Ad {
	t.Helper()

	return &Ad{
		Name:      "TestName",
		Descript:  "TestDescript",
		Price:     245.22,
		MainPic:   "link1",
		OtherPics: []string{"link2", "link3"},
	}
}
