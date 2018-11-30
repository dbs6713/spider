package storage

import (
	"github.com/donbstringham/spider/src/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPageRepository_Add(t *testing.T) {
	// arrange
	harness, err := GetPageRepository("mem")
	if err != nil {
		t.Error("could not get a repository")
	}
	p := models.NewPage("http://example.com")
	// act
	err = harness.Add(p)
	assert.Nil(t, err)
	actual, err := harness.Count()
	// assert
	assert.Nil(t, err)
	assert.Equal(t, uint(1), actual)
}
