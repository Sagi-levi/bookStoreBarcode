package DAL

import (
	"adraba/internal/common"
	"context"
	"database/sql"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

// books
func GetAllBooks(conn *sql.DB, bulkSize int) (<-chan common.Book, chan error) {
	bookChannel := make(chan common.Book, bulkSize)
	errorChannel := make(chan error, 1)
	var book common.Book
	go func() {
		exec, err := conn.Query("SELECT books.isbn, books.title, books.author,books.price  FROM books")
		defer common.RowsCloser(exec)
		if err != nil {
			errorChannel <- err
			return
		}
		err = exec.Scan(&book.Isbn, &book.Title, &book.Author, &book.Price)
		if err != nil {
			errorChannel <- err
			errClose := exec.Close()
			if err != nil {
				errorChannel <- errClose
			}
		}
		bookChannel <- book

		close(bookChannel)
		close(errorChannel)
	}()
	return bookChannel, errorChannel
}

func GetBookFromDB(isbn string, conn *sql.DB) (*common.Book, error) {
	book := &common.Book{}
	exec, err := conn.Query("SELECT books.isbn, books.title, books.author,books.price  FROM books  WHERE books.isbn=?", isbn)
	defer common.RowsCloser(exec)
	if err != nil {
		return nil, err
	}
	exec.Next()
	err = exec.Scan(&book.Isbn, &book.Title, &book.Author, &book.Price)
	if err != nil {
		return nil, err
	}
	exec.Close()

	return book, nil
}

func DeleteBook(isbn string, conn *sql.DB, logger *logrus.Logger) error {
	tx, err := conn.BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return err
	}
	rows, err := tx.Exec("DELETE FROM books WHERE isbn = ?", isbn)
	if err != nil {
		tx.Rollback()
		return err
	}
	affected, _ := rows.RowsAffected()
	logger.Infof("Deleted %v books , id = %v", affected, isbn)
	err = tx.Commit()
	return err
}

func InsertBook(book *common.Book, conn *sql.DB, logger *logrus.Logger) error {
	logger.Infof("insert subject %v", book.Isbn)
	tx, err := conn.BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return nil
	}
	statement, err := tx.Prepare("INSERT INTO books (isbn,title,author,price) VALUES (?, ?,?,?);")
	if err != nil {
		return err
	}
	_, err = statement.Exec(book.Isbn, book.Title, book.Author, book.Price)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	return err
}

// employees
func GetAllEmployees(conn *sql.DB, bulkSize int) (<-chan common.Employ, chan error) {
	employeesChannel := make(chan common.Employ, bulkSize)
	errorChannel := make(chan error, 1)
	var employ common.Employ
	go func() {
		exec, err := conn.Query("SELECT employees.id, employees.name  FROM employees")
		defer common.RowsCloser(exec)
		if err != nil {
			errorChannel <- err
			return
		}
		err = exec.Scan(&employ.Id, &employ.Name)
		if err != nil {
			errorChannel <- err
			errClose := exec.Close()
			if err != nil {
				errorChannel <- errClose
			}
		}
		employeesChannel <- employ

		close(employeesChannel)
		close(errorChannel)
	}()
	return employeesChannel, errorChannel
}

func GetEmployFromDB(id string, conn *sql.DB) (*common.Employ, error) {
	employ := &common.Employ{}
	exec, err := conn.Query("SELECT employees.id, employees.name, employees.is_active  FROM employees  WHERE employees.id=?", id)
	defer common.RowsCloser(exec)
	if err != nil {
		return nil, err
	}
	exec.Next()
	err = exec.Scan(&employ.Id, &employ.Name, &employ.IsActive)
	if err != nil {
		return nil, err
	}
	exec.Close()

	return employ, nil
}

func DeleteEmploy(id string, conn *sql.DB, logger *logrus.Logger) error {
	tx, err := conn.BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return err
	}
	rows, err := tx.Exec("DELETE FROM employees WHERE id = ?", id)
	if err != nil {
		tx.Rollback()
		return err
	}
	affected, _ := rows.RowsAffected()
	logger.Infof("Deleted %v emoloyees , id = %v", affected, id)
	err = tx.Commit()
	return err
}

func InsertEmploy(employ *common.Employ, conn *sql.DB, logger *logrus.Logger) error {
	logger.Infof("insert employ %v", employ.Id)
	tx, err := conn.BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return nil
	}
	statement, err := tx.Prepare("INSERT INTO employees (id,name,is_active) VALUES (?,?,?);")
	if err != nil {
		return err
	}
	_, err = statement.Exec(employ.Id, employ.Name, employ.IsActive)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	return err
}

// customers
func GetAllCustomers(conn *sql.DB, bulkSize int) (<-chan common.Customer, chan error) {
	customerChannel := make(chan common.Customer, bulkSize)
	errorChannel := make(chan error, 1)
	var customer common.Customer
	go func() {
		exec, err := conn.Query("SELECT customers.id, customers.name, customers.is_club_member,customers.phone_number  FROM customers")
		defer common.RowsCloser(exec)
		if err != nil {
			errorChannel <- err
			return
		}
		err = exec.Scan(&customer.Id, &customer.Name, &customer.IsClubMember, &customer.PhoneNumber)
		if err != nil {
			errorChannel <- err
			errClose := exec.Close()
			if err != nil {
				errorChannel <- errClose
			}
		}
		customerChannel <- customer

		close(customerChannel)
		close(errorChannel)
	}()
	return customerChannel, errorChannel
}

