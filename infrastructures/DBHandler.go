package infrastructures

import (
	"database/sql"
	"fmt"
	"github.com/djamboe/mtools-login-service/interfaces"
)

type DBHandler struct {
	Conn *sql.DB
}

func (handler *DBHandler) Execute(statement string) {
	handler.Conn.Exec(statement)
}

func (handler *DBHandler) Query(statement string) (interfaces.IRow, error) {
	rows, err := handler.Conn.Query(statement)

	if err != nil {
		fmt.Println(err)
		return new(DatabaseRow), err
	}
	row := new(DatabaseRow)
	row.Rows = rows
	return row, nil
}

type DatabaseRow struct {
	Rows *sql.Rows
}

func (r DatabaseRow) Scan(dest ...interface{}) error {
	err := r.Rows.Scan(dest...)
	if err != nil {
		return err
	}
	return nil
}

func (r DatabaseRow) Next() bool {
	return r.Rows.Next()
}
