package DAL

import (
	"adraba/internal/common"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"log"
	"testing"
)

func TestNewDb(t *testing.T) {

	dbw, err := BookStoreDBConnection("./booksStore.db")
	if err != nil {
		log.Fatal(err)
	}
	logger := logrus.New()
	for {
		book := &common.Book{uuid.New().String(), uuid.New().String(), uuid.New().String(), 14}
		err := InsertBook(book, dbw, logger)
		if err != nil {
			t.Error("adsfasdf")
		}

		//wg.Done()

	}
}
