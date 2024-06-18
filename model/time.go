package model

import (
	"database/sql"
	"time"
)

type Time struct {
	CreatedAt time.Time
	UpdatedAt sql.NullTime
	DeletedAt sql.NullTime
}
