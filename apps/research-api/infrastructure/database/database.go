package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type SQLHandler struct {
	Conn *gorm.DB
}

func NewSQLHandler(connInfo string, verbose bool) (*SQLHandler, error) {
	db, err := gorm.Open("postgres", connInfo)
	if err != nil {
		return nil, err
	}
	if verbose {
		db.LogMode(true)
	}
	return &SQLHandler{
		Conn: db,
	}, nil
}

func (h *SQLHandler) Close() error {
	return h.Conn.Close()
}
