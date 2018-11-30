package storage

import (
	"errors"
	"github.com/donbstringham/spider/src/contracts"
	"github.com/donbstringham/spider/src/models"
	"strings"
)

type PageRepository struct {
	adapter contracts.PageStorageAdapter
}

func GetPageRepository(adapter string, args ...string) (*PageRepository, error) {
	var a contracts.PageStorageAdapter

	switch strings.ToLower(adapter) {
	case "mem":
		a = NewMemDriver()
		break
	//case "mysql":
	//	drv, err := NewMySQLDriver(args[0], args[1], args[2], args[3], args[4])
	//	if err != nil {
	//		return nil, err
	//	}
	//	a = drv
	default:
		return nil, errors.New("no storage adapter found")
	}
	return New(a), nil
}

func New(a contracts.PageStorageAdapter) *PageRepository {
	return &PageRepository{adapter: a}
}

func (r *PageRepository) Add(p *models.Page) error {
	return r.adapter.Write(p)
}

func (r *PageRepository) All() ([]*models.Page, error) {
	return r.adapter.ReadAll()
}

func (r *PageRepository) ByURL(rawURL string) (*models.Page, error) {
	return nil, nil
}

func (r *PageRepository) Count() (uint, error) {
	return r.adapter.Count()
}
