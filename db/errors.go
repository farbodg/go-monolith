package db

import (
	"database/sql"
	"errors"
)

func IsDALNoRowsError(err error) bool {
	return errors.Is(err, sql.ErrNoRows)
}
