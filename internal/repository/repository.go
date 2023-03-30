package repository

import (
	"database/sql"
	"technodom/internal/entity"
)

type Repository interface {
	Close()
	GetList(limit int, offset int) (*sql.Rows, error)
	GetByID(id string) *sql.Row
	Post(url entity.Url) error
	Update(id, activeLink string) error
	Delete(id string) error
}
