package DAL

//
//import "database/sql"
//
//type Book struct {
//	Isbn   string
//	Title  string
//	Author string
//	Price  float32
//}
//
//// Update the AllBooks function so it accepts the connection pool as a
//// parameter.parameter
//func AllBooks(db *sql.DB) ([]Book, error) {
//	rows, err := db.Query("SELECT * FROM books")
//	if err != nil {
//		return nil, err
//	}
//	defer rows.Close()
//
//	var bks []Book
//
//	for rows.Next() {
//		var bk Book
//
//		err := rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
//		if err != nil {
//			return nil, err
//		}
//
//		bks = append(bks, bk)
//	}
//	if err = rows.Err(); err != nil {
//		return nil, err
//	}
//
//	return bks, nil
//}
//func BookByQuery(Books []Book) {
//	for b := range Books {
//		println(&b)
//	}
//
//	type Sell struct {
//		ID           int64
//		Buyer        string
//		Books        []Book
//		Price        float32
//		Discount     float32
//		IsClubMember bool
//	}
