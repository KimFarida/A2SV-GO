package models

type BookStatus string

const (
	BookAvailable BookStatus = "available"
	BookBorrowed BookStatus = "borrowed"
)

type Book struct {
	ID int
	Title string
	Author string
	Status BookStatus // can  be available or Borrowed -> Enum?
}