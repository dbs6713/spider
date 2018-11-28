package models

import (
	"github.com/magiconair/properties/assert"
	"reflect"
	"testing"
)

func TestNewPage(t *testing.T) {
	// arrange
	expected := "http://example.com"
	// act
	actual := NewPage(expected)
	// assert
	assert.Equal(t, reflect.TypeOf(&Page{}), reflect.TypeOf(actual))
	assert.Equal(t, actual.RawUrl, expected)
}

func TestPageToJsonSuccess(t *testing.T) {
	// arrange
	expected := "http://example.com"
	actual := NewPage(expected)
	// act
	j, err := actual.ToJSON()
	// assert
	assert.Equal(t, err, nil)
	assert.Equal(t, j, `{"url":"http://example.com"}`)
}

//func TestPageToJsonError(t *testing.T) {
//	// arrange
//	expected := "http://example.com"
//	actual := NewPage(expected)
//	// act
//	j, err := actual.ToJSON()
//	// assert
//	assert.Equal(t, err, "ERROR: JSON conversion")
//	assert.Equal(t, j, "")
//}