func GetCustomerFromDB(id string, conn *sql.DB) (*common.Customer, error) {
	customer := &common.Customer{}
	exec, err := conn.Query("SELECT customers.id, customers.name, customers.is_club_member,customers.phone_number  FROM customers  WHERE customers.id=?", id)
	defer common.RowsCloser(exec)
	if err != nil {
		return nil, err
	}
	exec.Next()
	err = exec.Scan(&customer.Id, &customer.Name, &customer.IsClubMember, &customer.PhoneNumber)
	if err != nil {
		return nil, err
	}
	exec.Close()

	return customer, nil
}

func DeleteCustomer(id string, conn *sql.DB, logger *logrus.Logger) error {
	tx, err := conn.BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return err
	}
	rows, err := tx.Exec("DELETE FROM customers WHERE id = ?", id)
	if err != nil {
		tx.Rollback()
		return err
	}
	affected, _ := rows.RowsAffected()
	logger.Infof("Deleted %v customers , id = %v", affected, id)
	err = tx.Commit()
	return err
}

func InsertCustomer(customer *common.Customer, conn *sql.DB, logger *logrus.Logger) error {
	logger.Infof("insert customer %v", customer.Id)
	tx, err := conn.BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return nil
	}
	statement, err := tx.Prepare("INSERT INTO customers (id,name,is_club_member,phone_number) VALUES (?,?,?);")
	if err != nil {
		return err
	}
	_, err = statement.Exec(customer.Id, customer.Name, customer.IsClubMember, customer.PhoneNumber)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	return err
}

// sales

func GetAllSells(conn *sql.DB, bulkSize int) (<-chan common.Sell, chan error) {
	sellsChannel := make(chan common.Sell, bulkSize)
	errorChannel := make(chan error, 1)
	var sell common.Sell
	go func() {
		var customerId string
		var employId string
		var booksString string
		exec, err := conn.Query("SELECT sells.id,sells.customer,sells.employ,sells.price,sells.Date,sells.books  FROM sells")
		defer common.RowsCloser(exec)
		if err != nil {
			errorChannel <- err
			return
		}
		var dateString string
		err = exec.Scan(&sell.Id, &sell.Price, &dateString, &customerId, &employId, &booksString)
		if err != nil {
			errorChannel <- err
			errClose := exec.Close()
			if err != nil {
				errorChannel <- errClose
			}
		}
		sell.Date, err = time.Parse(time.RFC3339, dateString)
		if err != nil {
			errorChannel <- err
		}
		customer, err := GetCustomerFromDB(customerId, conn)
		sell.Customer = *customer
		employ, err := GetEmployFromDB(employId, conn)
		sell.Employ = *employ
		books, err := Serializebooks(booksString, conn)
		sell.Books = books
		sellsChannel <- sell

		close(sellsChannel)
		close(errorChannel)
	}()
	return sellsChannel, errorChannel
}

func GetSellFromDB(id string, conn *sql.DB) (*common.Sell, error) {
	sell := &common.Sell{}
	var customerId string
	var employId string
	var booksString string
	exec, err := conn.Query("SELECT sells.id,sells.customer,sells.employ,sells.price,sells.Date,sells.books  FROM sells")
	defer common.RowsCloser(exec)
	if err != nil {
		return nil, err
	}
	var dateString string
	err = exec.Scan(&sell.Id, &sell.Price, &dateString, &customerId, &employId, &booksString)
	if err != nil {
		return nil, err
		errClose := exec.Close()
		if errClose != nil {
			return nil, errClose
		}
	}
	sell.Date, err = time.Parse(time.RFC3339, dateString)
	if err != nil {
		return nil, err
	}
	customer, err := GetCustomerFromDB(customerId, conn)
	sell.Customer = *customer
	employ, err := GetEmployFromDB(employId, conn)
	sell.Employ = *employ
	books, err := Serializebooks(booksString, conn)
	sell.Books = books
	return sell, nil
}

func InsertSell(sell *common.Sell, conn *sql.DB, logger *logrus.Logger) error {
	logger.Infof("insert sell %v", sell.Id)
	tx, err := conn.BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return nil
	}
	statement, err := tx.Prepare("INSERT INTO sells (id,customer,employ,price,Date,books) VALUES (?,?,?,?,?,?);")
	if err != nil {
		return err
	}
	booksString, err := DeserializeBooks(sell.Books)
	_, err = statement.Exec(sell.Id, sell.Customer, sell.Employ, sell.Price, sell.Date.String(), booksString)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	return err
}

func Serializebooks(booksString string, conn *sql.DB) ([]common.Book, error) {
	var books []common.Book
	booksIds := strings.Split(booksString, ",")
	for _, id := range booksIds {
		book, err := GetBookFromDB(id, conn)
		if err != nil {
			return nil, err
		}
		books = append(books, *book)
	}
	return books, nil
}

func DeserializeBooks(books []common.Book) (string, error) {
	var sb strings.Builder
	for _, b := range books {
		sb.WriteString(b.Isbn + ",")
	}
	return sb.String(), nil
}
