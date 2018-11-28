package contracts

import (
	"github.com/donbstringham/spider/src/models"
)

// PageCounter defines the count method
type PageCounter interface {
	Count() (uint, error)
}

// PageReader defines the reading method(s)
type PageReader interface {
	Read(rawURL string) (*models.Page, error)
	ReadAll() ([]*models.Page, error)
	ReadByDomain(rawDomain string) (*models.Page, error)
}

// PageRemover defines remove method(s)
type PageRemover interface {
	RemoveAll() error
	RemoveByDomain(rawDomain string) ([]*models.Page, error)
}

// PageUpdater defines the update method
type PageUpdater interface {
	UpdateByDomain(rawDomain string) ([]*models.Page, error)
}

// PageWriter defines the write method
type PageWriter interface {
	Write(p *models.Page) error
}

// PageStorageAdapter combines all interface for persistence
type PageStorageAdapter interface {
	PageCounter
	PageReader
	PageRemover
	PageUpdater
	PageWriter
}
