package DAL

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

func BookStoreDBConnection(dbName string) (*sql.DB, error) {
	if _, err := os.Stat(dbName); os.IsNotExist(err) {
		file, err := os.Create(dbName)
		if err != nil {
			return nil, err
		}
		defer file.Close()
	}
	sqliteDatabase, err := sql.Open("sqlite3", fmt.Sprintf("%v?_journal_mode=WAL", dbName))
	if err != nil {
		return nil, err
	}
	err = createBooksTableIfNotExists(sqliteDatabase)
	if err != nil {
		return nil, err
	}
	err = createCustomerTableIfNotExists(sqliteDatabase)
	if err != nil {
		return nil, err
	}
	err = createSellsTableIfNotExists(sqliteDatabase)
	if err != nil {
		return nil, err
	}
	err = createEmployeesTableIfNotExists(sqliteDatabase)
	if err != nil {
		return nil, err
	}
	return sqliteDatabase, nil
}
