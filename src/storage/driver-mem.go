package storage

import (
	"errors"
	"github.com/donbstringham/spider/src/models"
)

type MemDriver struct {
	pages []*models.Page
}

func NewMemDriver() *MemDriver {
	return &MemDriver{}
}

func (m *MemDriver) Count() (uint, error) {
	return uint(len(m.pages)), nil
}

func (m *MemDriver) Read(rawURL string) (*models.Page, error) {
	for i := range m.pages {
		p := m.pages[i]
		if rawURL == p.RawUrl {
			return p, nil
		}
	}
	return nil, errors.New(rawURL + " not found")
}

func (m *MemDriver) ReadAll() ([]*models.Page, error) {
	if len(m.pages) == 0 {
		return m.pages, errors.New("no pages")
	}
	return m.pages, nil
}

func (m *MemDriver) RemoveAll() error {
	m.pages = m.pages[:0]
	return nil
}

func (m *MemDriver) Write(p *models.Page) error {
	m.pages = append(m.pages, p)
	return nil
}
