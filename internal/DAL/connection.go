package DAL

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"time"
)

type ConnectionHandler struct {
	DB  *sql.DB
	log *logrus.Logger
}

func InitConnection() (*ConnectionHandler, error) {
	cfg := mysql.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PASSWORD"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "mydb",
		AllowNativePasswords: true,
	}
	// Get a database handle.
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s(%s)/%s", cfg.User, cfg.Passwd, cfg.Net, cfg.Addr, cfg.DBName))
	log.Println(cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
		return nil, err
	}
	fmt.Println("Connected!")
	db.SetMaxIdleConns(1)
	db.SetMaxOpenConns(1)
	db.SetConnMaxLifetime(time.Duration(3) * time.Minute)
	connectionHandler := ConnectionHandler{}
	connectionHandler.DB = db
	connectionHandler.log = logrus.New()
	return &connectionHandler, nil
}
