package models

import (
	"database/sql"
	"time"
)

type Time struct {
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}
