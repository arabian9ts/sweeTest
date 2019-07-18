package infrastructure

import (
	"database/sql"
	"github.com/arabian9ts/sweeTest/app/interface/database"
	"github.com/arabian9ts/sweeTest/config"
	_ "github.com/go-sql-driver/mysql"
)

type SqlHandler struct {
	Conn *sql.DB
}

func NewSqlHandler() database.SqlHandler {
	settings := config.GetSettings()
	uri := settings.GetRdbUri()

	conn, err := sql.Open("mysql", uri)
	if err != nil {
		panic(err.Error())
	}

	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}

func (handler *SqlHandler) Execute(statement string, args ...interface{}) (database.Result, error) {
	res := new(SqlResult)
	result, err := handler.Conn.Exec(statement, args...)
	if err != nil {
		return res, err
	}

	res.Result = result
	return res, nil
}

func (handler *SqlHandler) Query(statement string, args ...interface{}) (database.Row, error) {
	row := new(SqlRow)
	rows, err := handler.Conn.Query(statement, args...)
	if err != nil {
		return row, err
	}

	row.Rows = rows
	return row, nil
}
