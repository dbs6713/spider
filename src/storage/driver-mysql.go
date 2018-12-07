package storage

import (
	"database/sql"
	"errors"
	"github.com/donbstringham/spider/src/models"
	"strconv"

	log "github.com/spf13/jwalterweatherman"

	_ "github.com/go-sql-driver/mysql"
)

type MysqlDriver struct {
	pages []*models.Page
	db    *sql.DB
}

func NewMysqlDriver(u string, p string, h string, prt string, dbn string) *MysqlDriver {
	dsn := u + ":" + p + "@tcp(" + h + ":" + prt + ")/" + dbn + "?parseTime=true&timeout=10s&collation=utf8_general_ci"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.FATAL.Println(err)
		return nil
	}
	if err = db.Ping(); err != nil {
		log.FATAL.Println(err)
		return nil
	}
	db.SetMaxIdleConns(10000)
	return &MysqlDriver{db: db}
}

func (m *MysqlDriver) Count() (uint, error) {
	var nStr string
	var s = "SELECT COUNT(*) FROM t_pages"
	if err := m.db.QueryRow(s).Scan(&nStr); err != nil {
		return 0, err
	}
	r, err := strconv.Atoi(nStr)
	if err != nil {
		return 0, err
	}
	return uint(r), nil
}

func (m *MysqlDriver) Read(rawURL string) (*models.Page, error) {
	for i := range m.pages {
		p := m.pages[i]
		if rawURL == p.RawUrl {
			return p, nil
		}
	}
	return nil, errors.New(rawURL + " not found")
}

func (m *MysqlDriver) ReadAll() ([]*models.Page, error) {
	if len(m.pages) == 0 {
		return m.pages, errors.New("no pages")
	}
	return m.pages, nil
}

func (m *MysqlDriver) RemoveAll() error {
	m.pages = m.pages[:0]
	return nil
}

func (m *MysqlDriver) Write(p *models.Page) error {
	SQL := "INSERT INTO t_pages (" +
		"c_fetched," +
		"c_raw_body," +
		"c_raw_url" +
		") VALUES (?,?,?)"
	sysID, err := m.db.Exec(SQL,
		p.Fetched,
		p.RawBody,
		p.RawUrl,
	)
	if err != nil {
		return err
	}
	// Use sys_id and insert into t_urls table
	SQL = "INSERT INTO t_urls (" +
		"c_page_sys_id," +
		"c_url" +
		") VALUES (?,?)"
	for u := range p.Urls {
		_, err := m.db.Exec(SQL,
			sysID,
			u,
		)
		if err != nil {
			log.CRITICAL.Println(err)
		}
	}
	return nil
}
