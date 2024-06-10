package model

import (
	"database/sql"
	"time"
)

type Time struct {
	Created_at time.Time
	Updated_at sql.NullTime
	Deleted_at sql.NullTime
}
