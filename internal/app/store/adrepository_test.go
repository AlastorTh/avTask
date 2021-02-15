package store_test

import (
	"testing"

	"github.com/AlastorTh/avTask/internal/app/store"
	"github.com/AlastorTh/avTask/model"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, connectString)
	defer teardown("ads")

	tags := []string{"link2", "link3"}
	a, err := s.Ad().Create(&model.Ad{Name: "TestAd", Descript: "MyTest Yesyes nice", Price: 540.32, MainPic: "linkMain", OtherPics: tags})
	assert.NoError(t, err)
	assert.NotNil(t, a)
}
