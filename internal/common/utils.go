package common

import "database/sql"

func RowsCloser(rows *sql.Rows) {
	if rows != nil {
		rows.Close()
	}
}
