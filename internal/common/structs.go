package common

import "time"

type Book struct {
	Isbn   string
	Title  string
	Author string
	Price  float32
}
type Customer struct {
	Id           string
	Name         string
	IsClubMember bool
	PhoneNumber  string
}
type Employ struct {
	Id       string
	Name     string
	IsActive bool
}
type Sell struct {
	Id string
	Customer
	Employ
	Price float32
	Date  time.Time
	Books []Book
}
