package storage

import (
	"github.com/donbstringham/spider/src/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMemDriver_Count00(t *testing.T) {
	// arrange
	harness := NewMemDriver()
	p0 := models.NewPage("http://example.com")
	p1 := models.NewPage("http://google.com")
	err := harness.Write(p0)
	if err != nil {
		t.Error("count not create a new page object (p0)")
	}
	err = harness.Write(p1)
	if err != nil {
		t.Error("count not create a new page object (p1)")
	}
	// act
	actual, err := harness.Count()
	// assert
	assert.Nil(t, err)
	assert.Equal(t, uint(2), actual)
}

func TestMemDriver_Count01(t *testing.T) {
	// arrange
	harness := NewMemDriver()
	// act
	actual, err := harness.Count()
	// assert
	assert.Nil(t, err)
	assert.Equal(t, uint(0), actual)
}

func TestMemDriver_RemoveAll(t *testing.T) {
	// arrange
	harness := NewMemDriver()
	p0 := models.NewPage("http://example.com")
	p1 := models.NewPage("http://google.com")
	err := harness.Write(p0)
	if err != nil {
		t.Error("count not create a new page object (p0)")
	}
	err = harness.Write(p1)
	if err != nil {
		t.Error("count not create a new page object (p1)")
	}
	// act
	err = harness.RemoveAll()
	if err != nil {
		t.Error(err)
	}
	actual, err := harness.Count()
	// assert
	assert.Nil(t, err)
	assert.Equal(t, uint(0), actual)
}

func TestMemDriver_Read(t *testing.T) {
	// arrange
	harness := NewMemDriver()
	p0 := models.NewPage("http://example.com")
	p1 := models.NewPage("http://google.com")
	p2 := models.NewPage("http://bing.com")
	err := harness.Write(p0)
	if err != nil {
		t.Error("count not create a new page object (p0)")
	}
	err = harness.Write(p1)
	if err != nil {
		t.Error("count not create a new page object (p1)")
	}
	err = harness.Write(p2)
	if err != nil {
		t.Error("count not create a new page object (p2)")
	}
	// act
	p, err := harness.Read("http://google.com")
	// assert
	assert.Nil(t, err)
	assert.Equal(t, "http://google.com", p.RawUrl)
}

func TestMemDriver_ReadAll(t *testing.T) {
	// arrange
	harness := NewMemDriver()
	p0 := models.NewPage("http://example.com")
	p1 := models.NewPage("http://google.com")
	p2 := models.NewPage("http://bing.com")
	err := harness.Write(p0)
	if err != nil {
		t.Error("count not create a new page object (p0)")
	}
	err = harness.Write(p1)
	if err != nil {
		t.Error("count not create a new page object (p1)")
	}
	err = harness.Write(p2)
	if err != nil {
		t.Error("count not create a new page object (p2)")
	}
	// act
	p, err := harness.ReadAll()
	// assert
	assert.Nil(t, err)
	assert.Equal(t, 3, len(p))
}
