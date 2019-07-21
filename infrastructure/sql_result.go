package infrastructure

import (
	"database/sql"
)

type SqlResult struct {
	Result sql.Result
}

func (r *SqlResult) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()
}

func (r *SqlResult) RowAffected() (int64, error) {
	return r.Result.RowsAffected()
}
