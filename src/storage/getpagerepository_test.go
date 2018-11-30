package storage

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestGetPageRepositoryNoAdapter(t *testing.T) {
	// arrange
	// act
	_, err := GetPageRepository("")
	// assert
	assert.NotNil(t, err, "expecting an error")
	assert.Equal(t, "no storage adapter found", err.Error(), "should be equal strings")
}

func TestGetPageRepositoryMemAdapter00(t *testing.T) {
	// arrange
	// act
	actual, err := GetPageRepository("mem")
	// assert
	assert.Nil(t, err, "not expecting an error")
	assert.Equal(t, "*storage.PageRepository", reflect.TypeOf(actual).String(), "should be equal")
}

func TestGetPageRepositoryMemAdapter01(t *testing.T) {
	// arrange
	// act
	actual, err := GetPageRepository("Mem")
	// assert
	assert.Nil(t, err, "not expecting an error")
	assert.Equal(t, "*storage.PageRepository", reflect.TypeOf(actual).String(), "should be equal")
}

func TestGetPageRepositoryMemAdapter02(t *testing.T) {
	// arrange
	// act
	actual, err := GetPageRepository("mEm")
	// assert
	assert.Nil(t, err, "not expecting an error")
	assert.Equal(t, "*storage.PageRepository", reflect.TypeOf(actual).String(), "should be equal")
}

func TestGetPageRepositoryMemAdapter03(t *testing.T) {
	// arrange
	// act
	actual, err := GetPageRepository("meM")
	// assert
	assert.Nil(t, err, "not expecting an error")
	assert.Equal(t, "*storage.PageRepository", reflect.TypeOf(actual).String(), "should be equal")
}
