package infrastructure

import "database/sql"

type SqlRow struct {
	Rows *sql.Rows
}

func (r *SqlRow) Scan(dest ...interface{}) error {
	return r.Rows.Scan(dest...)
}

func (r *SqlRow) Next() bool {
	return r.Rows.Next()
}

func (r *SqlRow) Close() error {
	return r.Rows.Close()
}
