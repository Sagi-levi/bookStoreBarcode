package DAL

import (
	"adraba/internal/common"
	"database/sql"
	"github.com/sirupsen/logrus"
)

type Repository struct {
	conn   *sql.DB
	logger *logrus.Logger
	dbName string
}

func NewRepository(dbName string, logger *logrus.Logger) (*Repository, error) {
	connection, err := BookStoreDBConnection(dbName)
	if err != nil {
		return nil, err
	}
	return &Repository{conn: connection, logger: logger, dbName: dbName}, nil
}

func (r *Repository) IsSubjectExistsInDB(book *common.Book) error {
	return InsertBook(book, r.conn, r.logger)
}

//func (r *Repository) DeleteSubject(id string) error {
//	return DeleteSubject(id, r.conn, r.logger)
//}
//
//func (r *Repository) InsertSubject(subject *common.Subjects) error {
//	return InsertSubject(subject, r.conn, r.logger)
//}

//func (r *Repository) Flush() error {
//	var err error
//	r.conn, err = Flush(r.dbName)
//	if err != nil {
//		return err
//	}
//	return nil
//}

func (r *Repository) Close() error {
	err := r.conn.Close()
	if err != nil {
		return err
	}
	return nil
}
func (r *Repository) Ping() error {
	err := r.conn.Ping()
	if err != nil {
		return err
	}
	return nil
}
