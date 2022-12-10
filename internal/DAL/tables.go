package DAL

import (
	"database/sql"
)

func createBooksTableIfNotExists(sqliteDatabase *sql.DB) error {
	statement, err := sqliteDatabase.Prepare("CREATE TABLE IF NOT EXISTS \"books\" ( \"isbn\" TEXT NOT NULL UNIQUE,\"title\" TEXT NOT NULL,\"author\" TEXT NOT NULL,\"price\" REAL NOT NULL);")
	if err != nil {
		return err
	}
	_, err = statement.Exec()
	if err != nil {
		return err
	}
	return nil
}

func createCustomerTableIfNotExists(sqliteDatabase *sql.DB) error {
	statement, err := sqliteDatabase.Prepare("CREATE TABLE IF NOT EXISTS \"customers\" ( \"id\" TEXT NOT NULL UNIQUE,\"name\" TEXT NOT NULL , \"is_club_member\" NUMERIC NOT NULL DEFAULT 0,\"phone_number\" TEXT);")
	if err != nil {
		return err
	}
	_, err = statement.Exec()
	if err != nil {
		return err
	}
	return nil
}

func createSellsTableIfNotExists(sqliteDatabase *sql.DB) error {
	statement, err := sqliteDatabase.Prepare("CREATE TABLE IF NOT EXISTS \"sells\" ( \"id\" TEXT NOT NULL UNIQUE,\"customer\" TEXT NOT NULL,\"employ\" TEXT NOT NULL, \"books\" TEXT NOT NULL,\"price\" REAL NOT NULL,\"Date\" NUMERIC NOT NULL);")
	if err != nil {
		return err
	}
	_, err = statement.Exec()
	if err != nil {
		return err
	}
	return nil
}

func createEmployeesTableIfNotExists(sqliteDatabase *sql.DB) error {
	statement, err := sqliteDatabase.Prepare("CREATE TABLE IF NOT EXISTS \"employees\" (\"id\" TEXT NOT NULL UNIQUE,\"name\" TEXT NOT NULL , \"is_active\" NUMERIC NOT NULL DEFAULT 1);")
	if err != nil {
		return err
	}
	_, err = statement.Exec()
	if err != nil {
		return err
	}
	return nil
}
