package model_test

import (
	"testing"

	"github.com/AlastorTh/avTask/model"
	"github.com/stretchr/testify/assert"
)

func TestAd_Validate(t *testing.T) {
	a := model.TestAd(t)
	assert.NoError(t, a.Validate())
}
