package entity

import (
	"database/sql"
	"time"
)

type Tea struct {
	Id          uint64         `json:"id" db:"id"`
	Name        string         `json:"name" db:"name"`
	Description sql.NullString `json:"description,omitempty" db:"description,omitempty"`
	Price       int            `json:"price" db:"price"`
	UpdatedAt   time.Time      `json:"updated_at" db:"updated_at"`
	IdCategory  uint64         `json:"id_category" db:"id_category"`
}
